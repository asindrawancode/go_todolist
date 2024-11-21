package handlers

import (
	"encoding/json"
	"go_todolist/models"
	"net/http"
)

func GetToDoList(w http.ResponseWriter, r *http.Request) {
	var todos []models.ToDo
	user, _ := r.Context().Value("user").(*models.User)

	if err := models.DB.Where("user_id = ?", user.ID).Find(&todos).Error; err != nil {
		http.Error(w, "Error fetching todos", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todos)
}

func CreateToDo(w http.ResponseWriter, r *http.Request) {
	var todo models.ToDo
	user, _ := r.Context().Value("user").(*models.User)

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo.UserID = user.ID
	if err := models.DB.Save(&todo).Error; err != nil {
		http.Error(w, "Error creating todo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
