package controller

import (
	"app/db/repository"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	tasks, err := repository.GetTasks(ctx)
	if err != nil {
		return
	}
	for _, task := range tasks {
		log.Println("res:", task)
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]string{"message": "ok"}
	json.NewEncoder(w).Encode(response)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "taskID")
	taskIDInt, err := strconv.Atoi(taskID)
	if err != nil {
		return
	}

	ctx := context.Background()
	task, err := repository.GetTask(ctx, taskIDInt)
	if err != nil {
		return
	}
	log.Println("res:", task)

	w.WriteHeader(http.StatusOK)
	response := map[string]string{"message": "ok"}
	json.NewEncoder(w).Encode(response)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {

}

func UpdateTask(w http.ResponseWriter, r *http.Request) {

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {

}
