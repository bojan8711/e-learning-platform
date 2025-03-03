package handlers

import (
	"e-learning-platform/database"
	"e-learning-platform/models"
	"encoding/json"
	"net/http"

	"github.com/bojan8711/e-learning-platform/database"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Ne vazeci podaci ", http.StatusBadRequest)
		return
	}

	hashedPassword, err := models.HashPassword(user.Password)

	if err != nil {
		http.Error(w, "Greška pri hashiranju lozinke", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	_, err = database.DB.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)
	if err != nil {
		http.Error(w, "Greška pri čuvanju korisnika", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Uspešna registracija"})

}
