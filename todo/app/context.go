package app

import (
	"todo/repo"
)

type Context struct {
	TaskRepository repo.TaskRepository
}
