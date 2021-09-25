package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"path/filepath"
)

func QuitOnFilechange(ha HomeAssistant) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case _, ok := <-watcher.Events:
				if !ok {
					return
				}
				ha.done <- struct{}{}
				return
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	filepath.Walk("./apps/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("ERROR:", err)
		}
		if info.IsDir() {
			err = watcher.Add(path)
		}
		return nil
	})
	<-done
}
