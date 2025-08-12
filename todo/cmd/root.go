// Package cmd contains the command definitions
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"todo/util"
)

func createRootCmd(appContext *util.Context) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "todo",
		Short: "A simple CLI-based TODO list",
		Long:  `A simple CLI-based TODO list built for learning purposes on spf13/cobra`,
	}

	addCmd := createAddCmd(appContext)
	listCmd := createListCmd(appContext)
	markDoneCmd := createMarkDoneCmd(appContext)
	unmarkDoneCmd := createUnmarkDoneCmd(appContext)
	cmd.AddCommand(addCmd, listCmd, markDoneCmd, unmarkDoneCmd)

	return cmd
}

func Execute(appContext *util.Context) {
	rootCmd := createRootCmd(appContext)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
