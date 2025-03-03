package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() {
	var err error
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	DB, err = sqlx.Connect("postgress", connStr)
	if err != nil {
		log.Fatal("Neuspešno povezivanje sa bazom:", err)
	}
	fmt.Printf("✅ Uspostavljena konekcija sa bazom!")
}
