package cmd

import (
	"github.com/spf13/cobra"
)

var progressCmd = &cobra.Command{
	Use:   "progress",
	Short: "Show XP progress and journal",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("XP: 0\nJournal: (not implemented yet)")
	},
}

func init() {
	rootCmd.AddCommand(progressCmd)
}
