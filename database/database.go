package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL drajver
)

func ConnectDB() (*sql.DB, error) {
	// Definiši string za konekciju prema bazi
	connStr := "user=postgres password=yourpassword dbname=yourdbname sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("Neuspešno povezivanje sa bazom: %v", err)
	}

	// Testiraj konekciju sa bazom
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Neuspešno povezivanje sa bazom: %v", err)
	}

	return db, nil
}
