package main

import (
	"github.com/mauritsderuiter95/godaemon/apps"
	"github.com/mauritsderuiter95/godaemon/pkg/core"
	"log"
	"reflect"
)

func main() {
	log.Println("Starting GoDaemon")
	inst := core.GetInstance()
	go core.QuitOnFilechange(inst)

	log.Println("Registering apps")
	userApps := apps.Register()

	for _, app := range userApps {
		app.Initialize()
		log.Println("Registered", reflect.TypeOf(app).Name())
	}

	go inst.HandleEvents()
	log.Println("Listening to events")
	core.StartScheduler()
	log.Println("Started scheduler")
	log.Println("Ready")

	if err := inst.CloseConnection(); err != nil {
		log.Fatalln(err)
	}
}
