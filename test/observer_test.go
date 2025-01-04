package observer

import (
	"log"
	"os"
	"testing"

	"github.com/Felipe-CdS/observer"
)

func check(t *testing.T, e error) {
	if e != nil {
		t.Fatalf(e.Error())
	}
}

func TestWatch(t *testing.T) {

	o := observer.NewObserver("")

	o.Watch()

	d1 := []byte("Create testing 1\n")
	d2 := []byte("Create testing 2\n")

	err := os.WriteFile("/tmp/observer_test1", d1, 0644)
	check(t, err)

	log.Println(o.Results)

	f, err := os.Create("/tmp/observer_test2")
	check(t, err)

	_, err = f.Write(d2)
	check(t, err)

	log.Println(o.Results)

	_, err = f.WriteString("Create testing 3\n")
	check(t, err)

	log.Println(o.Results)

	f.Sync()
	f.Close()

	o.Watcher.Close()
}
