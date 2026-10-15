package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kvakusha32/pz3-http/internal/storage"
)

func TestCreateAndGetTask(t *testing.T) {
	h := NewHandlers(storage.NewMemoryStore())

	create := httptest.NewRequest(
		http.MethodPost,
		"/tasks",
		strings.NewReader(`{"title":"Buy milk"}`),
	)
	create.Header.Set("Content-Type", "application/json")

	created := httptest.NewRecorder()
	h.CreateTask(created, create)
	if created.Code != http.StatusCreated {
		t.Fatalf("create status = %d, want 201: %s",
			created.Code, created.Body.String())
	}

	get := httptest.NewRecorder()
	h.GetTask(get, httptest.NewRequest(http.MethodGet, "/tasks/1", nil))
	if get.Code != http.StatusOK ||
		!strings.Contains(get.Body.String(), `"title":"Buy milk"`) {
		t.Fatalf("get response = %d %s",
			get.Code, get.Body.String())
	}
}

func TestInvalidAndMissingTask(t *testing.T) {
	h := NewHandlers(storage.NewMemoryStore())

	invalid := httptest.NewRecorder()
	h.GetTask(invalid, httptest.NewRequest(http.MethodGet, "/tasks/abc", nil))
	if invalid.Code != http.StatusBadRequest {
		t.Fatalf("invalid id status = %d, want 400", invalid.Code)
	}

	missing := httptest.NewRecorder()
	h.GetTask(missing, httptest.NewRequest(http.MethodGet, "/tasks/9999", nil))
	if missing.Code != http.StatusNotFound {
		t.Fatalf("missing id status = %d, want 404", missing.Code)
	}
}
