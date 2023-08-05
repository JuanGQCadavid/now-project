package ports

import "github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/core/domain"

type Service interface {
	OnDateCreation(domain.DatesLocation) error
	OnDateRemoved(string) error
	OnDateStatusChanged(string, domain.DateState) error
}

// onlineStart, dateConfirmed -> Create  on DB
// onlineResume onlineStop -> Change status
// onlineFinalize, dateUnconfirmed -> Remove from DB
