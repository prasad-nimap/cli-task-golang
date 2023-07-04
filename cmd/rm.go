package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/thejunghare/task/db"
)

// rmCmd represents the do command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes the task form list",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int

		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				panic(err)
			} else {
				ids = append(ids, id)
			}
		}

		tasks, err := db.AllTask()
		if err != nil {
			fmt.Print("Something went wrong:", err)
			return
		}

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Print("Invalid task number:", id)
				continue // continue because their might be another vaild number after an invalid number
			}
			task := tasks[id-1]

			err := db.DeleteTask(task.Key)

			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed.Error %s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as completed.\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(rmCmd)
}
