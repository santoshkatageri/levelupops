package cmd

import (
	"github.com/spf13/cobra"
)

var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Remove stale sandboxes/resources",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Cleanup complete. (not implemented yet)")
	},
}

func init() {
	rootCmd.AddCommand(cleanupCmd)
}
