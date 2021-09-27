package core

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"strconv"
	"time"
)

type Schedule struct {
}

var s = gocron.NewScheduler(time.Now().Location())

func StartScheduler() {
	s.StartAsync()
}

func RunEveryDay(hour, minute int64, f func()) {
	sHour := strconv.FormatInt(hour, 10)
	sMinute := strconv.FormatInt(minute, 10)

	if _, err := s.Cron(fmt.Sprintf("%s %s * * *", sMinute, sHour)).Do(f); err != nil {
		fmt.Println(err)
	}
}

func RunEveryHour(minute int64, f func()) {
	sMinute := strconv.FormatInt(minute, 10)

	if _, err := s.Cron(fmt.Sprintf("%s * * * *", sMinute)).Do(f); err != nil {
		fmt.Println(err)
	}
}

func RunEveryMinute(f func()) {
	if _, err := s.Cron("* * * * *").Do(f); err != nil {
		fmt.Println(err)
	}
}

func RunEveryCron(cron string, f func()) {
	if _, err := s.Cron(cron).Do(f); err != nil {
		fmt.Println(err)
	}
}

func RunIn(hours, minutes int, f func()) {
	d, err := time.ParseDuration(fmt.Sprintf("%dh%dm", hours, minutes))
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		time.Sleep(d)
		f()
	}()
}
