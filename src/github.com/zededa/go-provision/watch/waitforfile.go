// Copyright (c) 2017 Zededa, Inc.
// All rights reserved.

// Wait for a named file to appear in a given directory

package watch

import (
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"path/filepath"
)

var debug = true

func WaitForFile(filename string) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err, ": NewWatcher")
	}
	defer w.Close()

	done := make(chan bool)
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				if debug {
					log.Println("stopping")
				}
				return
			case event := <-w.Events:
				if debug {
					log.Println("event:", event)
				}
				if event.Name != filename {
					if debug {
						log.Println("diff file:", event.Name)
					}
					break
				}
				if event.Op & fsnotify.Create != 0 {
					done <- true
				}
			case err := <-w.Errors:
				log.Println("WaitForFile error:", err)
			}
		}
	}()

	dirname := filepath.Dir(filename)
	err = w.Add(dirname)
	if err != nil {
		log.Fatal(err, ": ", filename)
	}
	_, err = os.Stat(filename)
	if err == nil {
		if debug {
			log.Println("found file:", filename)
		}
		stop <- true
		return
	}
	<-done
	stop <- true
}