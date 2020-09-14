package cmd

import (
	"fmt"
	"os"
	"strings"

	"akul.gupta/CLITaskManager/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add task to CLITaskManager",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")

		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("some error occured! Please try Agian.")
			os.Exit(1)
		}
		fmt.Printf("Task \"%v\" added successfully \n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
