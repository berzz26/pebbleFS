package main

import (
	"path/filepath"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get <file> [destination]",
	Short: "Download a file",
	Args:  cobra.RangeArgs(1, 2),

	RunE: func(cmd *cobra.Command, args []string) error {

		dest := filepath.Base(args[0])

		if len(args) == 2 {
			dest = args[1]
		}

		return app.Get(args[0], dest)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}