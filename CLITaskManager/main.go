package main

import (
	"fmt"
	"os"
	"path/filepath"

	"akul.gupta/CLITaskManager/db"

	"akul.gupta/CLITaskManager/cmd"
	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	err := db.CreateDB(dbPath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	cmd.RootCmd.Execute()
}
