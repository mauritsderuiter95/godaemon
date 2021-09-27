package apps

import (
	"github.com/mauritsderuiter95/godaemon/apps/motion"
	"github.com/mauritsderuiter95/godaemon/apps/nightmode"
	"github.com/mauritsderuiter95/godaemon/pkg/core"
)

func Register() []core.UserApp {
	apps := []core.UserApp{
		nightmode.Nightmode{},
		motion.Motion{},
	}

	return apps
}
