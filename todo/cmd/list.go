package cmd

import (
	"fmt"
	"time"

	"todo/app"

	"github.com/spf13/cobra"
)

const (
	dateFormat = time.DateTime
)

func createListCmd(appConext *app.Context) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		Run: func(cmd *cobra.Command, args []string) {
			tasks := appConext.TaskRepository.GetAll()

			if len(tasks) == 0 {
				fmt.Println("No current tasks")
			} else {
				fmt.Println("Current tasks:")

				for i, task := range tasks {
					date := task.CreatedAt.Format(dateFormat)

					fmt.Printf("%d. %s, from %s\n", i, task.Description, date)
				}
			}
		},
	}
	return cmd
}
