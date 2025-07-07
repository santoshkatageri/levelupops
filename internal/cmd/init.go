package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// The `var initCmd = &cobra.Command{}` statement is creating a new Cobra command named `initCmd`. This
// command has the following properties:
// - `Use: "init"`: Specifies the command name that will be used to invoke this command.
// - `Short: "Initialize config and XP store"`: Provides a short description of what the command does.
// - `Run: func(cmd *cobra.Command, args []string) { ... }`: Defines the function that will be executed
// when this command is run. In this case, it simply prints "LevelUpOps initialized!" to the console.
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize config and XP store",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("LevelUpOps initialized!")
	},
}

// The `init()` function adds the `initCmd` command to the `rootCmd`.
func init() {
	rootCmd.AddCommand(initCmd)
}
