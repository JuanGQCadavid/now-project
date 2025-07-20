package spotscore

import (
	"context"
	"encoding/json"
	"io"
	"strings"

	"github.com/JuanGQCadavid/now-project/services/fileService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/fileService/internal/core/ports"
	"github.com/rs/zerolog/log"
	"resty.dev/v3"
)

const (
	getUserEventAccessURI = "{{PULULAPP_DNS}}/spots/core/access/{{EVENT_ID}}/{{USER_ID}}"
	getUserDateAccessURI  = "{{PULULAPP_DNS}}/spots/core/access/{{EVENT_ID}}/{{USER_ID}}?dateId={{DATE_ID}}"
)

type SpotsCoreService struct {
	client *resty.Client
	dns    string
}

func NewSpotsCoreService(serviceDns string) *SpotsCoreService {
	return &SpotsCoreService{
		client: resty.New(),
		dns:    serviceDns,
	}
}

func (srv *SpotsCoreService) GetUserEventAccess(ctx context.Context, userID string, eventId string) (*domain.UserEventAccess, error) {

	var (
		bytes, err = srv.performGet(ctx, getUserEventAccessURI, map[string]string{
			"{{PULULAPP_DNS}}": srv.dns,
			"{{EVENT_ID}}":     eventId,
			"{{USER_ID}}":      userID,
		})
		logger          = log.Ctx(ctx)
		payload *Access = &Access{}
	)

	if err != nil {
		logger.Err(err).Str("method", "GetUserEventAccess").Send()
		return nil, err
	}

	if err := json.Unmarshal(bytes, payload); err != nil {
		logger.Err(err).Msg("Fail on casting payload")
		return nil, ports.ErrCallingBackend
	}

	return fomAccessToUserEventAcess(payload), nil
}

func (srv *SpotsCoreService) GetUserDateAccess(ctx context.Context, eventId, userID, dateId string) (*domain.UserEventAccess, error) {
	var (
		bytes, err = srv.performGet(ctx, getUserDateAccessURI, map[string]string{
			"{{PULULAPP_DNS}}": srv.dns,
			"{{EVENT_ID}}":     eventId,
			"{{USER_ID}}":      userID,
			"{{DATE_ID}}":      dateId,
		})
		logger          = log.Ctx(ctx)
		payload *Access = &Access{}
	)

	if err != nil {
		logger.Err(err).Str("method", "GetUserEventAccess").Send()
		return nil, err
	}

	if err := json.Unmarshal(bytes, payload); err != nil {
		logger.Err(err).Msg("Fail on casting payload")
		return nil, ports.ErrCallingBackend
	}

	return fomAccessToUserEventAcess(payload), nil
}

func (srv *SpotsCoreService) castURL(url string, replacements map[string]string) string {
	var (
		results = url
	)
	for key, val := range replacements {
		results = strings.ReplaceAll(results, key, val)
	}
	return results
}

func (srv *SpotsCoreService) performGet(ctx context.Context, url string, replacements map[string]string) ([]byte, error) {
	var (
		uri    = srv.castURL(getUserEventAccessURI, replacements)
		logger = log.Ctx(ctx)
	)

	logger.Debug().Str("URI", uri).Send()

	res, err := srv.client.R().
		EnableTrace().
		SetHeader("X-Auth", "ANONY"). // We should get the token from the ctx from the user token
		Get(uri)

	if err != nil {
		logger.Err(err).Str("URI", uri).Send()
		return nil, ports.ErrCallingBackend
	}

	bytes, err := io.ReadAll(res.Body)

	if err != nil {
		logger.Err(err).Str("URI", uri).Str("method", "Reading payload").Send()
		return nil, ports.ErrCallingBackend

	}

	logger.Info().Any("StatusCode", res.StatusCode()).Str("body", string(bytes)).Any("Trace", res.Request.TraceInfo()).Send()

	return bytes, nil

}
