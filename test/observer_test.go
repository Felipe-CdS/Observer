package observer

import (
	"fmt"
	"os"
	"slices"
	"sync"
	"testing"

	"github.com/Felipe-CdS/observer"
)

func check(t *testing.T, e error) {
	if e != nil {
		t.Fatalf(e.Error())
	}
}

func TestWatch(t *testing.T) {

	var wg sync.WaitGroup
	o := observer.NewObserver("")

	o.Watch()

	wg.Add(4)

	go func() {
		d1 := []byte("Create testing 1\n")
		err := os.WriteFile("/tmp/observer_test1", d1, 0644)
		check(t, err)

		wg.Done()
	}()

	go func() {
		d2 := []byte("Create testing 2\n")
		f, err := os.Create("/tmp/observer_test2")
		check(t, err)

		_, err = f.Write(d2)
		check(t, err)

		_, err = f.WriteString("Create testing 3\n")
		check(t, err)

		f.Sync()
		f.Close()
		wg.Done()
	}()

	go func() {
		d1 := []byte("Create testing 4\n")
		err := os.WriteFile("./testing_files/observer_test3", d1, 0644)
		check(t, err)

		wg.Done()
	}()

	go func() {
		counter := 0
		for not := range o.Notifications {
			counter++
			o.Results = append(o.Results, not)
			if counter >= 6 {
				wg.Done()
				break
			}
		}
	}()

	wg.Wait()
	o.Watcher.Close()

	cwd, _ := os.Getwd()

	if len(o.Results) != 4 {
		t.Fatalf("Error Test 1")
	}

	if !slices.Contains(o.Results, "/tmp/observer_test1") {
		t.Fatalf("Error Test 2")
	}

	if !slices.Contains(o.Results, fmt.Sprintf("%s/testing_files/observer_test3", cwd)) {
		t.Fatalf("Error Test 3")
	}

	if !slices.Contains(o.Results, "/tmp/observer_test2") {
		t.Fatalf("Error Test 4")
	}

	idx := slices.Index(o.Results, "/tmp/observer_test2")
	o.Results = append(o.Results[:idx], o.Results[idx+1:]...)

	if !slices.Contains(o.Results, "/tmp/observer_test2") {
		t.Fatalf("Error Test 5")
	}
}
