package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thomasreinecke/go-rest/internal/sushi"
)

// GetRolls returns all known rolls
func GetRolls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Request-Method", "GET")
	json.NewEncoder(w).Encode(sushi.GetMenu())
}

// GetRoll returns a single Roll
func GetRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range sushi.GetMenu() {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// CreateRoll creates a New Roll on the sushi menu
func CreateRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newRoll sushi.Roll
	json.NewDecoder(r.Body).Decode(&newRoll)
	newRoll = sushi.AddRoll(newRoll)
	json.NewEncoder(w).Encode(newRoll)
}

// UpdateRoll updates a given Roll on the sushi menu
func UpdateRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	var changedRoll sushi.Roll
	var err error

	json.NewDecoder(r.Body).Decode(&changedRoll)
	changedRoll, err = sushi.UpdateRoll(id, changedRoll)
	if err != nil {
		http.Error(w, "Forbidden", http.StatusForbidden)
	} else {
		json.NewEncoder(w).Encode(changedRoll)
	}
}

// DeleteRoll deletes a Roll given by its id from the sushi menu
func DeleteRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	err := sushi.DeleteRoll(id)
	if err != nil {
		http.Error(w, "Forbidden", http.StatusForbidden)
	} else {
		json.NewEncoder(w).Encode(sushi.GetMenu())
	}
}
