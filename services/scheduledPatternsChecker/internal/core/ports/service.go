package ports

import "errors"

var (
	ErrOnRepository        = errors.New("err on repository")
	ErrInvalidCors         = errors.New("The cores number is 0 or negative")
	ErrProcessingDates     = errors.New("We got an error while pocessing dates")
	ErrSendingConfirmation = errors.New("Confirmation service fail to send")
)

type Service interface {
}
