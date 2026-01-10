package event

import "time"

type Event struct {
	ID int64
	Title string
	Description string
	Location string
	EventTime time.Time
	Capacity int
	CreatedBy int64
	Status string
	CreatedAt time.Time
}

