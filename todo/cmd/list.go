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
					i = i + 1 // to make it more human friendly
					duration := calcDuration(task.CreatedAt)

					fmt.Printf("%d. %s, %s\n", i, task.Description, duration)
				}
			}
		},
	}
	return cmd
}

func calcDuration(t time.Time) string {
	duration := time.Since(t)

	hours := duration.Hours()

	if hours < 24 {
		return "new"
	}

	days := int(hours / 24)

	if days == 1 {
		return "1 day ago"
	} else {
		return fmt.Sprintf("%d days ago", days)
	}
}
