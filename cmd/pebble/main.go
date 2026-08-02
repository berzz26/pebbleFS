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

		if err := app.Put(os.Args[3]); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Upload completed successfully.")
	case "get":

		if len(os.Args) < 3 {
			log.Fatal("missing filename")
		}

		destination := os.Args[2]

		if len(os.Args) == 4 {
			destination = os.Args[3]
		}

		if err := app.Get(
			os.Args[2],
			destination,
		); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Download completed.")
	default:
		log.Fatalf("unknown command: %s", os.Args[1])
	}
}
