package main

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/spf13/cobra"
)

var diskListCmd = &cobra.Command{
	Use:   "list",
	Short: "List registered disks",

	RunE: func(cmd *cobra.Command, args []string) error {

		disks := app.Storage.Disks()

		fmt.Printf(
			"%-10s %-20s %-12s %-12s %-8s\n",
			"ID",
			"MOUNT",
			"FREE",
			"TOTAL",
			"HEALTH",
		)

		fmt.Println("----------------------------------------------------------------")

		for _, d := range disks {

			id := d.ID
			if len(id) > 8 {
				id = id[:8]
			}

			health := "No"
			if d.Healthy {
				health = "Yes"
			}

			fmt.Printf(
				"%-10s %-20s %-12s %-12s %-8s\n",
				id,
				d.MountPath,
				humanize.Bytes(d.FreeSpace),
				humanize.Bytes(d.TotalSpace),
				health,
			)
		}

		fmt.Println()

		fmt.Printf("Disks Attached : %d\n", app.Storage.Count())
		fmt.Printf("Total Capacity : %s\n", humanize.Bytes(app.Storage.TotalCapacity()))
		fmt.Printf("Free Space     : %s\n", humanize.Bytes(app.Storage.TotalFreeSpace()))

		return nil
	},
}

func init() {
	diskCmd.AddCommand(diskListCmd)
}