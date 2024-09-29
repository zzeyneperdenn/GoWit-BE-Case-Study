package route

import (
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/zzeyneperdenn/GoWit-BE-Case-Study/internal/api/handlers"
	"github.com/zzeyneperdenn/GoWit-BE-Case-Study/internal/appbase"
)

func BuildRouter(appBase *appbase.AppBase) *chi.Mux {
	mux := buildMux(appBase)

	handlers.InitRoutes(mux, appBase)

	return mux
}

func buildMux(appBase *appbase.AppBase) *chi.Mux {
	r := chi.NewRouter()

	r.Use(chiMiddleware.SetHeader("Content-Type", "application/json"))

	r.Use(chiMiddleware.Recoverer)
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Heartbeat("/health_check"))
	r.Use(chiMiddleware.Heartbeat("/healthz"))
	r.Use(chiMiddleware.Heartbeat("/readyz"))

	return r
}
