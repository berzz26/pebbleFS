package main

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info <file>",
	Short: "Display detailed information about a file",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {

		file, chunks, err := app.Info(args[0])
		if err != nil {
			return err
		}

		fmt.Println("File Information")
		fmt.Println("----------------")

		fmt.Println("Name   :", file.Name)
		fmt.Println("Size   :", humanize.Bytes(uint64(file.Size)))
		fmt.Println("Chunks :", len(chunks))
		fmt.Println("")

		fmt.Printf(
			"%-8s %-12s %-10s\n",
			"INDEX",
			"DISK",
			"SIZE",
		)

		fmt.Println("------------------------------------")

		for _, c := range chunks {

			id := c.DiskID
			if len(id) > 8 {
				id = id[:8]
			}

			fmt.Printf(
				"%-8d %-12s %-10s\n",
				c.ChunkIndex,
				id,
				humanize.Bytes(uint64(c.Size)),
			)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}