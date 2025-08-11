// Package util contains utility entities that do not represent the business logic
package util

import (
	"todo/repo"
)

type Context struct {
	TaskRepository repo.TaskRepository
}
