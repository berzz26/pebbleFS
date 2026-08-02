package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/berzz/pebbleFS/internal/service"
)

var app *service.Pebble

var rootCmd = &cobra.Command{
	Use:   "pebble",
	Short: "PebbleFS distributed storage engine",
}

func Execute() {

	var err error

	app, err = service.New()
	if err != nil {
		log.Fatal(err)
	}

	defer app.Metadata.Close()

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
