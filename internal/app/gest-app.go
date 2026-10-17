package app

import (
	http_server "github.com/martinpetrdev/Gest/internal/http-server"
	"github.com/martinpetrdev/Gest/internal/module"
)

type GestApp struct {
	appModule *module.Module
	server    *http_server.HttpServer
}

func Create(appModule module.Module) (*GestApp, error) {
	return &GestApp{
		appModule: &appModule,
		server:    http_server.Create(),
	}, nil
}

func (app *GestApp) Listen(httpInterface string) error {
	err := app.server.Listen(httpInterface)

	return err
}
