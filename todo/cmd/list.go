package cmd

import (
	"fmt"
	"time"

	"todo/util"

	"github.com/spf13/cobra"
)

const (
	dateFormat = time.DateTime
)

func createListCmd(appConext *util.Context) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		Run: func(cmd *cobra.Command, args []string) {
			tasks := appConext.TaskRepository.GetAll()

			if len(tasks) == 0 {
				fmt.Println("No current tasks")
			} else {
				fmt.Println("Current tasks:")
				lines := util.FormatTasks(tasks)

				for _, line := range lines {
					fmt.Println(line)
				}
			}
		},
	}
	return cmd
}
