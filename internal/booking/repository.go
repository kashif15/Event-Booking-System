package booking

import (
	"database/sql"
	"errors"
	"event-booking-api/pkg/database"
)

func Create(userID, eventID int64) error {

	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	
	var exists int
	err = tx.QueryRow(
		`SELECT 1 FROM bookings WHERE user_id = $1 AND event_id = $2`,
		userID, eventID,
	).Scan(&exists)

	if err == nil {
		return errors.New("already booked")
	}

	if err != sql.ErrNoRows {
		return err
	}

	
	var capacity int
	var bookedCount int

	err = tx.QueryRow(
		`SELECT capacity FROM events WHERE id = $1 AND status = 'ACTIVE'`,
		eventID,
	).Scan(&capacity)

	if err != nil {
		return errors.New("event not available")
	}

	err = tx.QueryRow(
		`SELECT COUNT(*) FROM bookings WHERE event_id = $1 AND status = 'CONFIRMED'`,
		eventID,
	).Scan(&bookedCount)

	if err != nil {
		return err
	}

	if bookedCount >= capacity {
		return errors.New("event fully booked")
	}

	
	_, err = tx.Exec(
		`INSERT INTO bookings (user_id, event_id, status)
		 VALUES ($1, $2, 'CONFIRMED')`,
		userID, eventID,
	)

	if err != nil {
		return err
	}

	return tx.Commit()
}


func GetByUser(userID int64) ([]Booking, error) {

	query := `SELECT id, user_id, event_id, status, created_at
		 FROM bookings
		 WHERE user_id = $1`

	rows, err := database.DB.Query(query, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var booking []Booking

	for rows.Next() {
		var b Booking

		err := rows.Scan(&b.ID, &b.UserID, &b.EventID, &b.Status, &b.CreatedAt)

		if err != nil {
		return nil, err
		}

		booking = append(booking, b)

	}

	return booking, nil

}

func Cancel(userID, eventID int64) error {
	query := `UPDATE bookings SET status = 'CANCELED' WHERE user_id = $1 AND event_id = $2`

	result, err := database.DB.Exec(query, userID, eventID)

	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("booking not found")
	}

	return nil
}