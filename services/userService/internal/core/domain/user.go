package domain

import "time"

type User struct {
	Name           string `json:"name" dynamodbav:"Name"`
	PhoneNumber    string `json:"phoneNumber" dynamodbav:"PhoneNumber"`
	Validated      bool   `json:"isValid" dynamodbav:"Validated"`
	ValidatedHash  string `json:"-" dynamodbav:"ValidatedHash"`
	UserId         string `json:"userId" dynamodbav:"UserId"`
	PhoneSignature string `json:"phoneSignature" dynamodbav:"PhoneSignature,omitempty"`
}

type UserProfile struct {
	UserName string `json:"user_name,omitempty" dynamodbav:"Name"`

	// Per user atribute there should be a flag for specifying wheter it is public
	FirstName             string `json:"first_name,omitempty" dynamodbav:"FirstName"`
	LastName              string `json:"last_name,omitempty" dynamodbav:"LastName"`
	IsFirstLastNamePublic bool   `json:"is_first_last_name_public,omitempty" dynamodbav:"IsFirstLastNamePublic"`

	// Phone number
	PhoneNumber         string `json:"phone_number,omitempty" dynamodbav:"PhoneNumber"`
	IsPhoneNumberPublic bool   `json:"is_phone_number_public,omitempty" dynamodbav:"IsPhoneNumberPublic"`
}

func (up *UserProfile) CleanSensitiveData() {
	if !up.IsFirstLastNamePublic {
		up.FirstName = ""
		up.LastName = ""
	}
	if !up.IsPhoneNumberPublic {
		up.PhoneNumber = ""
	}
}

type Tokens struct {
	TokenId           string    `json:"tokenId" dynamodbav:"TokenId"`
	UserID            string    `json:"userId" dynamodbav:"UserID"`
	RefreshToken      string    `json:"refreshToken" dynamodbav:"RefreshToken"`
	ShortLiveToken    string    `json:"shortLiveToken" dynamodbav:"ShortLiveToken"`
	ShortLiveTokenTTL time.Time `json:"shortLiveTokenTTL" dynamodbav:"ShortLiveTokenTTL"`
	JWT               string    `json:"jwt"`
}

type Login struct {
	PhoneNumber       string           `json:"phoneNumber"`
	PhoneSignature    string           `json:"phoneSignature"`
	MethodVerificator MethodVerifictor `json:"methodVerificator"`
}

type SingUp struct {
	PhoneNumber       string           `json:"phoneNumber"`
	UserName          string           `json:"userName"`
	PhoneSignature    string           `json:"phoneSignature"`
	MethodVerificator MethodVerifictor `json:"methodVerificator"`
}

type ValidateProcess struct {
	PhoneNumber string `json:"phoneNumber"`
	Code        []int  `json:"code"`
}

type MethodVerifictor struct {
	Language string `json:"language"`
	WhatsApp bool   `json:"whatsapp"`
	SMS      bool   `json:"sms"`
}

type NotificatorType string

const (
	SMS      NotificatorType = "SMS"
	WHATSAPP NotificatorType = "WHATSAPP"
	DEFAULT  NotificatorType = "DEFAULT"
)
