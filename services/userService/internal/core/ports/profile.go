package ports

import "github.com/JuanGQCadavid/now-project/services/userService/internal/core/domain"

type ProfileRepository interface {
	CreateProfile(*domain.UserProfile) error
	UpdateProfile(*domain.UserProfile) error

	// Fetch User profile from repository
	// Returns:
	//	- ErrUserDoesNotExist
	// 	- UserProfile
	GetUserProfile(userId string) (*domain.UserProfile, error)
}
