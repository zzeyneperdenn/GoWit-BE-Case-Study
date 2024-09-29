package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/samber/do"
	"github.com/zzeyneperdenn/GoWit-BE-Case-Study/internal/api/server"
	"github.com/zzeyneperdenn/GoWit-BE-Case-Study/internal/appbase"
	"github.com/zzeyneperdenn/GoWit-BE-Case-Study/internal/services"
)

type Handler struct {
	ticketsService services.ITicketsService
}

func newRootHandler(app *appbase.AppBase) *Handler {
	return &Handler{
		ticketsService: do.MustInvoke[services.ITicketsService](app.Injector),
	}
}

func InitRoutes(router chi.Router, app *appbase.AppBase) {
	server.HandlerFromMux(
		newRootHandler(app),
		router,
	)
}
