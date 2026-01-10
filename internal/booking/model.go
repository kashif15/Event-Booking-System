package booking

import "time"

type Booking struct {
	ID        int64
	UserID    int64
	EventID   int64
	Status 	   string
	CreatedAt time.Time 
}
