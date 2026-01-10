package database

import (
	"database/sql"
	"event-booking-api/pkg/config"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	dsn := config.Get("DB_URL")

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("Failed to Connect to database:", err)
	}

	DB = db
	log.Println("Database connected successfully")

}
