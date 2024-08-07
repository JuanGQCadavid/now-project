package ports

import "errors"

var (
	ErrRepositoryFail           = errors.New("we face an error while accessing the data")
	ErrEmptyDate                = errors.New("empty date, could be wrong date id ")
	ErrUserIsNotTheHost         = errors.New("User is not alowed to perform such operation")
	ErrDateIsOnline             = errors.New("You could not confirm/unconfirm a online date")
	ErrUpdatingDateOnRepository = errors.New("We could not update the date status due to a failure in the repository")
)

type Service interface {
	ConfirmDate(dateId string, userRequesterId string) error
	UnconfirmDate(dateId string, userRequesterId string) error
}
