package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize config and XP store",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("LevelUpOps initialized!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
