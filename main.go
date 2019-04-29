package main

import (
	"flag"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

var flagDir string

func init() {
	flag.StringVar(&flagDir, "d", "", "directory to watch for changes")
}

func main() {
	flag.Parse()
	if flagDir == "" {
		flag.Usage()
		os.Exit(1)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(flagDir)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("watching directory: ", flagDir)

	<-done
}
