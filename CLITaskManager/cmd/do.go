package cmd

import (
	"fmt"
	"os"
	"strconv"

	"akul.gupta/CLITaskManager/db"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "mark task as complete CLITaskManager",
	Run: func(cmd *cobra.Command, args []string) {
		tasksDone := []int{}
		for _, arg := range args {
			s, err := strconv.Atoi(arg)
			if err == nil {
				tasksDone = append(tasksDone, s)
			} else {
				fmt.Printf("can't do task: \"%v\", Give a valid task number\n", arg)
			}
		}

		tasks, e := db.ListAllTasks()
		if e != nil {
			fmt.Println("some error occured! Please try Agian.")
			os.Exit(1)
		}

		for _, taskID := range tasksDone {
			if taskID > len(tasks) {
				fmt.Println(taskID, "is not a valid task ID")
				continue
			}
			err := db.DeleteTask(tasks[taskID-1].ID)
			if err != nil {
				fmt.Println("some error occured! Please try Agian.")
				os.Exit(1)
			}
			fmt.Println("Task", taskID, "done successfully!")
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
