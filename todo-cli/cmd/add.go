package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add test")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
