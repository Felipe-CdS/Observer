package observer

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

type Observer struct {
	// Db         *sql.DB
	Watcher    *fsnotify.Watcher
	WatchPaths []string
	Results    []string
}

func NewObserver(configPath string) Observer {

	w, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatal(err)
	}

	o := Observer{
		Watcher:    w,
		WatchPaths: []string{"/tmp", "."},
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

	defer o.Watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-o.Watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Has(fsnotify.Write) {
					// Parse markdown
					// Write in database
					o.Results = append(o.Results, event.Name)
					log.Println("modified file:", event.Name)
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
	o.Watcher.Close()
}

func loadConfigs(configPath string) {

	//read toml file
}
