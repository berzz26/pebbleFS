package main

import (
	"fmt"
	"log"
	"os"

	"github.com/berzz/pebbleFS/internal/service"
)

func main() {

	app, err := service.New()
	if err != nil {
		log.Fatal(err)
	}
	defer app.Metadata.Close()

	// TODO: Make this configurable later.
	if err := app.Storage.AddDisk("/mnt/usb"); err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  pebble put <file>")
		return
	}

	switch os.Args[1] {

	case "put":

		if len(os.Args) < 3 {
			log.Fatal("missing file path")
		}

		if err := app.Put(os.Args[2]); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Upload completed successfully.")

	default:
		log.Fatalf("unknown command: %s", os.Args[1])
	}
}
