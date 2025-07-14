package domain

type UserEventAccess struct {
	IsHoster    bool
	IsAttending bool
}

func (acces *UserEventAccess) HasAdminRole() bool {
	return acces.IsHoster
}

func (acces *UserEventAccess) HasChatAccess() bool {
	return acces.IsAttending || acces.IsHoster
}
