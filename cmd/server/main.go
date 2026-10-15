package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kvakusha32/pz3-http/internal/api"
	"github.com/kvakusha32/pz3-http/internal/storage"
)

func main() {
	store := storage.NewMemoryStore()
	h := api.NewHandlers(store)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		api.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	mux.HandleFunc("GET /tasks", h.ListTasks)
	mux.HandleFunc("POST /tasks", h.CreateTask)
	mux.HandleFunc("GET /tasks/", h.GetTask)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port
	log.Println("listening on", addr)

	if err := http.ListenAndServe(addr, api.Logging(mux)); err != nil {
		log.Fatal(err)
	}
}
