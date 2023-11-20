package users

type OTP struct {
	PhoneNumber string `json:"PhoneNumber" dynamodbav:"PhoneNumber"`
}
