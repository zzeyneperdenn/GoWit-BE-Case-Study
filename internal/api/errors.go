package api

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
	"github.com/zzeyneperdenn/GoWit-BE-Case-Study/internal/api/server"
)

func RenderHTTPError(httpStatus int, err server.ErrorResponse, r *http.Request, w http.ResponseWriter) {
	log.Ctx(r.Context()).Error().Msg(err.Detail)

	render.Status(r, httpStatus)
	render.JSON(w, r, err)
}

func InternalServerError() server.ErrorResponse {
	return server.ErrorResponse{
		Status: http.StatusInternalServerError,
		Title:  http.StatusText(http.StatusInternalServerError),
		Detail: "There was a problem processing your request",
	}
}

func NotFoundError() server.ErrorResponse {
	return server.ErrorResponse{
		Status: http.StatusNotFound,
		Title:  http.StatusText(http.StatusNotFound),
		Detail: "Not found, please try again",
	}
}

func BadRequestError() server.ErrorResponse {
	return server.ErrorResponse{
		Status: http.StatusBadRequest,
		Title:  http.StatusText(http.StatusBadRequest),
		Detail: "There was a problem processing your request, please try again",
	}
}
