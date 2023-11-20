package domain

type User struct {
	Name           string `json:"name"`
	PhoneNumber    string `json:"phoneNumber"`
	Validated      bool   `json:"isValid"`
	UserId         string `json:"userId"`
	PhoneSignature string `json:"phoneSignature"`
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
