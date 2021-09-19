package main

import (
	"github.com/mauritsderuiter95/godaemon/godaemon/apps"
	"github.com/mauritsderuiter95/godaemon/pkg/core"
	"log"
)

func main() {
	userApps := apps.Register()

	for _, app := range userApps {
		app.Initialize()
	}

	inst := core.GetInstance()
	if err := inst.CloseConnection(); err != nil {
		log.Fatalln(err)
	}
}
