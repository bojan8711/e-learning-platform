package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

func ConnectDB() error {
	// Učitaj promenljive iz .env fajla
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Neuspešno učitavanje .env fajla", err)
	}

	// Učitaj vrednosti iz environment varijabli
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	// Definiši string za konekciju prema bazi
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", dbUser, dbPassword, dbName, dbSSLMode)
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("Neuspešno povezivanje sa bazom: %v", err)
	}

	// Proveri da li je baza dostupna
	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("Neuspešno povezivanje sa bazom: %v", err)
	}

	fmt.Println("Uspešno povezivanje sa bazom.")
	return nil
}
