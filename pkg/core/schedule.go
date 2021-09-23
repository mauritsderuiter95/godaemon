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

func RunDaily(hour, minute int64, f func()) {
	sHour := strconv.FormatInt(hour, 10)
	sMinute := strconv.FormatInt(minute, 10)

	if _, err := s.Cron(fmt.Sprintf("%s %s * * *", sMinute, sHour)).Do(f); err != nil {
		fmt.Println(err)
	}
}
