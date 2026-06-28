package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/lukas-arnold/strength-tracker/internal/models"
)

func TestAddStrength(t *testing.T) {
	t.Run("GET add strength", func(t *testing.T) {
		h := testHandler(t)
		h.store.AddExercise(models.ExerciseInput{Name: "Bench"})
		exercises, _ := h.store.GetExercises()

		req := httptest.NewRequest(http.MethodGet, "/strength/add", nil)
		req.SetPathValue("exerciseId", strconv.FormatInt(exercises[0].Id, 10))
		rec := httptest.NewRecorder()

		h.HandleAddStrengthGet(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rec.Code)
		}
	})

	t.Run("POST add strength", func(t *testing.T) {
		h := testHandler(t)
		h.store.AddExercise(models.ExerciseInput{Name: "Bench"})
		exercises, _ := h.store.GetExercises()

		form := url.Values{"date": {"2024-01-01"}, "load": {"100"}, "repetitions": {"5"}}
		req := httptest.NewRequest(http.MethodPost, "/strength/add", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.SetPathValue("exerciseId", strconv.FormatInt(exercises[0].Id, 10))
		rec := httptest.NewRecorder()

		h.HandleAddStrengthPost(rec, req)

		if rec.Code != http.StatusFound {
			t.Fatalf("expected 302, got %d", rec.Code)
		}

		ex, _ := h.store.GetExercise(exercises[0].Id)
		if len(ex.StrengthHistory) != 1 || ex.StrengthHistory[0].Load != 100 {
			t.Fatal("strength not added correctly")
		}
	})

	t.Run("invalid IDs return 500", func(t *testing.T) {
		h := testHandler(t)
		req := httptest.NewRequest(http.MethodGet, "/strength/add", nil)
		req.SetPathValue("exerciseId", "abc")
		rec := httptest.NewRecorder()

		h.HandleAddStrengthGet(rec, req)
		if rec.Code != http.StatusInternalServerError {
			t.Errorf("expected 500, got %d", rec.Code)
		}
	})
}

func TestEditSaveStrength(t *testing.T) {
	h := testHandler(t)
	h.store.AddExercise(models.ExerciseInput{Name: "Bench"})

	ex, _ := h.store.GetExercises()
	exID := ex[0].Id

	h.store.AddStrength(exID, models.StrengthInput{Date: "2024", Load: 50})

	updatedEx, _ := h.store.GetExercise(exID)
	strengthID := updatedEx.StrengthHistory[0].Id

	t.Run("GET edit strength", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/strength/edit", nil)
		req.SetPathValue("id", strconv.FormatInt(strengthID, 10))
		rec := httptest.NewRecorder()

		h.HandleEditStrength(rec, req)
		if rec.Code != http.StatusOK {
			t.Errorf("expected 200, got %d", rec.Code)
		}
	})

	t.Run("POST save strength", func(t *testing.T) {
		form := url.Values{"date": {"2025"}, "load": {"100"}, "repetitions": {"10"}}
		req := httptest.NewRequest(http.MethodPost, "/strength/save", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.SetPathValue("id", strconv.FormatInt(strengthID, 10))
		rec := httptest.NewRecorder()

		h.HandleSaveStrength(rec, req)

		updated, _ := h.store.GetStrength(strengthID)
		if updated.Load != 100 {
			t.Errorf("expected load 100, got %f", updated.Load)
		}
	})
}

func TestDeleteStrength(t *testing.T) {
	h := testHandler(t)
	h.store.AddExercise(models.ExerciseInput{Name: "Bench"})

	ex, _ := h.store.GetExercises()
	exID := ex[0].Id

	h.store.AddStrength(exID, models.StrengthInput{Date: "2024"})

	updatedEx, _ := h.store.GetExercise(exID)
	id := updatedEx.StrengthHistory[0].Id

	req := httptest.NewRequest(http.MethodGet, "/strength/delete", nil)
	req.SetPathValue("id", strconv.FormatInt(id, 10))
	rec := httptest.NewRecorder()

	h.HandleDeleteStrength(rec, req)

	strength, _ := h.store.GetStrength(id)
	if strength.Id != 0 {
		t.Fatal("strength was not deleted")
	}
}
