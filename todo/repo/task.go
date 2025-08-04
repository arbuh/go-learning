// Package repo contains clients for data persistence
package repo

import (
	"encoding/csv"
	"os"
	"todo/domain"
	"fmt"
)

const fileName = "todo-data.csv"

type TaskRepository interface {
	Add(task *domain.Task)
	// GetAll() []domain.Task
}

type CsvTaskRepository struct{}

func (csvRepository CsvTaskRepository) Add(task *domain.Task) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	row := []string{task.Description}
	err = writer.Write(row)
	if err != nil {
		panic(err)
	}

	fmt.Printf("The task %s saved!\n", task.Description)
}

func NewCsvTaskRepository() TaskRepository {
	return &CsvTaskRepository{}
}
