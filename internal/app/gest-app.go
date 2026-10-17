package app

import (
	"net/http"

	"github.com/martinpetrdev/Gest/internal/DI"
	http_server "github.com/martinpetrdev/Gest/internal/http-server"
	"github.com/martinpetrdev/Gest/internal/logger"
	"github.com/martinpetrdev/Gest/internal/module"
)

type GestApp struct {
	appModule *module.Module
	server    *http_server.HttpServer
	logger    *logger.Logger
}

func Create(appModule module.Module) (*GestApp, error) {
	appLogger := logger.New("GestApplication")
	appLogger.Log("Starting Gest application...")

	container := DI.New()
	mux := http.NewServeMux()

	if err := appModule.Register(container, mux, map[*module.Module]bool{}); err != nil {
		return nil, err
	}

	return &GestApp{
		appModule: &appModule,
		server:    http_server.Create(mux),
		logger:    appLogger,
	}, nil
}

func (app *GestApp) Listen(httpInterface string) error {
	err := app.server.Listen(httpInterface)

	return err
}
