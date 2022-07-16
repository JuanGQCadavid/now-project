package sessionservice

import (
	"fmt"
	"log"
	"time"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain/session"
	"github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/sessionService/domain"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

type SearchSessionDynamoDbService struct {
	dynamoDbConnector *DynamoDbConnector
	defaultTTLSeconds int
}

func NewSearchSessionDynamoDbService() *SearchSessionDynamoDbService {
	return &SearchSessionDynamoDbService{
		dynamoDbConnector: NewConnectorFromEnv(),
		defaultTTLSeconds: 1800,
	}
}

func (svc *SearchSessionDynamoDbService) CreateSession(sessionType session.SessionTypes) (*session.SessionConfig, error) {
	sessionId := svc.newUUID()
	ttl := svc.generateTTL()
	log.Println(sessionId, ttl)

	item := domain.SessionItem{
		SessionId: sessionId,
		State:     string(sessionType),
		TTL:       ttl,
	}

	itemMarshaled, err := dynamodbattribute.MarshalMap(item)

	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}

	input := &dynamodb.PutItemInput{
		TableName: &svc.dynamoDbConnector.TableName,
		Item:      itemMarshaled,
	}
	input.SetReturnConsumedCapacity(dynamodb.ReturnConsumedCapacityTotal)

	output, outErr := svc.dynamoDbConnector.Svc.PutItem(input)

	if outErr != nil {
		log.Fatalf("Got error on the output of the put item : %s", outErr)
	}

	fmt.Printf("%+v\n", output)

	return &session.SessionConfig{
		TTL:         ttl,
		SessionType: sessionType,
		SessionId:   sessionId,
	}, nil
}

func (svc *SearchSessionDynamoDbService) GetSessionData(sessionId string, sessionType session.SessionTypes) (session.SessionData, error) {
	return session.SessionData{}, nil
}

func (svc *SearchSessionDynamoDbService) newUUID() string {
	return uuid.NewString()
}

func (svc *SearchSessionDynamoDbService) generateTTL() int64 {
	now := time.Now().Add(time.Second * time.Duration(svc.defaultTTLSeconds))
	return now.Unix()
}
