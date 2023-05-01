// Package status (of task)
package status

// Status (of a task)
type Status int64

// Enums of Status
const (
	Waiting   Status = 0
	Started   Status = 1
	Finished  Status = 2
	Failed    Status = 3
	Cancelled Status = 4
)
