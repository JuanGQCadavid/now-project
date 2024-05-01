package utils

import (
	"net/http"

	"github.com/JuanGQCadavid/now-project/services/authService/internal/core/domain"
)

func GetHeaders(headers http.Header) *domain.UserDetails {

	var userName, userPhone, userId string = headers.Get(domain.USER_NAME_HEADER), headers.Get(domain.USER_PHONE_HEADER), headers.Get(domain.USER_ID_HEADER)

	if len(userId) == 0 {
		return domain.AnonymousUser
	}

	return &domain.UserDetails{
		UserID:      userId,
		Name:        userName,
		PhoneNumber: userPhone,
	}
}
