package main

import (
	"fmt"
	"github.com/berzz/pebblefs/internal/service"
	"log"
)

func main() {
	fmt.Println("Initializing PebbleFS")

	app, err := service.New()
	if err != nil {
		log.Fatal(err)
	}

	defer app.close()

	fmt.Println("PebbleFS initialized successfully")
}
