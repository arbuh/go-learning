package cmd

import (
	"fmt"

	"todo/app"

	"github.com/spf13/cobra"
)

func createListCmd(appConext *app.Context) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Current tasks:")

			tasks := appConext.TaskRepository.GetAll()
			for i, task := range tasks {
				fmt.Printf("%d. %s\n", i, task.Description)
			}
		},
	}
	return cmd
}
