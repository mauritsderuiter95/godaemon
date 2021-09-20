package event

import (
	"fmt"
	"github.com/mauritsderuiter95/godaemon/pkg/core"
)

type Event struct {
	Name string
}

func Get(name string) Event {
	return Event{Name: name}
}

func (e Event) OnChange(id string, f func(event core.HaEvent)) {
	ha := core.GetInstance()

	ha.Callbacks[id] = append(ha.Callbacks[id], func(event core.HaEvent) {
		fmt.Println(event.Event.EventType)
		if event.Event.EventType == e.Name {
			f(event)
		}
	})
}
