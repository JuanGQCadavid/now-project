package users

import "time"

type TableKey struct {
	PhoneNumber string `json:"PhoneNumber" dynamodbav:"PhoneNumber"`
}

type UserOTP struct {
	OTP      []int         `json:"otp"`
	TTL      time.Duration `json:"ttl"`
	Attempts int           `json:"attempts"`
}
