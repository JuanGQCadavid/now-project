package httphdl

import (
	"fmt"
	"log"
	"strconv"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain/session"
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/filter/internal/handlers/httphdl/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HTTPHandler struct {
	service        ports.FilterService
	session        ports.SearchSessionService
	defaultRadious float64
}

func NewHTTPHandler(service ports.FilterService, session ports.SearchSessionService) *HTTPHandler {
	return &HTTPHandler{
		service:        service,
		session:        session,
		defaultRadious: 0.5,
	}
}

func (hdl *HTTPHandler) FilterSpots(context *gin.Context) {
	queryLat, isLatPresent := context.GetQuery("cpLat")
	queryLon, isLonPresent := context.GetQuery("cpLon")
	queryRadious := context.DefaultQuery("radious", fmt.Sprintf("%f", hdl.defaultRadious))
	queryCreateSession := context.DefaultQuery("createSession", "false")
	headerSessionUUID := context.Request.Header.Get("X-Now-Search-Session-Id")

	traceId := hdl.newUUID()
	log.SetPrefix(traceId + " ")
	context.Header("X-Trace-Id", traceId)

	if !isLatPresent || !isLonPresent {
		context.AbortWithStatusJSON(400, map[string]interface{}{
			"errorMessage": "Missing Query params (cpLat, cpLon)",
		})
		return
	}

	cpLat, errLat := strconv.ParseFloat(queryLat, 64)
	cpLon, errLon := strconv.ParseFloat(queryLon, 64)
	radious, errRad := strconv.ParseFloat(queryRadious, 64)

	if errLat != nil || errLon != nil || errRad != nil {
		context.AbortWithStatusJSON(400, map[string]interface{}{
			"errorMessage": "Bad format on query params",
		})
		return
	}

	var onSessionError bool = false
	var onError domain.OnError = domain.OnError{}

	// Default if not session created.
	var sessionData session.SearchSessionData = session.SearchSessionData{
		SessionData: session.SessionData{
			SessionConfiguration: session.SessionConfig{
				SessionType: session.Empty,
			},
		},
		Spots: make(map[string][]string),
	}

	if len(headerSessionUUID) > 0 {
		log.Println("Session UUID Present on request: UUID:", headerSessionUUID)

		sessionSearch, err := hdl.session.GetSessionData(headerSessionUUID, session.SpotsReturned)

		if err != nil {
			log.Println("Error -> It is not possible to get session")
			log.Println(err)
			onSessionError = true
			onError = domain.OnError{
				Error:   domain.SessionNotFounded,
				TraceId: traceId,
			}
		} else {
			sessionData = sessionSearch
			log.Println("Session UUID Fetched perfetcly: SessionSearch:", fmt.Sprintf("%+v", sessionData))
		}

	} else if queryCreateSession == "true" {
		log.Println("Create session query present on the request.")
		sessionConfig, err := hdl.session.CreateSession(session.SpotsReturned)

		if err != nil {
			log.Println("Error -> It is not possible to create session")
			log.Println(err)
			onSessionError = true
			onError = domain.OnError{
				Error:   domain.SessionNotCreated,
				TraceId: traceId,
			}
		} else {
			log.Println("Session Configuration created correctly: SessionConfig:", fmt.Sprintf("%+v", sessionConfig))
			sessionData.SessionConfiguration = *sessionConfig
			log.Println("SessionData:", fmt.Sprintf("%+v", sessionData))

		}
	} else {
		log.Println("Nor Query or header where present on the request. Proceeding with default session.")
	}

	response := hdl.service.FilterByProximity(cpLat, cpLon, radious, sessionData)

	if sessionData.SessionConfiguration.SessionType != session.Empty {
		log.Println("Checking if we need to update the session.")
		if len(response.Places) > 0 {
			log.Println("There is values in the result, storing them in the session.")
			var spotsIds []string = make([]string, len(response.Places))

			for index, place := range response.Places {
				spotsIds[index] = place.EventInfo.UUID
			}
			err := hdl.session.AddSpotsToSession(sessionData.SessionConfiguration.SessionId, sessionData.SessionConfiguration.SessionType, spotsIds)
			log.Println(fmt.Sprintf("Error: %+v", err))
		} else {
			log.Println("Empty response, omiting storing ids on session.")
		}
		httpResponse := &domain.ProxymityResult{
			Result: response,
			SearchSessionResponse: domain.SearchSessionResponse{
				SessionDetail: *domain.NewSessionDetails(sessionData.SessionConfiguration.SessionId, sessionData.SessionConfiguration.TTL),
			},
		}
		context.JSON(200, httpResponse)
		return
	}

	if onSessionError {
		httpResponse := &domain.ProxymityResult{
			Result: response,
			SearchSessionResponse: domain.SearchSessionResponse{
				OnError: onError,
			},
		}
		context.JSON(200, httpResponse)
	} else {
		httpResponse := &domain.ProxymityResult{
			Result: response,
		}
		context.JSON(200, httpResponse)
	}

}

func (hdl *HTTPHandler) newUUID() string {
	return uuid.NewString()
}
