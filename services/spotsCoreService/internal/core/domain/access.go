package domain

type AccessRole string

const (
	AttendantAccessRole AccessRole = "Attendant"
	HostAccessRole      AccessRole = "Host"
	NoAccessRole        AccessRole = "NoAccess"
)

type Access struct {
	UserName string
	UserId   string
	Role     AccessRole
}
