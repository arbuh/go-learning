package util

import (
	"fmt"
	"time"

	"todo/domain"
)

func FormatTasks(tasks []*domain.Task) []string {
	var lines []string
	for i, task := range tasks {
		nr := i + 1
		duration := calcDuration(task.CreatedAt)
		isDone := boolToString(task.IsDone)

		line := fmt.Sprintf("%d. %s, %s, %s", nr, task.Description, duration, isDone)
		lines = append(lines, line)
	}
	return lines
}

func calcDuration(t time.Time) string {
	duration := time.Since(t)

	hours := duration.Hours()

	if hours < 24 {
		return "new"
	}

	days := int(hours / 24)

	if days == 1 {
		return "1 day ago"
	} else {
		return fmt.Sprintf("%d days ago", days)
	}
}

func boolToString(b bool) string {
	if (b) {
		return "☑"
	} else {
		return "☐"
	}
}
