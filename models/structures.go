// Package models ...
package models

import (
	"loco-tasks/status"
	"time"
)

/* TODO:
- add a command structure
- a task has one command (or several command?) -> one is simpler
*/

// Task ...
type Task struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Path        string   `json:"path"`
	Args        []string `json:"args"`
	Env         []string `json:"env"`
}

/*

 */

// TaskRecord ...
type TaskRecord struct {
	Name              string        `json:"name"`
	Description       string        `json:"description"`
	CurrentStatus     status.Status `json:"currentStatus"`
	StatusDescription string        `json:"statusDescription"`
	CreatedAt         time.Time     `json:"createdAt"` // record the time the event was requested
	StartedAt         time.Time     `json:"startedAt"`
	FinishedAt        time.Time     `json:"finishedAt"`
	CancelledAt       time.Time     `json:"cancelledAt"`
	FailedAt          time.Time     `json:"failedAt"`
	CheckedAt         time.Time     `json:"checkedAt"`
	Pid               int           `json:"pid"`
	Path              string        `json:"path"`
	Args              []string      `json:"args"`
	Env               []string      `json:"env"`
}

// TaskError ...
type TaskError struct {
	Error string `json:"error"`
	Name  string `json:"name"`
}
