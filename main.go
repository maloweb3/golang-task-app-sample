package main

import (
	"app/controller"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Create Router
	r := chi.NewRouter()
	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", controller.GetTasks)
		r.Get("/{taskID}", controller.GetTask)
		r.Post("/", controller.CreateTask)
		r.Patch("/{taskID}", controller.UpdateTask)
		r.Delete("/{taskID}", controller.DeleteTask)
	})

	// Start HTTP server.
	http.ListenAndServe(":8080", r)
}
