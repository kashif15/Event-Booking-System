package user

import (
	"database/sql"
	"errors"
	"event-booking-api/pkg/database"
)

func Create(u *User) error {
	query := `INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id, created_at`

	err := database.DB.QueryRow(query, u.Name, u.Email, u.Password, u.Role). Scan(&u.ID, &u.CreatedAt)

	return err
}

func GetByEmail(email string) (*User, error) {
	query := `SELECT id, name, email, password, role, created_at From users WHERE email = $1`

	var u User

	err := database.DB.QueryRow(query, email).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Role, &u.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func GetByID(id int64) (*User, error) {
	query := `
		SELECT id, name, email, role, created_at
		FROM users
		WHERE id = $1
	`

	var user User

	err := database.DB.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Role,
		&user.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

