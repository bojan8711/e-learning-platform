package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL drajver
)

var DB *sql.DB // Globalna varijabla za bazu

// Funkcija za povezivanje sa bazom
func ConnectDB() error {
	// Definiši string za konekciju prema bazi
	connStr := "user=admin password=Ste11@2016 dbname=elearning sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("Neuspešno povezivanje sa bazom: %v", err)
	}

	// Testiraj konekciju sa bazom
	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("Neuspešno povezivanje sa bazom: %v", err)
	}

	fmt.Println("Povezan sa bazom!")
	return nil

}
