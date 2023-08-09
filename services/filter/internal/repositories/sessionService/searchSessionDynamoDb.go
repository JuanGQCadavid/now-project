package sessionservice

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain/session"
	"github.com/JuanGQCadavid/now-project/services/filter/internal/repositories/sessionService/domain"
	"github.com/aws/aws-sdk-go/aws"
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
	log.Println("Testing on Search session handler -------------------------")
	log.Println(fmt.Sprintf("CreateSession -> Creating session with sessionType: %s", string(sessionType)))
	sessionId := svc.newUUID()
	ttl := svc.generateTTL()
	log.Println(sessionId, ttl)

	item := domain.SessionItem{
		SessionId: sessionId,
		State:     string(sessionType),
		TTL:       ttl,
	}
	log.Println(fmt.Sprintf("CreateSession -> Item %+v", item))

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
		// TODO -> We should handle this error
		log.Fatalf("Got error on the output of the put item : %s", outErr)
	}

	fmt.Printf("CreateSession -> Goes well, result: %+v\n", output)

	return &session.SessionConfig{
		TTL:         ttl,
		SessionType: sessionType,
		SessionId:   sessionId,
	}, nil
}

func (svc *SearchSessionDynamoDbService) GetSessionData(sessionId string, sessionType session.SessionTypes) (session.SearchSessionData, error) {
	log.Println(fmt.Sprintf("GetSessionData -> Getting session with sessionType: %s, sessionId: %s", string(sessionType), sessionId))
	key, err := dynamodbattribute.MarshalMap(domain.SessionItem{SessionId: sessionId, State: string(sessionType)})

	if err != nil {
		log.Fatalf("Got error marshalling get item: %s", err)
	}

	input := &dynamodb.GetItemInput{
		TableName: &svc.dynamoDbConnector.TableName,
		Key:       key,
	}
	output, err := svc.dynamoDbConnector.Svc.GetItem(input)

	if err != nil {
		log.Fatalf("Got error geting item: %s", err)
	}

	log.Println(fmt.Sprintf("Get operation goes well, Output:%+v", output))

	if output.Item == nil {
		log.Println("Session Not founded")
		return session.SearchSessionData{}, errors.New("Session not founded")
	}

	item := output.Item
	sessionData := session.SearchSessionData{
		SessionData: session.SessionData{
			SessionConfiguration: session.SessionConfig{},
		},
		Spots: make(map[string][]string),
	}

	for key := range item {

		if key == "TTL" {
			log.Println("Key: ", key, " Value: ", item[key])
			ttl := ""
			dynamodbattribute.Unmarshal(item[key], &ttl)

			log.Println("TTL: ", ttl)
			convertedValue, err := strconv.ParseInt(ttl, 10, 64)

			if err != nil {
				log.Println("TTL bad formated, asigning 0")
				convertedValue = 0
			}
			sessionData.SessionConfiguration.TTL = convertedValue
		} else if key == "State" {
			state := ""
			svc.getString(item, &state, key)

			log.Println(state)

			if strings.ToUpper(state) == strings.ToUpper(string(session.SpotsReturned)) {
				sessionData.SessionConfiguration.SessionType = session.SpotsReturned
			} else {
				log.Println("Empty State, assigning Empty")
				sessionData.SessionConfiguration.SessionType = session.Empty
			}

		} else if key == "SessionId" {
			sessionData.SessionConfiguration.SessionId = sessionId
		} else {
			log.Println("Key: ", key, " Value: ", item[key])

			var sessionItems []string = make([]string, 15)
			dynamodbattribute.Unmarshal(item[key], &sessionItems)
			log.Println(fmt.Sprintf("%+v", sessionItems))

			sessionData.Spots[key] = sessionItems
		}
	}

	log.Println("Result: ", fmt.Sprintf("%+v", sessionData))
	return sessionData, nil
}

func (svc *SearchSessionDynamoDbService) getString(source map[string]*dynamodb.AttributeValue, destinatio *string, key string) {
	log.Println("Key: ", key, " Source: ", source[key])
	dynamodbattribute.Unmarshal(source[key], destinatio)
	log.Println("Value: ", *destinatio)
}

func (svc *SearchSessionDynamoDbService) AddSpotsToSession(sessionId string, sessionType session.SessionTypes, spots []string) error {
	log.Println(fmt.Sprintf("AddSpotsToSession -> Adding spots session with sessionType: %s, sessionId: %s, Spots: %+v", string(sessionType), sessionId, spots))

	if len(spots) == 0 {
		log.Println("Empty spots, avoiding process")
		return nil
	}
	key, err := dynamodbattribute.MarshalMap(domain.SessionItem{SessionId: sessionId, State: string(sessionType)})

	if err != nil {
		log.Fatalf("Got error marshalling key item: %s", err)
	}

	log.Println(fmt.Sprintf("AddSpotsToSession -> key %+v", key))

	expressionAttributeValues := map[string]*dynamodb.AttributeValue{}
	updateExpression := "set"
	expressionAttributesNames := map[string]*string{}

	var startIndex, cutIndex, count, slope, maxIndex int = 0, 0, 0, 15, len(spots)
	var stopFlag bool = false

	for !stopFlag {

		cutIndex = startIndex + slope

		if cutIndex >= maxIndex {
			cutIndex = maxIndex
			stopFlag = true
		}

		identifier := fmt.Sprintf("spots%d", count)
		attributeNameValue := time.Now().Format(time.RFC3339Nano)
		updateExpression = fmt.Sprintf("%s #%s = :%s,", updateExpression, identifier, identifier)
		expressionAttributesNames[fmt.Sprintf("#%s", identifier)] = aws.String(attributeNameValue)
		expressionAttributeValues[fmt.Sprintf(":%s", identifier)] = &dynamodb.AttributeValue{
			SS: aws.StringSlice(spots[startIndex:cutIndex]),
		}

		startIndex = cutIndex
		count = count + 1
	}

	// Remove last " , " in the updateExpression
	updateExpression = strings.TrimSuffix(updateExpression, ",")

	input := &dynamodb.UpdateItemInput{
		Key:                       key,
		TableName:                 &svc.dynamoDbConnector.TableName,
		ExpressionAttributeValues: expressionAttributeValues,
		ExpressionAttributeNames:  expressionAttributesNames,
		ReturnConsumedCapacity:    aws.String(dynamodb.ReturnConsumedCapacityIndexes),
		UpdateExpression:          &updateExpression,
	}

	_, err = svc.dynamoDbConnector.Svc.UpdateItem(input)

	return err
}

func (svc *SearchSessionDynamoDbService) newUUID() string {
	return uuid.NewString()
}

func (svc *SearchSessionDynamoDbService) generateTTL() int64 {
	now := time.Now().Add(time.Second * time.Duration(svc.defaultTTLSeconds))
	return now.Unix()
}
