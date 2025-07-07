package cmd

import (
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Run diagnostics",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Diagnostics passed. (not implemented yet)")
	},
}

// The `init` function adds the `doctorCmd` command to the `rootCmd`.
func init() {
	rootCmd.AddCommand(doctorCmd)
}
