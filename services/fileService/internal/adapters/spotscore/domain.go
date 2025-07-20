package spotscore

import "github.com/JuanGQCadavid/now-project/services/fileService/internal/core/domain"

type AccessRole string

const (
	AttendantAccessRole AccessRole = "Attendant"
	HostAccessRole      AccessRole = "Host"
	NoAccessRole        AccessRole = "NoAccess"
)

type Access struct {
	UserName string     `json:"userName,omitempty"`
	UserId   string     `json:"userId,omitempty"`
	Role     AccessRole `json:"role,omitempty"`
}

func fomAccessToUserEventAcess(payload *Access) *domain.UserEventAccess {
	switch payload.Role {
	case HostAccessRole:
		return &domain.UserEventAccess{
			IsHoster: true,
		}
	case AttendantAccessRole:
		return &domain.UserEventAccess{
			IsAttending: true,
		}
	}

	return &domain.UserEventAccess{
		IsHoster:    false,
		IsAttending: false,
	}
}
