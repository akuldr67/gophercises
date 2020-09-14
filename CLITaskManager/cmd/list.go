package cmd

import (
	"fmt"
	"os"

	"akul.gupta/CLITaskManager/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list tasks in CLITaskManager",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ListAllTasks()
		if err != nil {
			fmt.Println("some error occured! Please try Agian.")
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no task to do!")
			return
		}

		fmt.Println("You have following tasks to do:")
		for i, task := range tasks {
			fmt.Println(i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
