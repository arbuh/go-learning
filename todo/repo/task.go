// Package repo contains clients for data persistence
package repo

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
	"todo/domain"
)

const (
	fileName   = "todo-data.csv"
	dateFormat = time.DateTime
)

type TaskRepository interface {
	Add(task *domain.Task)
	GetAll() []*domain.Task
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

	row := []string{task.Description, task.CreatedAt.Format(dateFormat)}
	err = writer.Write(row)
	if err != nil {
		panic(err)
	}

	fmt.Printf("The task '%s' saved!\n", task.Description)
}

func (csvRepository CsvTaskRepository) GetAll() []*domain.Task {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDONLY, 0644)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var tasks []*domain.Task

	for _, row := range data {
		description := row[0]
		createdAt, err := time.Parse(dateFormat, row[1])
		if err != nil {
			panic(err)
		}

		task := domain.Task{Description: description, CreatedAt: createdAt}
		tasks = append(tasks, &task)
	}
	return tasks
}

func NewCsvTaskRepository() TaskRepository {
	return &CsvTaskRepository{}
}
