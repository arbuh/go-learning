package cmd

import (
	"fmt"
	"time"

	"todo/domain"
	"todo/util"

	"github.com/spf13/cobra"
)

func createAddCmd(appConext *util.Context) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new task",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		Run: func(cmd *cobra.Command, args []string) {
			task := &domain.Task{Description: args[0], CreatedAt: time.Now()}
			appConext.TaskRepository.Add(task)

			fmt.Printf("The task '%s' saved!\n", task.Description)
		},
	}
	return cmd
}
