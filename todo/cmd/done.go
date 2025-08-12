package cmd

import (
	"fmt"
	"strconv"

	"todo/util"

	"github.com/spf13/cobra"
)

func createMarkDoneCmd(appConext *util.Context) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "done",
		Short: "Mark a task as done",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		Run: func(cmd *cobra.Command, args []string) {
			nrStr := args[0]
			nr, err := strconv.Atoi(nrStr)
			if err != nil {
				fmt.Printf("Error! '%s' is not a valid number\n", nrStr)
				return
			}

			err = appConext.TaskRepository.MarkDone(nr)
			if err != nil {
				fmt.Printf("Cannot mark the task as done: %s\n", err)
				return
			}

			fmt.Printf("The task #%d is marked as done!\n", nr)
		},
	}
	return cmd
}

func createUnmarkDoneCmd(appConext *util.Context) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "undone",
		Short: "Unmark a task as done",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		Run: func(cmd *cobra.Command, args []string) {
			nrStr := args[0]
			nr, err := strconv.Atoi(nrStr)
			if err != nil {
				fmt.Printf("Error! '%s' is not a valid number\n", nrStr)
				return
			}

			err = appConext.TaskRepository.UnmarkDone(nr)
			if err != nil {
				fmt.Printf("Cannot unmark the task as done: %s\n", err)
				return
			}

			fmt.Printf("The task #%d is unmarked as done!\n", nr)
		},
	}
	return cmd
}
