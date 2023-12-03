package httphdl

type ErrorMessageId string

const (
	UserNotFound            ErrorMessageId = "UserNotFound"
	Internal                ErrorMessageId = "Internal"
	BadBodyRequest          ErrorMessageId = "BadBodyRequest"
	PhoneNumberAlreadyTaken ErrorMessageId = "PhoneNumberAlreadyTaken"
	OTPAlive                ErrorMessageId = "OTPAlive"
	OTPDied                 ErrorMessageId = "OTPDied"
	WrongOTP                ErrorMessageId = "WrongOTP"
	NoPendingOTP            ErrorMessageId = "NoPendingOTP"
	OTPMaxTriesReached      ErrorMessageId = "OTPMaxTriesReached"
)

type ErrorMessage struct {
	Id            ErrorMessageId `json:"id,omitempty"`
	Message       string         `json:"message,omitempty"`
	InternalError string         `json:"internalError,omitempty"`
}
