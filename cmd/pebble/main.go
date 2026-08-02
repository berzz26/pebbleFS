package main

import (
	"fmt"
	"github.com/berzz/pebbleFS/internal/service"
	"log"
)

func main() {
	fmt.Println("Initializing PebbleFS")

	app, err := service.New()
	if err != nil {
		log.Fatal(err)
	}

	defer app.Metadata.Close()

	fmt.Println("PebbleFS initialized successfully")
}
