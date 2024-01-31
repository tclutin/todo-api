package app

import (
	"log/slog"
	"net/http"
	"time"
	"todo/internal/config"
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

func (a *App) startHTTP() {
	a.logger.Info("Initial http server on", a.cfg.Host+":"+a.cfg.Port)
	a.httpServer = &http.Server{
		Addr:         a.cfg.Host + ":" + a.cfg.Port,
		Handler:      a.router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	a.httpServer.ListenAndServe()
}
