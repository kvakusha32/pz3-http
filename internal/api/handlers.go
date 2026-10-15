package api

import (
	"encoding/json"
	"errors"
	"mime"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/kvakusha32/pz3-http/internal/storage"
)

type Handlers struct {
	Store *storage.MemoryStore
}

func NewHandlers(store *storage.MemoryStore) *Handlers {
	return &Handlers{Store: store}
}

func (h *Handlers) ListTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.Store.List()
	q := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("q")))

	if q != "" {
		filtered := make([]*storage.Task, 0)
		for _, t := range tasks {
			if strings.Contains(strings.ToLower(t.Title), q) {
				filtered = append(filtered, t)
			}
		}
		tasks = filtered
	}

	JSON(w, http.StatusOK, tasks)
}

type createTaskRequest struct {
	Title string `json:"title"`
}

func (h *Handlers) CreateTask(w http.ResponseWriter, r *http.Request) {
	mediaType, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil || mediaType != "application/json" {
		BadRequest(w, "Content-Type must be application/json")
		return
	}

	var req createTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, "invalid json")
		return
	}

	req.Title = strings.TrimSpace(req.Title)
	if req.Title == "" {
		BadRequest(w, "title is required")
		return
	}
	if n := utf8.RuneCountInString(req.Title); n < 3 || n > 140 {
		JSON(w, http.StatusUnprocessableEntity,
			ErrorResponse{Error: "title must be 3..140 characters"})
		return
	}
	t := h.Store.Create(req.Title)
	JSON(w, http.StatusCreated, t)
}

func (h *Handlers) GetTask(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) != 2 || parts[0] != "tasks" {
		NotFound(w, "invalid path")
		return
	}

	id, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil || id < 1 {
		BadRequest(w, "invalid id")
		return
	}

	t, err := h.Store.Get(id)
	if errors.Is(err, storage.ErrNotFound) {
		NotFound(w, "task not found")
		return
	}
	if err != nil {
		Internal(w, "unexpected error")
		return
	}

	JSON(w, http.StatusOK, t)
}
