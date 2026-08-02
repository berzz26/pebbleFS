package main

import "github.com/spf13/cobra"

var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "Manage Pebble disks",
}

func init() {
	rootCmd.AddCommand(diskCmd)
}
