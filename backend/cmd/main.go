package main

import (
	api "backend/cmd/modes/api"
	menu "backend/cmd/modes/techUI"
	"backend/cmd/registry"
	"log"
)

func main() {
	app := registry.App{}

	err := app.Config.ParseConfig("config.json", "../config")
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run()

	if err != nil {
		log.Fatal(err)
	}

	if app.Config.Mode == "tech" {
		app.Logger.Info("Start with tech ui!")
		err = menu.RunMenu(app.Services)
		if err != nil {
			log.Fatal(err)
		}	
	} else if app.Config.Mode == "api" {
		app.Logger.Info("Start with api!")
		api.SetupServer(&app)
	} else {
		app.Logger.Error("Wrong app mode", "mode", app.Config.Mode)
	}
}
