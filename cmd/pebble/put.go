package main

import "github.com/spf13/cobra"

var putCmd = &cobra.Command{
	Use:   "put <file>",
	Short: "Upload a file",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		return app.Put(args[0])
	},
}

func init() {
	rootCmd.AddCommand(putCmd)
}