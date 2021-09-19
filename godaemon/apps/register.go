package apps

import (
	"github.com/mauritsderuiter95/godaemon/godaemon/apps/nightmode"
	"github.com/mauritsderuiter95/godaemon/pkg/core"
)

func Register() []core.UserApp {
	apps := []core.UserApp{
		nightmode.Nightmode{},
	}

	return apps
}
