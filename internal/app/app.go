package app

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mdvasilyev/avito-shop/internal/config"
	"github.com/mdvasilyev/avito-shop/internal/router"
	sloggin "github.com/samber/slog-gin"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunServer(lgr *slog.Logger, cfg *config.Config, dbConn *sql.DB) {
	lgr.Info("Running server")

	rtr := gin.New()

	rtr.Use(sloggin.New(lgr))
	rtr.Use(gin.Recovery())

	router.SetupRoutes(lgr, rtr, dbConn)

	srv := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: rtr.Handler(),
	}

	go func() {
		lgr.Info("Listening for connections")

		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			lgr.Error("Error while listen and serve", "error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	lgr.Info("Shutting down server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		lgr.Error("Server shutdown", "error", err)
	}
	select {
	case <-ctx.Done():
		lgr.Info("5 second timeout.")
	}

	lgr.Info("Exiting server")
}
