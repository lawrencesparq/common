package utils

import (
	"database/sql"
	"log"
)

func InitStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	log.Println("Database connection successful")
}
