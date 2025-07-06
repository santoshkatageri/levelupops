package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/spf13/cobra"
)

var questsCmd = &cobra.Command{
	Use:   "quests",
	Short: "Quest-related commands",
}

var browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "List available quests",
	Run: func(cmd *cobra.Command, args []string) {
		questsDir := "quests"
		entries, err := os.ReadDir(questsDir)
		if err != nil {
			fmt.Printf("Error reading quests directory: %v\n", err)
			return
		}
		for _, entry := range entries {
			if entry.IsDir() {
				questYaml := filepath.Join(questsDir, entry.Name(), "quest.yaml")
				if _, err := os.Stat(questYaml); err == nil {
					fmt.Println(entry.Name())
				}
			}
		}
	},
}

func init() {
	questsCmd.AddCommand(browseCmd)
	rootCmd.AddCommand(questsCmd)
}
