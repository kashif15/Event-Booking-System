package auth

import (
	"errors"
	"event-booking-api/pkg/database"
	"time"
)

func SaveRefreshToken(userID int64, token string, expiresAt time.Time) error {
	query := `INSERT INTO refresh_tokens (user_id, token, expires_at) VALUES ($1, $2, $3)`

	_, err := database.DB.Exec(query, userID, token, expiresAt)

	if err != nil {
		return errors.New("couldn't save refresh token")
	}

	return err
}

func GetRefreshToken(token string) (*RefereshToken, error) {
	query := `SELECT id, user_id, token, expires_at FROM refresh_tokens WHERE token = $1`

	var rt RefereshToken
	err := database.DB.QueryRow(query, token).Scan(&rt.ID, &rt.UserID, &rt.Token, &rt.ExpiresAt)

	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	if time.Now().After(rt.ExpiresAt) {
		return nil, errors.New("refresh token expired")
	}

	return &rt, nil
}

func DeleteRefreshToken(token string) error {
	_, err := database.DB.Exec(
		`DELETE FROM refresh_tokens WHERE token = $1`,
		token,
	)
	return err
}

