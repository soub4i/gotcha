package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

const (
	// Version of the application
	Version = "1.0.0"
	NO_ARGS = "No arguments provided"
)


func help()  {
		fmt.Println("Help")
		fmt.Println("Usage: watcher <path>")
		fmt.Println("  version - prints the version of the application")
		fmt.Println("  help - prints the help")
		fmt.Println("  <path> - the path to the directory to watch")
}


func main() {

	args := os.Args

	if len(args) == 1 {
		fmt.Println(NO_ARGS)
		return
	}

	// events := make(chan fsnotify.Event)

	watcher, err := fsnotify.NewWatcher()
	defer watcher.Close()

	if err != nil {
		log.Fatal(err)
	}
	switch args[1] {
	case "version":
		fmt.Println(Version)
	case "help":
		help()
	case "watch":
		if len(args) < 2 {
			fmt.Println("No path provided")
			return
		}
		path := args[2]
		fmt.Printf("Watching %s", path)

		go func() {
			for {
				fmt.Println("de")
				select {
				case event, ok := <-watcher.Events:
					if !ok {
						log.Fatal("Error:", err)
						return
					}
					fmt.Println(event)
				}
			}
		}()
		err := watcher.Add(path)
		if err != nil {
			log.Fatal(err)
		}
		<-make(chan struct{})

	default: help()
	}

}
