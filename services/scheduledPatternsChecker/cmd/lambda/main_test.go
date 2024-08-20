package main

import (
	"encoding/json"
	"strconv"
	"sync"
	"testing"

	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/cmd/lambda/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
)

var (
	scheduleBody string = `
	{
		"scheduleId":"Holi"
	}
	`

	generateBatchBody string = `
	{
		"Operation":"generateDatesFromSchedulePatterns",
		"TimeWindow": 60
	}
	`
)

func TestScheduleEventsPipeline(t *testing.T) {
	var (
		done chan interface{} = make(chan interface{})
	)

	var tests = []struct {
		name             string
		spAppendedCount  int
		spResumedCount   int
		spConcludedCount int
		spFreezedCount   int
	}{
		{
			name:             "All Append",
			spAppendedCount:  5,
			spResumedCount:   0,
			spConcludedCount: 0,
			spFreezedCount:   0,
		},
		{
			name:             "All Resumed",
			spAppendedCount:  0,
			spResumedCount:   5,
			spConcludedCount: 0,
			spFreezedCount:   0,
		},
		{
			name:             "Creation operations",
			spAppendedCount:  5,
			spResumedCount:   5,
			spConcludedCount: 0,
			spFreezedCount:   0,
		},
		{
			name:             "All Conclude",
			spAppendedCount:  0,
			spResumedCount:   0,
			spConcludedCount: 5,
			spFreezedCount:   0,
		},
		{
			name:             "All Frezed",
			spAppendedCount:  0,
			spResumedCount:   0,
			spConcludedCount: 0,
			spFreezedCount:   5,
		},
		{
			name:             "Delete Operations",
			spAppendedCount:  0,
			spResumedCount:   0,
			spConcludedCount: 5,
			spFreezedCount:   5,
		},
		{
			name:             "Mixed",
			spAppendedCount:  5,
			spResumedCount:   5,
			spConcludedCount: 5,
			spFreezedCount:   5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				creationRecords int      = tt.spAppendedCount + tt.spResumedCount
				deleteRecords   int      = tt.spConcludedCount + tt.spFreezedCount
				totalRecords    int      = creationRecords + deleteRecords
				records         []Record = make([]Record, totalRecords)
				recCounter      int
				events          chan Record = make(chan Record, totalRecords)
			)

			for operation, count := range map[Operations]int{
				SchedulePatternAppended:  tt.spAppendedCount,
				SchedulePatternConcluded: tt.spConcludedCount,
				SchedulePatternFreezed:   tt.spFreezedCount,
				SchedulePatternResumed:   tt.spResumedCount,
			} {
				for i := 0; i < count; i++ {
					records[recCounter] = createDummyRecord(operation, recCounter)
					events <- records[recCounter]
					recCounter++
				}
			}
			close(events)

			createStream, deleteStream := scheduleEventsIdentifer(done, events)
			checker := make([]bool, totalRecords)

			var wg sync.WaitGroup

			if creationRecords > 0 {
				t.Log("Checking creation results")
				wg.Add(1)
				go func() {
					defer wg.Done()
					for result := range createStream {
						index, _ := strconv.Atoi(result.SpotId)
						checker[index] = true
						t.Logf("%+v\n", result)
					}
				}()

			}

			if deleteRecords > 0 {
				t.Log("Checking deletion results")
				wg.Add(1)
				go func() {
					defer wg.Done()
					for result := range deleteStream {
						index, _ := strconv.Atoi(result)
						checker[index] = true
						t.Logf("%+v\n", result)
					}
				}()
			}

			wg.Wait()

			for index, ok := range checker {
				if !ok {
					t.Error("Spot/Scheduler with ID: ", index, " was not returned")
				}
			}
		})
	}

}

func createDummyRecord(operation Operations, counter int) Record {
	var id string = strconv.Itoa(counter)
	var body utils.Body = utils.Body{
		ScheduleId: id,
		SpotId:     id,
		UserId:     id,
		SpotRequest: utils.SpotRequest{
			SpotInfo: utils.SpotInfo{
				SpotId: id,
			},
			SpotPatterns: []utils.SpotPatterns{
				{
					PatternId: id,
				},
			},
		},
	}

	jsonBytes, _ := json.Marshal(body)

	return Record{
		operation: operation,
		body:      string(jsonBytes),
	}

}

func TestEmitRecordsEventsDispatched(t *testing.T) {
	var (
		done chan interface{} = make(chan interface{})
	)

	var singleRecord = events.SQSMessage{
		MessageAttributes: map[string]events.SQSMessageAttribute{
			Operation: {
				StringValue: aws.String(string(SchedulePatternAppended)),
			},
		},
		Body: scheduleBody,
	}

	var tests = []struct {
		name          string
		baseRecord    events.SQSMessage
		numberOfTimes int
	}{
		{
			"0 records where sent",
			singleRecord,
			0,
		},
		{
			"1 record where sent",
			singleRecord,
			1,
		},
		{
			"1000 record where sent",
			singleRecord,
			1000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			payload := make([]events.SQSMessage, tt.numberOfTimes)

			for i := 0; i < tt.numberOfTimes; i++ {
				payload[i] = tt.baseRecord
			}

			resultingChan := emitRecords(done, &events.SQSEvent{Records: payload})

			var counter int
			for range resultingChan {
				counter++
			}
			if counter != tt.numberOfTimes {
				t.Errorf("got %d, want %d", counter, tt.numberOfTimes)
			}

		})
	}
}

func TestEmitRecordsOperation(t *testing.T) {
	var (
		done chan interface{} = make(chan interface{})
	)

	var tests = []struct {
		name              string
		payload           events.SQSEvent
		operationExpected Operations
	}{
		{
			"Operation taken from the message attributes",
			events.SQSEvent{
				Records: []events.SQSMessage{
					{
						MessageAttributes: map[string]events.SQSMessageAttribute{
							Operation: {
								StringValue: aws.String(string(SchedulePatternAppended)),
							},
						},
						Body: scheduleBody,
					},
				},
			},
			SchedulePatternAppended,
		},
		{
			"Operation taken from the message body",
			events.SQSEvent{
				Records: []events.SQSMessage{
					{
						Body: generateBatchBody,
					},
				},
			},
			GenrateDatesFromSchedulePatterns,
		},
		{
			"Operation taken from the message body omitting attributes",
			events.SQSEvent{
				Records: []events.SQSMessage{
					{ // one with Operation set in message attributes but wrong, leaving it on body
						MessageAttributes: map[string]events.SQSMessageAttribute{
							Operation: {
								StringValue: aws.String("AmIWrong"),
							},
						},
						Body: generateBatchBody,
					},
				},
			},
			GenrateDatesFromSchedulePatterns,
		},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			resultingChan := emitRecords(done, &tt.payload)
			ans := <-resultingChan
			if ans.operation != tt.operationExpected {
				t.Errorf("got %s, want %s", string(ans.operation), string(tt.operationExpected))
			}
		})
	}
}
