package service

import (
	"github.com/JuanGQCadavid/now-project/services/confirmationService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/confirmationService/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
)

type ConfirmationService struct {
	repository ports.Repository
	notifyer   ports.Notify
}

func (srv *ConfirmationService) ConfirmDate(dateId string, userRequesterId string) error {
	return srv.changeDateStatus(dateId, userRequesterId, true, ports.DateConfirmed)
}
func (srv *ConfirmationService) UnconfirmDate(dateId string, userRequesterId string) error {
	return srv.changeDateStatus(dateId, userRequesterId, false, ports.DateUnconfirmed)
}

func (srv *ConfirmationService) changeDateStatus(dateId string, userRequesterId string, confirmed bool, spotActiviy ports.NotifyOperator) error {
	logs.Info.Printf("ConfirmDate: Date Id = %s, User Requets = %s \n", dateId, userRequesterId)

	date, err := srv.repository.FetchDate(dateId, userRequesterId)

	if err != nil {
		logs.Error.Printf("We face an error on the repository")
		return ports.ErrRepositoryFail
	}

	if date == nil {
		logs.Warning.Println("Empty date, could be wrong date id or invalid user requestId")
		return ports.ErrEmptyDate
	}

	if date.Confirmed {
		logs.Info.Printf("The date %s is already confirmed, aborting job.\n", date.Id)
		return nil
	}
	date.Confirmed = confirmed
	err = srv.repository.UpdateDateOnConfirmed(date.Id, date.Confirmed)

	if err != nil {
		logs.Error.Println("We could not update the date status due to a failure in the repository")
		return ports.ErrUpdatingDateOnRepository
	}

	logs.Info.Printf("Date confirmed updated sucessfully to %v \n", date.Confirmed)

	srv.notifyer.ConfirmationActivity(spotActiviy, domain.Notification{
		DateId:           date.Id,
		SpotId:           date.SpotId,
		UserId:           userRequesterId,
		Aditionalpayload: date,
	})

	return nil

}
