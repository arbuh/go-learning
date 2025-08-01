package cmd

import (
	"todo/app"
	"todo/domain"

	"github.com/spf13/cobra"
)

func createAddCmd(appConext *app.Context) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new task",
		Run: func(cmd *cobra.Command, args []string) {
			task := &domain.Task{Description: "do the things"}
			appConext.TaskRepository.Add(task)
		},
	}
	return cmd
}
