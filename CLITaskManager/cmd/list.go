package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list tasks in CLITaskManager",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List called")
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
