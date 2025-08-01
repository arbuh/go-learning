package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"

	"todo/app"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A simple CLI-based TODO list",
	Long: `A simple CLI-based TODO list built for learning purposes on spf13/cobra`,
}

func Execute(appContext *app.Context) {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
