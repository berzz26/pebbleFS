package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "rm <file>",
	Short: "Delete a stored file",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {

		err := app.Delete(args[0])
		if err != nil {
			return err
		}

		fmt.Println("File deleted successfully.")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}