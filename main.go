package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		taskId := r.URL.Path[len("/tasks/"):]
		log.Println("taskId:", taskId)

		w.WriteHeader(http.StatusOK)
		response := map[string]string{"message": "ok"}
		json.NewEncoder(w).Encode(response)
	})
	http.ListenAndServe(":8080", nil)
}
