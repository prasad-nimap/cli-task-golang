package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thejunghare/task/db"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all the tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTask()
		if err != nil {
			fmt.Print("Something went wrong:", err.Error())
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Print("You have no tasks to complete")
			return
		}

		fmt.Print("You have following task: \n")
		for i, task := range tasks {
			/* index + 1 as the index starts from zero and we want task to display from 1 */
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(lsCmd)
}
