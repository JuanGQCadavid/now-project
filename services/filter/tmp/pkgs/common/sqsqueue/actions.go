package sqsqueue

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var (
	ErrSQSArnEnvNotFound   = errors.New("The sqs arn env is empty or does not exist")
	ErrParsingMessage      = errors.New("We found an error while parsing the sqs message")
	ErrOnSQSPublishMessage = errors.New("We found a problem while sending the message")
)

type SQSQueueActions struct {
	sqsService *sqs.SQS
	targetArn  string
	bulkJumper int
}

func NewSQSQueueActionsFromEnv(envName string) (*SQSQueueActions, error) {
	sqsArn, isPresent := os.LookupEnv(envName)
	if !isPresent {
		logs.Error.Println(envName, " is not configured in the env.")
		return nil, ErrSQSArnEnvNotFound
	}
	return NewSQSTopicActionsFromArn(sqsArn)
}

func NewSQSTopicActionsFromArn(snsArn string) (*SQSQueueActions, error) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)

	return &SQSQueueActions{
		sqsService: svc,
		targetArn:  snsArn,
		bulkJumper: 10,
	}, nil

}

func (action *SQSQueueActions) SendMessage(payload interface{}) error {

	if payload == nil {
		logs.Info.Println("Empty payload, aborting job")
		return nil
	}

	bodyMarshaled, err := json.Marshal(&payload)

	if err != nil {
		return ErrParsingMessage
	}

	input := &sqs.SendMessageInput{
		QueueUrl:    &action.targetArn,
		MessageBody: aws.String(string(bodyMarshaled)),
	}

	out, err := action.sqsService.SendMessage(input)

	if err != nil {
		logs.Error.Println("We found an error while sending message to queue ", action.targetArn, " Error:", err.Error())
	}

	logs.Info.Printf("SQS send message done, output: %s \n", out)

	return nil
}

func (action *SQSQueueActions) SendBulkMessages(payload []interface{}) map[interface{}]error {

	var startPointer int
	var threshold int
	errors := make(map[interface{}]error)

	if len(payload) == 0 {
		logs.Info.Println("Empty payload, aborting job")
		return nil
	}

	for {
		threshold = startPointer + action.bulkJumper
		var err map[interface{}]error = nil

		if threshold < len(payload) {
			err = action.sendBuklMessage(payload[startPointer:threshold])
			startPointer = threshold
		} else if (len(payload)-startPointer) > 0 && (len(payload)-startPointer) <= action.bulkJumper {
			err = action.sendBuklMessage(payload[startPointer:])
			startPointer = len(payload)
		} else {
			break
		}

		if err != nil && len(err) > 0 {
			logs.Error.Println("We got an error while processing the bulk operation on bacth")
			for key, errrordata := range err {
				errors[key] = errrordata
			}
		}
	}

	if errors != nil && len(errors) > 0 {
		logs.Error.Println("We found erros while processing the bacth request!")
		for key, err := range errors {
			logs.Error.Printf("%+v - %s \n", key, err.Error())
		}
		return errors
	}

	logs.Info.Println("All messages where sent successfully")
	return nil
}

// How should we handle this kind of errors ?
func (action *SQSQueueActions) sendBuklMessage(payload []interface{}) map[interface{}]error {
	entries := make([]*sqs.SendMessageBatchRequestEntry, len(payload))
	date := strings.ReplaceAll(time.Now().Format(time.DateTime), " ", "_")
	date = strings.ReplaceAll(date, ":", "_")

	errors := make(map[interface{}]error)

	index := 0
	for _, entry := range payload {
		bodyMarshaled, err := json.Marshal(&entry)

		if err != nil {
			logs.Error.Printf("We face an error while marshaling the interface for payload %+v \n", entry)
			errors[entry] = ErrParsingMessage
			continue
		}

		id := fmt.Sprintf("%d_%s", index, date)
		logs.Info.Println("Id: ", id)
		entries[index] = &sqs.SendMessageBatchRequestEntry{
			Id:          aws.String(id),
			MessageBody: aws.String(string(bodyMarshaled)),
		}
		index++
	}

	if index < (len(payload) - 1) {
		logs.Error.Println("As we found some messages with errors then we shrink the array to ", index)
		entries = entries[0:index]
	}

	input := &sqs.SendMessageBatchInput{
		QueueUrl: &action.targetArn,
		Entries:  entries,
	}

	out, err := action.sqsService.SendMessageBatch(input)

	if err != nil {
		logs.Error.Println("We found a problem while sending the message: ", err.Error())

		errors[out] = ErrOnSQSPublishMessage
	}

	logs.Info.Printf("Send bulk message response: %+v \n", out)

	return errors
}
