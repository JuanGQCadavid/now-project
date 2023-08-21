// Package docs Filter service.
//
// This service filters schedule and in-time spots
// base on the user proximity, furthermore there will be
// filters base on treanding spots.
//
//	    	Schemes: https, http
//	    	BasePath: /filter
//	    	Version: 1.0.0
//	    	Host: api.pululapp.com
//
//		   	TermsOfService: http://swagger.io/terms/
//			Contact: Juan Gonzalo Quiroz Cadavid <jquirozcadavid@gmail.com> http://john.blogs.com
//			License: MIT http://opensource.org/licenses/MIT
//
//	    	Consumes:
//	    	- application/json
//
//	    	Produces:
//	    	- application/json
//
// swagger:meta
package httphdl

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

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

// swagger:route GET /proximity ByProximity filter
//
// Lists spots filtered by some parameters, also it stores the result if specified
// in order to avoid returning the same data in future calls.
//
//	    Consumes:
//	    - application/json
//
//	    Produces:
//	    - application/json
//
//	    Schemes: http, https
//
//	    Parameters:
//	      + name: cpLat
//	        in: query
//	        description: Central point location on latitude coordinates.
//	        required: true
//	        type: number
//			 + name: cpLon
//	        in: query
//	        description: Central point location on longitude coordinates.
//	        required: true
//	        type: number
//	      + name: radious
//	        in: query
//	        description: radious from cpLat, cpLon.
//	        required: false
//	        type: number
//			   ddefault: 0.5
//			 + name: createSession
//	        in: query
//	        description: If true then it returns a search session id with its details in the response.
//	        required: false
//	        type: boolean
//			   default: false
//			 + name: format
//	        in: query
//	        description: if small then some attributes are omited, full will return all data.
//	        required: false
//	        type: string
//			   default: small
//			 + name: X-Now-Search-Session-Id
//	        in: header
//	        description: Session Id, if present then it will be used to return spots that were not prevoiusly returned in the session
//	        required: false
//	        type: string
//
//	    Responses:
//	      default: ProxymityResult
//	      200: ProxymityResult
//	      422: ProxymityResult
func (hdl *HTTPHandler) FilterSpots(context *gin.Context) {
	traceId := hdl.newUUID()
	var onSessionError bool = false
	var onError *domain.OnError

	log.SetPrefix(traceId + " ")
	context.Header("X-Trace-Id", traceId)

	sessionData, onError := hdl.getSession(context, traceId)
	params, err := hdl.getParams(context)

	if err != nil {
		context.AbortWithStatusJSON(400, map[string]interface{}{
			"errorMessage": err.Error(),
		})
		return
	}

	response, err := hdl.service.FilterByProximity(params.Lat, params.Lon, params.Radious, sessionData, params.Format)

	if err != nil {
		httpResponse := &domain.ProxymityResult{
			OnError: domain.OnError{
				Error:   domain.ErrorType(err.Error()),
				TraceId: traceId,
			},
		}
		context.JSON(500, httpResponse)
		return
	}

	if sessionData.SessionConfiguration.SessionType != session.Empty {
		log.Println("Checking if we need to update the session.")
		if len(response.Places) > 0 {
			log.Println("There is values in the result, storing them in the session.")
			var spotsIds []string = make([]string, len(response.Places))

			for index, place := range response.Places {
				spotsIds[index] = place.DateInfo.Id
			}
			err := hdl.session.AddSpotsToSession(sessionData.SessionConfiguration.SessionId, sessionData.SessionConfiguration.SessionType, spotsIds)
			log.Println(fmt.Sprintf("Error: %+v", err))
		} else {
			log.Println("Empty response, omiting storing ids on session.")
		}
		httpResponse := &domain.ProxymityResult{
			Result: *response,
			SearchSessionResponse: domain.SearchSessionResponse{
				SessionDetail: *domain.NewSessionDetails(sessionData.SessionConfiguration.SessionId, sessionData.SessionConfiguration.TTL),
			},
		}
		context.JSON(200, httpResponse)
		return
	}

	if onSessionError {
		httpResponse := &domain.ProxymityResult{
			Result: *response,
			SearchSessionResponse: domain.SearchSessionResponse{
				OnError: *onError,
			},
		}
		context.JSON(200, httpResponse)
	} else {
		httpResponse := &domain.ProxymityResult{
			Result: *response,
		}
		context.JSON(200, httpResponse)
	}

}

type ProximityFilterParams struct {
	Lat     float64
	Lon     float64
	Radious float64
	Format  ports.OutputFormat
}

func (hdl *HTTPHandler) getParams(context *gin.Context) (*ProximityFilterParams, error) {
	queryLat, isLatPresent := context.GetQuery("cpLat")
	queryLon, isLonPresent := context.GetQuery("cpLon")
	queryRadious := context.DefaultQuery("radious", fmt.Sprintf("%f", hdl.defaultRadious))

	queryFormat := context.DefaultQuery("format", "small")

	if !isLatPresent || !isLonPresent {
		return nil, errors.New("Missing Query params (cpLat, cpLon)")
	}

	cpLat, errLat := strconv.ParseFloat(queryLat, 64)
	cpLon, errLon := strconv.ParseFloat(queryLon, 64)
	radious, errRad := strconv.ParseFloat(queryRadious, 64)
	format := ports.SMALL_FORMAT

	if strings.ToUpper(queryFormat) == string(ports.FULL_FORMAT) {
		log.Println(fmt.Sprintf("Using FULL format, format value passed by query: %s", queryFormat))
		format = ports.FULL_FORMAT
	} else {
		log.Println(fmt.Sprintf("Using default (SMALL) format, format value passed by query: %s", queryFormat))
	}

	if errLat != nil || errLon != nil || errRad != nil {
		return nil, errors.New("Bad format on query params")
	}

	return &ProximityFilterParams{
		Lat:     cpLat,
		Lon:     cpLon,
		Radious: radious,
		Format:  format,
	}, nil

}

func (hdl *HTTPHandler) getSession(context *gin.Context, traceId string) (session.SearchSessionData, *domain.OnError) {
	queryCreateSession := context.DefaultQuery("createSession", "false")
	headerSessionUUID := context.Request.Header.Get("X-Now-Search-Session-Id")
	var onError domain.OnError = domain.OnError{}
	var onSessionError bool = false

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

	if onSessionError {
		return sessionData, &onError
	} else {
		return sessionData, nil
	}
}
func (hdl *HTTPHandler) newUUID() string {
	return uuid.NewString()
}
