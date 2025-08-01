package repo

import (
	"fmt"
	"todo/domain"
)

type TaskRepository interface {
	Add(task *domain.Task)
	// GetAll() []domain.Task
}

type CsvTaskRepository struct{}

func (csvRepository CsvTaskRepository) Add(task *domain.Task) {
	fmt.Printf("New task added: %s", task.Description)
}

func NewCsvTaskRepository() TaskRepository {
	return &CsvTaskRepository{}
}
