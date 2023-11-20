package domain

type User struct {
	Name           string `json:"name" dynamodbav:"Name"`
	PhoneNumber    string `json:"phoneNumber" dynamodbav:"PhoneNumber"`
	Validated      bool   `json:"isValid" dynamodbav:"Validated"`
	UserId         string `json:"userId" dynamodbav:"UserId"`
	PhoneSignature string `json:"phoneSignature" dynamodbav:"PhoneSignature,omitempty"`
}

type Tokens struct {
	TokenId              string
	UserId               string
	LongLiveRefreshToken string
	ShortLiveToken       string
	ShortLiveTokenTTL    int
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
