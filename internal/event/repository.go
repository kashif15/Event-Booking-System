package event

import (
	"database/sql"
	"errors"
	"event-booking-api/pkg/database"
	"strconv"
	"time"
)

func Create(event *Event) error {

	query := `INSERT INTO events (title, description,location, event_time, capacity, created_by, status )
			VALUES ($1, $2, $3, $4, $5, $6, 'ACTIVE') 
			RETURNING id, created_at
	`

	err := database.DB.QueryRow(
		query,
		event.Title,
		event.Description,
		event.Location,
		event.EventTime,
		event.Capacity,
		event.CreatedBy,
	).Scan(&event.ID, &event.CreatedAt)

	return err
}

func GetByID(id int64) (*Event, error) {
	query := `
		SELECT id, title, description, location, event_time, capacity, created_by, status, created_at
		FROM events
		WHERE id = $1
	`

	var e Event

	err := database.DB.QueryRow(query, id).Scan(
		&e.ID,
		&e.Title,
		&e.Description,
		&e.Location,
		&e.EventTime,
		&e.Capacity,
		&e.CreatedBy,
		&e.Status,
		&e.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("event not found")
	}

	if err != nil {
		return nil, err
	}

	return &e, nil
}

func GetAllEvents() ([]Event, error) {

	query := `
		SELECT id, title, description, location, event_time,
		       capacity, created_by, status, created_at
		FROM events
		WHERE status = 'ACTIVE'
		ORDER BY event_time
	`

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var e Event

		err := rows.Scan(
			&e.ID,
			&e.Title,
			&e.Description,
			&e.Location,
			&e.EventTime,   
			&e.Capacity,    
			&e.CreatedBy,   
			&e.Status,      
			&e.CreatedAt, 
		)

		if err != nil {
			return nil, err
		}

		events = append(events, e)
	}

	return events, nil
}


func Delete(id int64) error {
	query := `DELETE FROM events WHERE id = $1`
	_, err := database.DB.Exec(query, id)
	return err
}

func GetWillFilter(page int, limit int, status string, createdBy *int64, fromDate *time.Time,
	search *string)([]Event, error) {


	offset := (page - 1) * limit

	query := `
		SELECT id, title, description, location, event_time, capacity,
		       created_by, status, created_at
		FROM events
		WHERE 1=1
	`

	args := []interface{}{}
	argPos := 1

	if status != "" {
		query += " AND status =$" + strconv.Itoa(argPos)
		args = append(args, status)
		argPos++
	}

	if createdBy != nil {
		query += " AND created_by = $" + strconv.Itoa(argPos)
		args = append(args, *createdBy)
		argPos++
	}

	if fromDate != nil {
		query += " AND event_time >= $" + strconv.Itoa(argPos)
		args = append(args, *fromDate)
		argPos++
	}

	if search != nil && *search != "" {
		query += `
		AND (
			title ILIKE $` + strconv.Itoa(argPos) + `
			OR location ILIKE $` + strconv.Itoa(argPos+1) + `
		)
	   `

		searchPattern := "%" + *search + "%"

		args = append(args, searchPattern, searchPattern)
		argPos += 2
	}

		query += `
		ORDER BY event_time
		LIMIT $` + strconv.Itoa(argPos) + `
		OFFSET $` + strconv.Itoa(argPos+1)

	args = append(args, limit, offset)

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var e Event
		err := rows.Scan(
			&e.ID,
			&e.Title,
			&e.Description,
			&e.Location,
			&e.EventTime,
			&e.Capacity,
			&e.CreatedBy,
			&e.Status,
			&e.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}

	return events, nil

}
