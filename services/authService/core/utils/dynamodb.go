package utils

import (
	"log"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func DynamoGetAndMapTo(keyName, keyValue, tableName string, mapTo any, svc *dynamodb.DynamoDB) error {
	key, err := dynamodbattribute.MarshalMap(map[string]interface{}{
		keyName: keyValue,
	})

	logs.Info.Printf("%+v\n", key)
	if err != nil {
		log.Fatalf("Got error marshalling key item: %s", err)
		return err
	}

	out, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       key,
	})

	if err != nil {
		logs.Error.Println("We fail to get the OTP, error: ", err.Error())
		return err
	}

	logs.Info.Printf("%+v\n", out)

	err = dynamodbattribute.UnmarshalMap(out.Item, mapTo)

	if err != nil {
		logs.Error.Println("We fail to unmarshal the OTP, error: ", err.Error())
		return err
	}

	return nil
}

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
		logs.Error.Println("We fail to get the OTP, error: ", err.Error())
		return err
	}

	logs.Info.Printf("%+v\n", out)

	log.Println("We spent -> ", out.ConsumedCapacity)

	if len(out.Items) == 0 {
		logs.Info.Println("Not data found")
		return nil
	}

	err = dynamodbattribute.UnmarshalMap(out.Items[0], mapTo)

	if err != nil {
		logs.Error.Println("We fail to unmarshal the OTP, error: ", err.Error())
		return err
	}

	return nil
}
