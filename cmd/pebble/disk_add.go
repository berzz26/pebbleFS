package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var diskAddCmd = &cobra.Command{
	Use:   "add <mount-path>",
	Short: "Register a new storage device",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {

		disk, err := app.Storage.AddDisk(args[0])
		if err != nil {
			return err
		}

		if err := app.Metadata.RegisterDisk(*disk); err != nil {
			return err
		}

		fmt.Println("Disk registered successfully.")

		return nil
	},
}

func init() {
	diskCmd.AddCommand(diskAddCmd)
}