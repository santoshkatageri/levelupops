package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start <id>",
	Short: "Start a specific quest",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		questID := args[0]
		questPath := "quests/" + questID + "/quest.yaml"
		if _, err := os.Stat(questPath); err != nil {
			fmt.Printf("Quest '%s' not found.\n", questID)
			return
		}
		fmt.Printf("Quest '%s' started!\n", questID)
	},
}

// The `init` function in Go adds the `startCmd` command to the `rootCmd`.
func init() {
	rootCmd.AddCommand(startCmd)
}
