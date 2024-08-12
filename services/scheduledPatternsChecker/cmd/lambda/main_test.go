package main

import (
	"os"
	"testing"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	credFinderCore "github.com/JuanGQCadavid/now-project/services/pkgs/credentialsFinder/core/core/domain"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/stretchr/testify/mock"
)

func init() {
	os.Setenv("neo4jUser", "DUMMY")
	os.Setenv("neo4jPassword", "DUMMY")
	os.Setenv("neo4jUri", "DUMMY")
}

// Mock implementation of the SSMCredentialsFinder interface
type MockSSMCredentialsFinder struct {
	mock.Mock
}

func (m *MockSSMCredentialsFinder) FindNeo4jCredentialsFromDefaultEnv() (neo4j.Driver, error) {
	args := m.Called()
	return args.Get(0).(neo4j.Driver), args.Error(1)
}
func (m *MockSSMCredentialsFinder) GetDBCredentialsFromDefaultEnv() (*credFinderCore.DBCredentials, error) {
	args := m.Called()
	return args.Get(0).(*credFinderCore.DBCredentials), args.Error(1)
}

// mockCredsFinder := new(MockSSMCredentialsFinder)
// // mockDriver := nil // Replace with your actual driver type
// mockCredsFinder.On("FindNeo4jCredentialsFromDefaultEnv").Return(nil, nil)

func TestEmitRecords(t *testing.T) {
	var (
		done              chan interface{} = make(chan interface{})
		generateBatchBody string           = `
		{
			"Operation":"generateDatesFromSchedulePatterns",
			"TimeWindow": 60
		}
		`
		scheduleBody string = `
		{
			"scheduleId":"Holi"
		}
		`
		payload events.SQSEvent = events.SQSEvent{
			Records: []events.SQSMessage{
				{ // one with Operation set in message attributes
					MessageAttributes: map[string]events.SQSMessageAttribute{
						Operation: {
							StringValue: aws.String(string(SchedulePatternAppended)),
						},
					},
					Body: scheduleBody,
				},
				{ // one with Operation set in body
					Body: generateBatchBody,
				},
				{ // one with Operation set in message attributes but wrong, leaving it on body
					MessageAttributes: map[string]events.SQSMessageAttribute{
						Operation: {
							StringValue: aws.String("AmIWrong"),
						},
					},
					Body: generateBatchBody,
				},
			},
		}
	)
	resultingChan := emitRecords(done, &payload)

	for value := range resultingChan {
		logs.Info.Printf("Op: %s, body: %s \n", string(value.operation), value.body)
	}

	logs.Info.Println("Done")

}
