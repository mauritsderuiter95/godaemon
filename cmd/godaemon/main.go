package main

import (
	"github.com/mauritsderuiter95/godaemon/godaemon/apps"
	"github.com/mauritsderuiter95/godaemon/pkg/core"
	"log"
)

func main() {
	inst := core.GetInstance()

	userApps := apps.Register()

	for _, app := range userApps {
		app.Initialize()
	}

	go inst.HandleEvents()

	if err := inst.CloseConnection(); err != nil {
		log.Fatalln(err)
	}
}
