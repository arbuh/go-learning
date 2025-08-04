// Package app contain shared utility entities that do not represent the business logic
package app

import (
	"todo/repo"
)

type Context struct {
	TaskRepository repo.TaskRepository
}
