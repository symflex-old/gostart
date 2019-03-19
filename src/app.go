package main

import "log"

type AppInterface interface {
	Bootstrap(file string) *App
	Process()
}

type App struct {
	AppInterface
	Config *Config
}

func (app *App) Bootstrap(config *Config) *App {
	app.Config = config
	return app
}

func (app *App) Process() {
	log.Println(app.Config.Debug)
}
