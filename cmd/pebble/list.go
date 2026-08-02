package main

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List stored files",

	RunE: func(cmd *cobra.Command, args []string) error {

		files, err := app.ListFiles()
		if err != nil {
			return err
		}

		fmt.Printf(
			"%-30s %-12s %-20s\n",
			"NAME",
			"SIZE",
			"CREATED",
		)

		fmt.Println("--------------------------------------------------------------------")

		for _, file := range files {

			fmt.Printf(
				"%-30s %-12s %-20s\n",
				file.Name,
				humanize.Bytes(uint64(file.Size)),
				file.CreatedAt.Format("2006-01-02 15:04"),
			)
		}

		fmt.Println()
		fmt.Printf("%d file(s)\n", len(files))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}