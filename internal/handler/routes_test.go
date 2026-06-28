package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutes(t *testing.T) {
	h := testHandler(t)
	mux := http.NewServeMux()
	RegisterRoutes(mux, h)

	tests := []struct {
		name   string
		method string
		path   string
	}{
		{"Home", http.MethodGet, "/"},

		{"StaticWeb", http.MethodGet, "/web/"},
		{"ServiceWorker", http.MethodGet, "/service-worker.js"},

		{"ExerciseAddGet", http.MethodGet, "/exercise/add"},
		{"ExerciseAddPost", http.MethodPost, "/exercise/add"},
		{"ExerciseEdit", http.MethodGet, "/exercise/edit/1"},
		{"ExerciseSave", http.MethodPost, "/exercise/save/1"},
		{"ExerciseDelete", http.MethodGet, "/exercise/delete/1"},
		{"ExerciseHistory", http.MethodGet, "/exercise/history/1"},

		{"StrengthAddGet", http.MethodGet, "/strength/add/1"},
		{"StrengthAddPost", http.MethodPost, "/strength/add/1"},
		{"StrengthEdit", http.MethodGet, "/strength/edit/1"},
		{"StrengthSave", http.MethodPost, "/strength/save/1"},
		{"StrengthDelete", http.MethodGet, "/strength/delete/1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			rec := httptest.NewRecorder()

			mux.ServeHTTP(rec, req)

			if rec.Code == http.StatusNotFound {
				t.Errorf("route missing or returned 404: %s %s", tt.method, tt.path)
			}
		})
	}
}
