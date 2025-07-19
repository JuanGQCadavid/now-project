package domain

type AccessRole string

func AccessRoleFromString(role string) AccessRole {
	switch role {
	case "HOST":
		return HostAccessRole
	case "Attendant":
		return AttendantAccessRole
	default:
		return NoAccessRole
	}
}

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
