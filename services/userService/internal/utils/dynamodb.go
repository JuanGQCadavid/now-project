package utils

import (
	"errors"
	"log"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var (
	ErrDynamoDBMarshaling = errors.New("err while marshaling to dynamodb")
	ErrDynamoDBInserting  = errors.New("err while inserting to dynamodb")
)

func DynamoQueryOneAndMapTo(keyName, keyValue, tableName, index string, mapTo any, svc *dynamodb.DynamoDB) error {
	key, err := dynamodbattribute.MarshalMap(map[string]interface{}{
		keyName: keyValue,
	})

	logs.Info.Printf("%+v\n", key)
	if err != nil {
		log.Fatalf("Got error marshalling key item: %s", err)
		return err
	}

	out, err := svc.Query(&dynamodb.QueryInput{
		TableName: aws.String(tableName),
		IndexName: &index,
		KeyConditions: map[string]*dynamodb.Condition{
			keyName: {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(keyValue),
					},
				},
			},
		},
	})

	if err != nil {
		logs.Error.Println("We fail to get data, error: ", err.Error())
		return err
	}

	logs.Info.Printf("%+v\n", out)
	err = dynamodbattribute.UnmarshalMap(out.Items[0], mapTo)
	if err != nil {
		logs.Error.Println("We fail to unmarshal the data, error: ", err.Error())
		return err
	}
	return nil
}

func DynamoGetAndMapTo(keyName, keyValue, tableName string, mapTo any, svc *dynamodb.DynamoDB) error {
	key, err := dynamodbattribute.MarshalMap(map[string]interface{}{
		keyName: keyValue,
	})

	logs.Info.Printf("%+v\n", key)
	if err != nil {
		logs.Error.Printf("Got error marshalling key item: %s\n", err)
		return err
	}

	out, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       key,
	})

	if err != nil {
		logs.Error.Println("We fail to get the item, error: ", err.Error())
		return err
	}

	logs.Info.Printf("%+v\n", out)

	err = dynamodbattribute.UnmarshalMap(out.Item, mapTo)

	if err != nil {
		logs.Error.Println("We fail to unmarshal the item, error: ", err.Error())
		return err
	}

	return nil
}

func UpdateItem(item any, tableName string, svc *dynamodb.DynamoDB) error {

	marshaled, err := dynamodbattribute.MarshalMap(item)

	if err != nil {
		logs.Error.Println("err we fail marshaling the data, err: ", err.Error())
		return errors.Join(
			ErrDynamoDBMarshaling,
			err,
		)
	}

	userInput := &dynamodb.PutItemInput{
		Item:                   marshaled,
		TableName:              &tableName,
		ReturnConsumedCapacity: aws.String(dynamodb.ReturnConsumedCapacityTotal),
	}

	_, err = svc.PutItem(userInput)

	if err != nil {
		logs.Error.Println("err we fail putting the data, err", err.Error())
		return errors.Join(
			ErrDynamoDBInserting,
			err,
		)
	}

	return nil
}
