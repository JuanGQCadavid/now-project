package main

import (
	"testing"

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
