package handler

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lukas-arnold/strength-tracker/internal/models"
	"github.com/lukas-arnold/strength-tracker/internal/storage"
)

func testHandler(t *testing.T) *Handler {
	t.Helper()
	store := storage.New(filepath.Join(t.TempDir(), "test.json"))
	return New(store)
}

func TestHandleView(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		h := testHandler(t)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		h.HandleView(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rec.Code)
		}
	})

	t.Run("with exercises", func(t *testing.T) {
		h := testHandler(t)
		err := h.store.AddExercise(models.ExerciseInput{
			Name:        "Bench Press",
			MuscleGroup: "Chest",
			Machine:     "Barbell",
		})
		if err != nil {
			t.Fatal(err)
		}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		h.HandleView(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rec.Code)
		}
		if !strings.Contains(rec.Body.String(), "Bench Press") {
			t.Fatal("exercise name not rendered in response body")
		}
	})
}

func TestTemplateFuncs(t *testing.T) {
	funcs := getTemplateFuncs()

	t.Run("map existence", func(t *testing.T) {
		if funcs == nil {
			t.Fatal("expected template funcs")
		}
		for name, fn := range funcs {
			if fn == nil {
				t.Errorf("template func %s is nil", name)
			}
		}
	})

	t.Run("formatting logic", func(t *testing.T) {
		load := funcs["formatLoad"].(func(float64) string)
		reps := funcs["formatRepetitions"].(func(int64) string)

		if load(100.5) != "100.5" {
			t.Error("formatLoad failed")
		}
		if reps(8) != "8" {
			t.Error("formatRepetitions failed")
		}
	})

	t.Run("translation", func(t *testing.T) {
		tFunc := funcs["T"].(func(string) string)
		if result := tFunc("test_key"); result == "" {
			t.Fatal("expected translation result")
		}
	})
}

func TestRenderTemplate(t *testing.T) {
	h := testHandler(t)

	t.Run("render without data", func(t *testing.T) {
		rec := httptest.NewRecorder()
		h.renderTemplate(rec, "templates/exercise/add.html", nil)
		if rec.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rec.Code)
		}
	})

	t.Run("render with data", func(t *testing.T) {
		rec := httptest.NewRecorder()
		h.renderTemplate(rec, "templates/exercise/edit.html", models.Exercise{})
		if rec.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rec.Code)
		}
	})
}
