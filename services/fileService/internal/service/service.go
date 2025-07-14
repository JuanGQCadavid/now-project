package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	authDomain "github.com/JuanGQCadavid/now-project/services/authService/core/core/domain"
	"github.com/JuanGQCadavid/now-project/services/fileService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/fileService/internal/core/ports"
)

type FileSerice struct {
	objectRepository    ports.ObjectRepository
	spotsCoreRepository ports.SpotsCoreRepository
}

func NewFileSerice(
	objectRepository ports.ObjectRepository,
	spotsCoreRepository ports.SpotsCoreRepository,
) *FileSerice {
	return &FileSerice{
		objectRepository:    objectRepository,
		spotsCoreRepository: spotsCoreRepository,
	}
}

const (
	PROFILE_PATH = "profiles"
)

var (
	ErrMissingAuthentication               error = fmt.Errorf("err user should be log in")
	ErrScopeNotDefined                     error = errors.New("err undefined scope")
	ErrInternalError                       error = errors.New("err internal error")
	ErrNoChatAcess                         error = errors.New("err you down belong to the chat")
	ErrIsNotAHoster                        error = errors.New("err the user does not have the right, they is not a hoster")
	ErrMissingControlAccessDataForTheScope error = errors.New("err scope is defined but it is not reflected in the controlAccess")
)

func (fs *FileSerice) UploadFile(ctx context.Context, userDetails authDomain.UserDetails, fileMetadata *domain.FileMetadata) (*domain.PresignedURL, error) {
	var (
		filePath      string
		loggerr                             = log.Ctx(ctx)
		controlAccess *domain.ControlAccess = &domain.ControlAccess{}
	)

	if len(userDetails.UserID) == 0 || userDetails.UserID == authDomain.AnonymousUser.UserID {
		loggerr.Warn().Str("UserId", userDetails.UserID).Msg("User is not authenticated")
		return nil, ErrMissingAuthentication
	}

	switch fileMetadata.Scope {
	case domain.ProfileScope:
		// users_files/:userId/profile/profile_picture.jpg
		filePath = fmt.Sprintf("users_files/%s/profile/profile_picture.%s", userDetails.UserID, fileMetadata.Type)
		controlAccess = nil
	case domain.ChatScope:
		// TODO - So far chat Id is date Id.

		// users_files/:userId/chats/:chatId/uuid.[type]
		if len(fileMetadata.ControlAccess.ChatId) == 0 {
			loggerr.Warn().Any("ControlAccess", fileMetadata.ControlAccess).Msg("Control access missing ChatId for ChatScope")
			return nil, ErrMissingControlAccessDataForTheScope
		}
		access, err := fs.spotsCoreRepository.GetUserDateAccess(userDetails.UserID, fileMetadata.ControlAccess.EventId)

		if err != nil {
			loggerr.Err(err).
				Str("UserID", userDetails.UserID).
				Str("ChatId", fileMetadata.ControlAccess.ChatId).
				Msg("Err while calling GetUserDateAccess")

			return nil, ErrInternalError
		}

		if !access.HasChatAccess() {
			return nil, ErrNoChatAcess
		}

		filePath = fmt.Sprintf("users_files/%s/chats/%s/%s.%s",
			userDetails.UserID,
			fileMetadata.ControlAccess.ChatId,
			uuid.NewString(),
			fileMetadata.Type,
		)

		controlAccess.ChatId = fileMetadata.ControlAccess.ChatId
	case domain.DateScope, domain.EventScope:

		if len(fileMetadata.ControlAccess.EventId) == 0 {
			loggerr.Warn().Any("ControlAccess", fileMetadata.ControlAccess).Msg("Control access missing EventId for Date or event scope")
			return nil, ErrMissingControlAccessDataForTheScope
		}

		// TODO - check that the user admins the date or the event
		access, err := fs.spotsCoreRepository.GetUserEventAccess(userDetails.UserID, fileMetadata.ControlAccess.EventId)

		if err != nil {
			loggerr.Err(err).
				Str("UserID", userDetails.UserID).
				Str("EventID", fileMetadata.ControlAccess.EventId).
				Msg("Err while calling GetUserEventAccess")

			return nil, ErrInternalError
		}

		if !access.HasAdminRole() {
			loggerr.Warn().
				Str("UserID", userDetails.UserID).
				Str("EventID", fileMetadata.ControlAccess.EventId).
				Msg("User is not admin")
			return nil, ErrIsNotAHoster
		}

		// users_files/:userId/events/:eventId/dates/:dateId/uuid.[type]
		var (
			subFilePath = fmt.Sprintf("users_files/%s/events/%s", userDetails.UserID, fileMetadata.ControlAccess.EventId)
			objectId    = uuid.NewString()
		)

		// users_files/:userId/events/:eventId/dates/:dateId/uuid.[type]
		if fileMetadata.Scope == domain.DateScope {
			if len(fileMetadata.ControlAccess.DateId) == 0 {
				loggerr.Warn().Any("ControlAccess", fileMetadata.ControlAccess).Msg("Control access missing Date for Date scope")
				return nil, ErrMissingControlAccessDataForTheScope
			}
			filePath = fmt.Sprintf("%s/dates/%s/%s.%s", subFilePath, fileMetadata.ControlAccess.DateId, objectId, fileMetadata.Type)
			controlAccess.DateId = fileMetadata.ControlAccess.DateId
			controlAccess.EventId = fileMetadata.ControlAccess.EventId
			break
		}
		// users_files/:userId/events/:eventId/uuid.[type]
		filePath = fmt.Sprintf("%s/%s.%s", subFilePath, objectId, fileMetadata.Type)
		controlAccess.EventId = fileMetadata.ControlAccess.EventId

	default:
		return nil, ErrScopeNotDefined
	}

	return fs.objectRepository.GeneratePresignedURL(filePath, controlAccess)
}
