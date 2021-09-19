package core

import "time"

type UserApp interface {
	Initialize()
}

type App struct {

}

func (a *App) ListenState(f func(), id string) {

}

func (a *App) RunDaily(f func(), at time.Time) {

}
