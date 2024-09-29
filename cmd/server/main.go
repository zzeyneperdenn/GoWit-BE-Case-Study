package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/zzeyneperdenn/GoWit-BE-Case-Study/internal/api/route"
	"github.com/zzeyneperdenn/GoWit-BE-Case-Study/internal/appbase"
)

const shutdownDuration = 30 * time.Second

func main() {
	ctx, mainCtxStop := context.WithCancel(context.Background())

	app := appbase.New(
		appbase.Init(),
		appbase.WithDependencyInjector(),
		appbase.WithLogger(),
	)
	defer app.Shutdown()

	router := route.BuildRouter(app)

	timeout := time.Duration(app.Config.ServerTimeout) * time.Second

	httpServer := &http.Server{
		Addr:              app.Config.ServerAddress,
		Handler:           router,
		ReadTimeout:       timeout,
		WriteTimeout:      timeout,
		IdleTimeout:       timeout,
		ReadHeaderTimeout: timeout,
	}

	handleSignals(ctx, mainCtxStop, func() {
		shutdownErr := httpServer.Shutdown(ctx)
		if shutdownErr != nil {
			log.Panic().Err(shutdownErr).Msg("server shutdown failed")
		}
	})

	log.Info().Msgf("started server on %s", app.Config.ServerAddress)

	serverErr := httpServer.ListenAndServe()
	if serverErr != nil {
		log.Err(serverErr).Msg("server stopped")
	}

	<-ctx.Done()

}

func handleSignals(ctx context.Context, cancelCtx context.CancelFunc, callback func()) {
	sig := make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig

		shutdownCtx, cancel := context.WithTimeout(ctx, shutdownDuration)

		go func() {
			<-shutdownCtx.Done()

			if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
				panic("graceful shutdown timed out.. forcing exit.")
			}
		}()

		callback()

		cancel()
		cancelCtx()
	}()
}
