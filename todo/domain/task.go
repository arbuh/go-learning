// Package domain contain shared business logic entities
package domain

import "time"

type Task struct {
	Description string
	CreatedAt time.Time
}
