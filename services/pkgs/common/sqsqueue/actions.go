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

func (action *SQSQueueActions) SendBulkMessages(payload []interface{}) error {

	var startPointer int
	var threshold int

	if len(payload) == 0 {
		logs.Info.Println("Empty payload, aborting job")
		return nil
	}

	for {
		threshold = startPointer + action.bulkJumper
		var err error = nil
		if threshold < len(payload) {
			err = action.sendBuklMessage(payload[startPointer:threshold])
			startPointer = threshold
		} else if (len(payload)-startPointer) > 0 && (len(payload)-startPointer) <= action.bulkJumper {
			err = action.sendBuklMessage(payload[startPointer:])
			startPointer = len(payload)
		} else {
			break
		}

		// What should we do here ?
		if err != nil {
			logs.Error.Println("We got an error while processing the bulk operation ")
			return err
		}
	}

	logs.Info.Println("All messages where sent successfully")

	return nil
}

// How should we handle this kind of errors ?
func (action *SQSQueueActions) sendBuklMessage(payload []interface{}) error {
	entries := make([]*sqs.SendMessageBatchRequestEntry, len(payload))
	date := strings.ReplaceAll(time.Now().Format(time.DateTime), " ", "_")
	date = strings.ReplaceAll(date, ":", "_")
	for index, entry := range payload {
		bodyMarshaled, err := json.Marshal(&entry)

		if err != nil {
			return ErrParsingMessage
		}

		id := fmt.Sprintf("%d_%s", index, date)
		logs.Info.Println("Id: ", id)
		entries[index] = &sqs.SendMessageBatchRequestEntry{
			Id:          aws.String(id),
			MessageBody: aws.String(string(bodyMarshaled)),
		}

	}

	input := &sqs.SendMessageBatchInput{
		QueueUrl: &action.targetArn,
		Entries:  entries,
	}

	out, err := action.sqsService.SendMessageBatch(input)

	if err != nil {
		logs.Error.Println("We found a problem while sending the message: ", err.Error())
		return ErrOnSQSPublishMessage
	}

	logs.Info.Printf("Send bulk message response: %+v \n", out)

	return nil
}
