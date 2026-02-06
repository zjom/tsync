package tasks

import (
	"time"

	"github.com/zjom/tsync/pkg/core/priority"
)

type Task struct {
	Id       string
	Name     string
	Deadline time.Time
	Subtasks []*Task
	Priority priority.Priority
}
