// Package cmd contains the command definitions
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"todo/app"
)

func createRootCmd(appContext *app.Context) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "todo",
		Short: "A simple CLI-based TODO list",
		Long:  `A simple CLI-based TODO list built for learning purposes on spf13/cobra`,
	}

	addCmd := createAddCmd(appContext)
	cmd.AddCommand(addCmd)
	listCmd := createListCmd(appContext)
	cmd.AddCommand(listCmd)

	return cmd
}

func Execute(appContext *app.Context) {
	rootCmd := createRootCmd(appContext)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
