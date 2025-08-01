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
			fmt.Println("list test")
		},
	}
	return cmd
}
