package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCreatePipelineHandler_ClientError(t *testing.T) {
	// Ensure NewClient fails by pointing KUBECONFIG to a definitely-invalid file and isolating HOME
	os.Setenv("KUBECONFIG", "C:/this/path/does/not/exist/kubeconfig")
	tmp := os.Getenv("HOME")
	os.Setenv("HOME", "C:/this/path/does/not/exist")
	defer os.Setenv("HOME", tmp)

	req := httptest.NewRequest("GET", "/pipeline/create", nil)
	w := httptest.NewRecorder()

	CreatePipelineHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusInternalServerError {
		t.Fatalf("expected status 500; got %d", res.StatusCode)
	}

	var body map[string]string
	_ = json.NewDecoder(res.Body).Decode(&body)
	if _, ok := body["error"]; !ok {
		t.Fatalf("expected error key in response body")
	}
}

func TestDeletePipelineHandler_ClientError(t *testing.T) {
	os.Setenv("KUBECONFIG", "C:/this/path/does/not/exist/kubeconfig")
	tmp := os.Getenv("HOME")
	os.Setenv("HOME", "C:/this/path/does/not/exist")
	defer os.Setenv("HOME", tmp)

	req := httptest.NewRequest("GET", "/pipeline/delete", nil)
	w := httptest.NewRecorder()

	DeletePipelineHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusInternalServerError {
		t.Fatalf("expected status 500; got %d", res.StatusCode)
	}

	var body map[string]string
	_ = json.NewDecoder(res.Body).Decode(&body)
	if _, ok := body["error"]; !ok {
		t.Fatalf("expected error key in response body")
	}
}
