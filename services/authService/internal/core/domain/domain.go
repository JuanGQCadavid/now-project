package domain

import "time"

const (
	APP_TOKEN         = "x-auth"
	ANONYMOUS_KEY     = "ANONYMOUS"
	USER_NAME_HEADER  = "x-internal-user-name"
	USER_PHONE_HEADER = "x-internal-user-phone-number"
	USER_ID_HEADER    = "x-internal-user-id"
)

var (
	AnonymousUser *UserDetails = &UserDetails{
		UserID:      ANONYMOUS_KEY,
		Name:        ANONYMOUS_KEY,
		PhoneNumber: ANONYMOUS_KEY,
	}
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

type UserDetails struct {
	UserID      string `json:"UserID"`
	Name        string `json:"Name"`
	PhoneNumber string `json:"PhoneNumber"`
}

type Token struct {
	TokenID    string
	TokenValue string
}
