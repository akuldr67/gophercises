package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add task to CLITaskManager",
	Run: func(cmd *cobra.Command, args []string) {
		item := strings.Join(args, " ")
		fmt.Println(item)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
