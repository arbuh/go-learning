package main

import (
	"todo/cmd"
	"todo/repo"
	"todo/util"
)

func main() {
	taskRepo := repo.NewCsvTaskRepository()

	context := &util.Context{TaskRepository: taskRepo}

	cmd.Execute(context)
}
