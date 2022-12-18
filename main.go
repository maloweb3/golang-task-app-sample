package main

import (
	"app/db/repository"
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		taskId := r.URL.Path[len("/tasks/"):]
		log.Println("taskId:", taskId)

		task, err := repository.GetTask(taskId)
		if err != nil {
			return
		}
		log.Println("task:", task)

		w.WriteHeader(http.StatusOK)
		response := map[string]string{"message": "ok"}
		json.NewEncoder(w).Encode(response)
	})
	http.ListenAndServe(":8080", nil)
}
