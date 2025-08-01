package main

import (
	"todo/app"
	"todo/cmd"
	"todo/repo"
)

func main() {
	taskRepo := repo.NewCsvTaskRepository()

	context := &app.Context{TaskRepository: taskRepo}

	cmd.Execute(context)
}
