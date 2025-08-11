// Package repo contains clients for data persistence
package repo

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
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
	MarkDone(taskNr int) error
}

type CsvTaskRepository struct{}

func openFile() *os.File {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)

	if err != nil {
		panic(err)
	}

	return file
}

func readAllTasks(file *os.File) []*domain.Task {
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3
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
		isDone, err := strconv.ParseBool(row[2])
		if err != nil {
			panic(err)
		}

		task := domain.Task{Description: description, CreatedAt: createdAt, IsDone: isDone}
		tasks = append(tasks, &task)
	}

	return tasks
}

func taskToString(task *domain.Task) []string {
	return []string{task.Description, task.CreatedAt.Format(dateFormat), strconv.FormatBool(task.IsDone)}
}

func (csvRepository CsvTaskRepository) Add(task *domain.Task) {
	file := openFile()
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	row := taskToString(task)
	err := writer.Write(row)

	if err != nil {
		panic(err)
	}
}

func (csvRepository CsvTaskRepository) GetAll() []*domain.Task {
	file := openFile()
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3
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
		isDone, err := strconv.ParseBool(row[2])
		if err != nil {
			panic(err)
		}

		task := domain.Task{Description: description, CreatedAt: createdAt, IsDone: isDone}
		tasks = append(tasks, &task)
	}
	return tasks
}

func (csvRepository *CsvTaskRepository) MarkDone(taskNr int) error {
	file := openFile()
	defer file.Close()

	tasks := readAllTasks(file)

	var targetTask *domain.Task
	for i, task := range tasks {
		if i == taskNr-1 {
			targetTask = task
		}
	}

	if targetTask == nil {
		details := fmt.Sprintf("Cannot find a task with the number %d", taskNr)
		return errors.New(details)
	}

	targetTask.IsDone = true

	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, task := range tasks {
		row := taskToString(task)
		writer.Write(row)
	}

	return nil
}

func NewCsvTaskRepository() TaskRepository {
	return &CsvTaskRepository{}
}
