package handlers

import (
	"article/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func HandlerUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		Createusers(w, r)
	case http.MethodGet:
		getusers(w, r)
	case http.MethodPut:
		updateusers(w, r)
	}
}
func Createusers(w http.ResponseWriter, r *http.Request) {
	var newuser models.User
	err := json.NewDecoder(r.Body).Decode(&newuser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newuser.Id = len(models.Users) + 1
	fmt.Println(newuser)
	models.Users = append(models.Users, newuser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(newuser)
}
func getusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Users)
}

func updateusers(w http.ResponseWriter, r *http.Request) {
	var newusers models.User
	idString := r.URL.Query()["id"][0]
	id, err := strconv.Atoi(idString)
	fmt.Println(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	body := json.NewDecoder(r.Body).Decode(&newusers)

	if body != nil {
		http.Error(w, body.Error(), http.StatusBadRequest)
	}

	index := -1

	for i, user := range models.Users {
		if user.Id == id {
			index = i
			break
		}
	}

	if index != -1 {
		models.Users = append(models.Users[:index], models.Users[index+1:]...)
	}

	models.Users = append(models.Users, newusers)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Users)
}

