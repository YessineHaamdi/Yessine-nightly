package api

import (
	"encoding/json"
	"net/http"

	"tekton-backend/internal/tekton"
)

// jsonResponse is a small helper for consistent responses.
func jsonResponse(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

// CreatePipelineHandler creates the todo pipeline in the provided namespace (default: default)
func CreatePipelineHandler(w http.ResponseWriter, r *http.Request) {
	ns := r.URL.Query().Get("namespace")
	if ns == "" {
		ns = "default"
	}

	client, err := tekton.NewClient()
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, map[string]string{"error": "failed to create tekton client"})
		return
	}

	p, err := tekton.CreateTodoPipeline(client, ns)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	jsonResponse(w, http.StatusCreated, map[string]string{"pipeline": p.Name})
}

// DeletePipelineHandler deletes the todo pipeline in the provided namespace (default: default)
func DeletePipelineHandler(w http.ResponseWriter, r *http.Request) {
	ns := r.URL.Query().Get("namespace")
	if ns == "" {
		ns = "default"
	}

	client, err := tekton.NewClient()
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, map[string]string{"error": "failed to create tekton client"})
		return
	}

	if err := tekton.DeleteTodoPipeline(client, ns); err != nil {
		jsonResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	jsonResponse(w, http.StatusOK, map[string]string{"deleted": "todo-pipeline"})
}
