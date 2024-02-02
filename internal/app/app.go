package app

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo/internal/config"
	"todo/internal/handler/middleware"
)

type App struct {
	cfg        *config.Config
	logger     *slog.Logger
	httpServer *http.Server
	router     *http.ServeMux
}

func New(cfg *config.Config, logger *slog.Logger, router *http.ServeMux) App {
	return App{cfg: cfg, logger: logger, router: router}
}

func (a *App) Run() {
	a.startHTTP()
}

func (a *App) Shutdown() {
	err := a.httpServer.Shutdown(context.Background())
	if err != nil {
		a.logger.Error("error occured on server shuttind down", err)
		os.Exit(1)
	}
}

func (a *App) startHTTP() {
	a.logger.Info("Initial http server on", a.cfg.Host+":"+a.cfg.Port)
	handler := middleware.Logging(a.router)
	a.httpServer = &http.Server{
		Addr:         a.cfg.Host + ":" + a.cfg.Port,
		Handler:      handler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		err := a.httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			a.logger.Error(err.Error())
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	a.logger.Info("Shutting down")
	a.Shutdown()
}
