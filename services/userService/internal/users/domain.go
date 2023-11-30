package users

import "time"

type TableKey struct {
	PhoneNumber string `json:"PhoneNumber" dynamodbav:"PhoneNumber"`
}

type UserOTP struct {
	OTP      []int     `json:"otp"`
	TTL      time.Time `json:"ttl"`
	Attempts int       `json:"attempts"`
}

type UserValidate struct {
	Validated bool `json:"isValid" dynamodbav:"Validated"`
}

type UserOTPBody struct {
	OTP *UserOTP `json:"OTP" dynamodbav:"OTP"`
}
