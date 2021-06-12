package controllers

import (
	"encoding/json"
	"net/http"
)

// Welcome a user
func Home(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Welcome!")
}