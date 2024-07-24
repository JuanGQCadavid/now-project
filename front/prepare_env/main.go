package main

import (
	"bytes"
	_ "embed"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"text/template"

	// "encoding/base64"
	// "fmt"
	// "strings"
	"time"

	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

type User struct {
	Name           string `json:"name" dynamodbav:"Name"`
	PhoneNumber    string `json:"phoneNumber" dynamodbav:"PhoneNumber"`
	Validated      bool   `json:"isValid" dynamodbav:"Validated"`
	UserId         string `json:"userId" dynamodbav:"UserId"`
	PhoneSignature string `json:"phoneSignature" dynamodbav:"PhoneSignature,omitempty"`
}

type Tokens struct {
	TokenId           string    `json:"tokenId" dynamodbav:"TokenId"`
	UserID            string    `json:"userId" dynamodbav:"UserID"`
	RefreshToken      string    `json:"refreshToken" dynamodbav:"RefreshToken"`
	ShortLiveToken    string    `json:"shortLiveToken" dynamodbav:"ShortLiveToken"`
	ShortLiveTokenTTL time.Time `json:"shortLiveTokenTTL" dynamodbav:"ShortLiveTokenTTL"`
}

type AllUserRelate struct {
	UserStuff   User
	TokensStuff Tokens
}

type Spot struct {
	Id  string `json:"id"`
	Lat float64
	Lng float64
}

type CreateSpotResponse struct {
	EventInfo Spot `json:"eventInfo"`
}

const (
	usersCount        = 1
	UserTableName     = "Users-staging"
	TokensTableName   = "Tokens-staging"
	BaseLat           = 6.250467
	MaxJump           = 0.05
	BaseLng           = -75.593271
	StandarOnlineBody = "{\"durationApproximated\": 300,\"maximunCapacity\": 58}"
	ScheduleOne       = `{
		\"patterns\": [
			{
				\"host": {
					\"hostId\": \"%s\"
				},
				\"day\": 127,
				\"fromDate\": \"2023-07-10\",
				\"toDate\": \"2024-10-23\",
				\"StartTime\": \"%s\",
				\"endTime\": \"%s\"
			}
		]
	}`
)

var (
	svc *dynamodb.DynamoDB

	//go:embed new_spot.json
	createTemplate string

	// hours []string = []string{"00", "", ""}
)

func main() {

	// // 1 Create users on dynamo
	// usersList := createUsers(usersCount)

	// for _, userStuff := range usersList {
	// 	insertDynamo(userStuff.UserStuff, UserTableName)
	// 	insertDynamo(userStuff.TokensStuff, TokensTableName)
	// }
	// // 2 Create two online

	// for index, user := range usersList {
	// 	spotCreate := createSpotOnService(index, user)
	// 	goOnline(spotCreate, user)
	// }

	// startTime := time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC)
	startTime, _ := time.Parse(time.DateTime, "2006-01-01 00:00:00")
	for {
		log.Println(startTime.Date())
		println(startTime.Format(time.TimeOnly))

		log.Printf(ScheduleOne+"\n", "JOST", startTime.Format(time.TimeOnly), startTime.Add(time.Hour*2).Format(time.TimeOnly))
		startTime = startTime.Add(time.Hour * 2)

		if strings.Contains(startTime.Format(time.TimeOnly), "22") {
			break
		}
	}

	// 3 create two schedule
}

func goOnline(spotCreate *CreateSpotResponse, user AllUserRelate) {
	client := &http.Client{}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("https://staging.pululapp.com/spots/online/%s/start", spotCreate.EventInfo.Id),
		strings.NewReader(StandarOnlineBody),
	)

	if err != nil {
		log.Fatalln("Err on online request 1", err.Error())
	}

	req.Header.Add("X-Auth", user.TokensStuff.ShortLiveToken)
	req.Header.Add("Content-Type", "application/json")

	if resp, err := client.Do(req); err != nil {
		log.Fatalln("Err on online request 2", err.Error())
	} else {
		log.Println("Online response: ", resp.Status)
	}

}

func createSpotOnService(index int, user AllUserRelate) *CreateSpotResponse {
	client := &http.Client{}

	results, _ := populateTemplate(generateSpot(index), createTemplate)
	req, err := http.NewRequest("POST", "https://staging.pululapp.com/spots/core/", results)

	if err != nil {
		log.Fatalln(err.Error())
	}
	req.Header.Add("X-Auth", user.TokensStuff.ShortLiveToken)
	req.Header.Add("Content-Type", "application/json")

	if resp, err := client.Do(req); err != nil {
		log.Fatalln(err.Error())
	} else {
		respBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		respString := string(respBytes)
		log.Println(respString)

		var result CreateSpotResponse
		json.Unmarshal(respBytes, &result)

		log.Println(result.EventInfo.Id)

		return &result
	}

	return nil
}

func generateSpot(index int) Spot {
	r := rand.Float64() * (MaxJump)
	id := strings.Split(generateId(), "-")[1]

	if index%2 == 0 {
		log.Println("Inverting R")
		r = r * -1
	}

	return Spot{
		Id:  id,
		Lat: BaseLat + r,
		Lng: BaseLng + r,
	}
}

func createUsers(usersCount int) []AllUserRelate {
	usersList := make([]AllUserRelate, usersCount, usersCount)

	for index := range usersList {
		userId := generateId()
		tokenId := generateId()
		usersList[index] = AllUserRelate{
			UserStuff: User{
				Name:        fmt.Sprintf("User-%s", strings.Split(userId, "-")[0]),
				UserId:      userId,
				Validated:   true,
				PhoneNumber: fmt.Sprintf("+57999999+%s", strings.Split(userId, "-")[0]),
			},
			TokensStuff: Tokens{
				TokenId:           tokenId,
				UserID:            userId,
				RefreshToken:      base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s+%s", tokenId, generateId()))),
				ShortLiveToken:    base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s+%s", tokenId, generateId()))),
				ShortLiveTokenTTL: time.Now(),
			},
		}
	}

	return usersList
}

func populateTemplate(data any, tmplParser string) (*bytes.Buffer, error) {
	id := uuid.NewString()

	tmpl, err := template.New(fmt.Sprintf("template/%s.tmpl", id)).Parse(tmplParser)
	if err != nil {
		log.Println("error while creating the template ", err.Error())
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, data)

	if err != nil {
		log.Println("error Executing the template ", err.Error())
		return nil, err
	}

	return buf, nil
}

func generateId() string {
	return uuid.NewString()
}

func insertDynamo(stuff any, onTable string) {
	if userData, err := dynamodbattribute.MarshalMap(stuff); err != nil {
		log.Fatalf("Got error marshalling item: %s", err)
	} else {
		if _, err := svc.PutItem(
			&dynamodb.PutItemInput{
				Item:      userData,
				TableName: aws.String(onTable),
			}); err != nil {
			log.Fatalf("Got error calling PutItem: %s", err)
		}
	}

}

func init() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc = dynamodb.New(sess)
}
