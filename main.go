package main

import (
	"app/db/repository"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		taskId := r.URL.Path[len("/tasks/"):]
		taskIdInt, err := strconv.Atoi(taskId)
		if err != nil {
			return
		}

		ctx := context.Background()
		task, err := repository.GetTask(ctx, taskIdInt)
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
