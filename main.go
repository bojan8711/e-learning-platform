package main

import (
	"log"
	"net/http"

	"github.com/bojan8711/e-learning-platform/database"
	"github.com/bojan8711/e-learning-platform/handlers"

	"github.com/gorilla/mux"
)

func main() {
	database.ConnectDB()
	database.CreateTables()

	r := mux.NewRouter()
	r.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")

	log.Println("✅ Server pokrenut na portu 8080")
	http.ListenAndServe(":8080", r)
}
