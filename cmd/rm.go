package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
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
		fmt.Print("Completed:",ids)
	},
}

func init() {
	RootCmd.AddCommand(rmCmd)
}
