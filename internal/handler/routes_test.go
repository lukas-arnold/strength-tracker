package handler

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/lukas-arnold/strength-tracker/internal/storage"
)

func TestRoutes(t *testing.T) {

	store := storage.New(
		filepath.Join(
			t.TempDir(),
			"test.json",
		),
	)

	h := New(store)

	mux := http.NewServeMux()

	RegisterRoutes(
		mux,
		h,
	)

	tests := []struct {
		method string
		path   string
	}{
		{
			method: "GET",
			path:   "/",
		},

		{
			method: "GET",
			path:   "/web",
		},

		{
			method: "GET",
			path:   "/service-worker.js",
		},

		{
			method: "GET",
			path:   "/exercise/add",
		},
		{
			method: "POST",
			path:   "/exercise/add",
		},
		{
			method: "GET",
			path:   "/exercise/edit/1",
		},
		{
			method: "POST",
			path:   "/exercise/save/1",
		},
		{
			method: "GET",
			path:   "/exercise/delete/1",
		},
		{
			method: "GET",
			path:   "/exercise/history/1",
		},

		{
			method: "GET",
			path:   "/strength/add/1",
		},
		{
			method: "POST",
			path:   "/strength/add/1",
		},
		{
			method: "GET",
			path:   "/strength/edit/1",
		},
		{
			method: "POST",
			path:   "/strength/save/1",
		},
		{
			method: "GET",
			path:   "/strength/delete/1",
		},
	}

	for _, test := range tests {

		req := httptest.NewRequest(
			test.method,
			test.path,
			nil,
		)

		rec := httptest.NewRecorder()

		mux.ServeHTTP(
			rec,
			req,
		)

		if rec.Code == http.StatusNotFound {
			t.Fatalf(
				"route missing: %s %s",
				test.method,
				test.path,
			)
		}
	}
}
