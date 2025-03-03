package database

import "fmt"

func RunMigrations() {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        email VARCHAR(100) UNIQUE NOT NULL,
        password TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`

	_, err := DB.Exec(query)
	if err != nil {
		fmt.Println("❌ Greška pri kreiranju tabela:", err)
	} else {
		fmt.Println("✅ Tabela 'users' je kreirana ili već postoji!")
	}
}
