package handler

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/JuanGQCadavid/now-project/services/fileService/internal/core"
	"github.com/JuanGQCadavid/now-project/services/fileService/internal/core/domain"
	"github.com/gin-gonic/gin"
)

const (
	TRACE_ID_HEADER = "X-Trace-Id"
)

type HttpHandler struct {
	srv *core.FileService
}

func NewHttpHandler(srv *core.FileService) *HttpHandler {
	return &HttpHandler{
		srv: srv,
	}
}

func (hdl *HttpHandler) SetRouter(router *gin.Engine) {
	router.POST("/files/generate/url", hdl.GetPresignedURL)
}

func (hdl *HttpHandler) getOrGenerateTraceID(context *gin.Context) string {
	var (
		headerTraceId = context.Request.Header.Get(TRACE_ID_HEADER)
	)

	if len(headerTraceId) > 0 {
		return headerTraceId
	}

	return uuid.NewString()
}

func (hdl *HttpHandler) GetPresignedURL(ctx *gin.Context) {

	var (
		TraceID                       = hdl.getOrGenerateTraceID(ctx)
		logger                        = log.With().Str(TRACE_ID_HEADER, TraceID).Logger()
		metadata *domain.FileMetadata = &domain.FileMetadata{}
	)

	if err := ctx.BindJSON(metadata); err != nil {
		logger.Err(err).Msg("Error while casting payload")

		ctx.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage{
			Message: "Unable to cast payload",
		})
		return
	}

	presigned, err := hdl.srv.GeneratePresignedURL(logger.WithContext(context.Background()), metadata)

	if err != nil {
		logger.Err(err).Msg("Error while generating the Presigned URL")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage{
			Message:       "It is us, not you...",
			InternalError: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, presigned)
}
