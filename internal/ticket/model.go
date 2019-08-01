package ticket

import (
	"time"
)

type Ticket struct {
	ID              string
	TicketID        string
	AccountID       int64
	Status          Status
	Priority        Priority
	Title           string
	Description     string
	Assignee        Person
	InvolvedPersons []Person
	DueUntil        time.Time
	Created         time.Time
	Updated         time.Time
	Deleted         time.Time
}

type Status string
const (
	StatusOpen       = "open"
	StatusInProgress = "in_progress"
	StatusClosed     = "closed"
)

type Priority int32
const (
	PriorityLow = iota
	PriorityMedium
	PriorityHigh
)

type Person struct {
	UserID    string
	AccountID int64
}
