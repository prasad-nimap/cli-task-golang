package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/thejunghare/task/db"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Print("Something went wrong:", err.Error())
			return
		}
		fmt.Printf("Added: %s", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
