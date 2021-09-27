package motion

import (
	"fmt"
	"github.com/mauritsderuiter95/godaemon/pkg/core"
)

type Motion struct {
}

func (m Motion) Initialize() {
	core.Entity{EntityId: "binary_sensor.motion_sensor_overloop"}.OnChange(m.OnChange)
}

func (m Motion) OnChange(event core.Event) {
	if event.Data.NewState.State == "on" {
		core.Entity{EntityId: "light.bulb_overloop"}.TurnOn(nil)
		core.RunIn(0, 1, m.CheckMotion)
	}
}

func (m Motion) CheckMotion() {
	e, err := core.GetEntity("binary_sensor.motion_sensor_overloop")
	if err != nil {
		fmt.Println(err)
		return
	}
	if e.State.State == "on" {
		core.RunIn(0, 1, m.CheckMotion)
	} else {
		core.Entity{EntityId: "light.bulb_overloop"}.TurnOff()
	}
}
