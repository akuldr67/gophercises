package cmd

import (
	"fmt"
	"strconv"

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
				fmt.Println("can't add task:", arg)
			}
		}
		fmt.Println(tasksDone)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
