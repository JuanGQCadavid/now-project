package sessionservice

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDbConnector struct {
	Svc       *dynamodb.DynamoDB
	TableName string
}

func NewConnectorFromEnv() *DynamoDbConnector {
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	tableName, isPresentTableName := os.LookupEnv("sessionTableName")

	if !isPresentTableName {
		log.Fatalln("Missgin Table Name Env")
	}

	return &DynamoDbConnector{
		Svc:       dynamodb.New(session),
		TableName: tableName,
	}
}
