package observer

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

type Observer struct {
	// Db         *sql.DB
	Watcher       *fsnotify.Watcher
	WatchPaths    []string
	Results       []string
	Notifications chan string
}

func NewObserver(configPath string) Observer {

	w, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatal(err)
	}

	// These paths come from the config file
	cwd, _ := os.Getwd()

	o := Observer{
		Watcher: w,
		// These paths will come from the config file.
		// Here just for testing
		WatchPaths:    []string{"/tmp", fmt.Sprintf("%s/testing_files", cwd)},
		Notifications: make(chan string),
	}

	for _, p := range o.WatchPaths {
		err = o.Watcher.Add(p)

		if err != nil {
			log.Fatal(err)
		}
	}

	return o
}

func (o Observer) Watch() {
	go func() {
		for {
			select {
			case event, ok := <-o.Watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					log.Println("event:", event)
					o.Notifications <- event.Name
				}
			case err, ok := <-o.Watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()
}

func (o Observer) CloseWatcher() {
	err := o.Watcher.Close()
	if err != nil {
		log.Println("error:", err)
	}
}

func loadConfigs(configPath string) {

	//read toml file
}
