package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"example.com/pz3-http/internal/storage"
)

type Handlers struct {
	Store *storage.MemoryStore
}

func NewHandlers(store *storage.MemoryStore) *Handlers {
	return &Handlers{Store: store}
}

type createTaskRequest struct {
	Title string `json:"title"`
}

func (h *Handlers) ListTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.Store.List()
	q := strings.TrimSpace(r.URL.Query().Get("q"))
	if q != "" {
		filtered := tasks[:0]
		for _, t := range tasks {
			if strings.Contains(strings.ToLower(t.Title), strings.ToLower(q)) {
				filtered = append(filtered, t)
			}
		}
		tasks = filtered
	}
	JSON(w, http.StatusOK, tasks)
}

func (h *Handlers) CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "" && !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		BadRequest(w, "Content-Type must be application/json")
		return
	}

	var req createTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, "invalid json: "+err.Error())
		return
	}

	req.Title = strings.TrimSpace(req.Title)
	if len(req.Title) == 0 {
		BadRequest(w, "title is required")
		return
	}
	if len(req.Title) < 3 || len(req.Title) > 140 {
		JSON(w, 422, ErrorResponse{Error: "title length must be 3..140 characters"})
		return
	}

	t := h.Store.Create(req.Title)
	JSON(w, http.StatusCreated, t)
}

func (h *Handlers) GetTask(w http.ResponseWriter, r *http.Request) {
	id, ok := parseIDFromPath(r.URL.Path)
	if !ok {
		BadRequest(w, "invalid id")
		return
	}

	t, err := h.Store.Get(id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			NotFound(w, "task not found")
			return
		}
		Internal(w, "unexpected error")
		return
	}
	JSON(w, http.StatusOK, t)
}

func (h *Handlers) UpdateTaskDone(w http.ResponseWriter, r *http.Request) {
	id, ok := parseIDFromPath(r.URL.Path)
	if !ok {
		BadRequest(w, "invalid id")
		return
	}

	t, err := h.Store.UpdateDone(id, true)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			NotFound(w, "task not found")
			return
		}
		Internal(w, "unexpected error")
		return
	}
	JSON(w, http.StatusOK, t)
}

func (h *Handlers) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, ok := parseIDFromPath(r.URL.Path)
	if !ok {
		BadRequest(w, "invalid id")
		return
	}

	if err := h.Store.Delete(id); err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			NotFound(w, "task not found")
			return
		}
		Internal(w, "unexpected error")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func parseIDFromPath(path string) (int64, bool) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) != 2 {
		return 0, false
	}
	id, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return 0, false
	}
	return id, true
}
