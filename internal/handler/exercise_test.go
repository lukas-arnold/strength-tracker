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

func TestAddExerciseGet(t *testing.T) {

	h := testHandler(t)

	req :=
		httptest.NewRequest(
			"GET",
			"/exercise/add",
			nil,
		)

	rec :=
		httptest.NewRecorder()

	h.HandleAddExerciseGet(
		rec,
		req,
	)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"got %d",
			rec.Code,
		)
	}
}

func TestAddExercisePost(t *testing.T) {

	h := testHandler(t)

	form := url.Values{}

	form.Set(
		"name",
		"Bench Press",
	)

	form.Set(
		"muscleGroup",
		"Chest",
	)

	form.Set(
		"machine",
		"Barbell",
	)

	req :=
		httptest.NewRequest(
			"POST",
			"/exercise/add",
			strings.NewReader(
				form.Encode(),
			),
		)

	req.Header.Set(
		"Content-Type",
		"application/x-www-form-urlencoded",
	)

	rec :=
		httptest.NewRecorder()

	h.HandleAddExercisePost(
		rec,
		req,
	)

	if rec.Code != http.StatusFound {
		t.Fatalf(
			"got %d",
			rec.Code,
		)
	}

	exercises, err :=
		h.store.GetExercises()

	if err != nil {
		t.Fatal(err)
	}

	if len(exercises) != 1 {
		t.Fatal(
			"exercise missing",
		)
	}

	if exercises[0].Name != "Bench Press" {
		t.Fatal(
			"wrong exercise",
		)
	}
}

func TestEditExercise(t *testing.T) {

	h := testHandler(t)

	h.store.AddExercise(
		models.ExerciseInput{
			Name: "Bench",
		},
	)

	exercises, _ :=
		h.store.GetExercises()

	req :=
		httptest.NewRequest(
			"GET",
			"/exercise/edit",
			nil,
		)

	req.SetPathValue(
		"id",
		strconv.FormatInt(
			exercises[0].Id,
			10,
		),
	)

	rec :=
		httptest.NewRecorder()

	h.HandleEditExercise(
		rec,
		req,
	)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"got %d",
			rec.Code,
		)
	}
}

func TestEditExerciseInvalidID(t *testing.T) {

	h := testHandler(t)

	req :=
		httptest.NewRequest(
			"GET",
			"/exercise/edit",
			nil,
		)

	req.SetPathValue(
		"id",
		"abc",
	)

	rec :=
		httptest.NewRecorder()

	h.HandleEditExercise(
		rec,
		req,
	)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf(
			"got %d",
			rec.Code,
		)
	}
}

func TestSaveExercise(t *testing.T) {

	h := testHandler(t)

	h.store.AddExercise(
		models.ExerciseInput{
			Name: "Bench",
		},
	)

	exercises, _ :=
		h.store.GetExercises()

	form := url.Values{}

	form.Set(
		"name",
		"Squat",
	)

	form.Set(
		"muscleGroup",
		"Legs",
	)

	form.Set(
		"machine",
		"Barbell",
	)

	req :=
		httptest.NewRequest(
			"POST",
			"/exercise/save",
			strings.NewReader(
				form.Encode(),
			),
		)

	req.Header.Set(
		"Content-Type",
		"application/x-www-form-urlencoded",
	)

	req.SetPathValue(
		"id",
		strconv.FormatInt(
			exercises[0].Id,
			10,
		),
	)

	rec :=
		httptest.NewRecorder()

	h.HandleSaveExercise(
		rec,
		req,
	)

	if rec.Code != http.StatusFound {
		t.Fatalf(
			"got %d",
			rec.Code,
		)
	}

	updated, _ :=
		h.store.GetExercises()

	if updated[0].Name != "Squat" {
		t.Fatal(
			"exercise not updated",
		)
	}
}

func TestSaveExerciseInvalidID(t *testing.T) {

	h := testHandler(t)

	form := url.Values{}

	form.Set(
		"name",
		"Squat",
	)

	form.Set(
		"muscleGroup",
		"Legs",
	)

	form.Set(
		"machine",
		"Barbell",
	)

	req :=
		httptest.NewRequest(
			"POST",
			"/exercise/save",
			strings.NewReader(
				form.Encode(),
			),
		)

	req.Header.Set(
		"Content-Type",
		"application/x-www-form-urlencoded",
	)

	req.SetPathValue(
		"id",
		"abc",
	)

	rec :=
		httptest.NewRecorder()

	h.HandleSaveExercise(
		rec,
		req,
	)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf(
			"got %d",
			rec.Code,
		)
	}
}

func TestDeleteExercise(t *testing.T) {

	h := testHandler(t)

	err := h.store.AddExercise(
		models.ExerciseInput{
			Name: "Squat",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	exercises, _ :=
		h.store.GetExercises()

	req :=
		httptest.NewRequest(
			"GET",
			"/exercise/delete",
			nil,
		)

	req.SetPathValue(
		"id",
		strconv.FormatInt(
			exercises[0].Id,
			10,
		),
	)

	rec :=
		httptest.NewRecorder()

	h.HandleDeleteExercise(
		rec,
		req,
	)

	if rec.Code != http.StatusFound {
		t.Fatalf(
			"got %d",
			rec.Code,
		)
	}

	result, _ :=
		h.store.GetExercises()

	if len(result) != 0 {
		t.Fatal(
			"exercise not deleted",
		)
	}
}

func TestDeleteExerciseInvalidID(t *testing.T) {

	h := testHandler(t)

	req :=
		httptest.NewRequest(
			"GET",
			"/exercise/delete",
			nil,
		)

	req.SetPathValue(
		"id",
		"abc",
	)

	rec :=
		httptest.NewRecorder()

	h.HandleDeleteExercise(
		rec,
		req,
	)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf(
			"got %d",
			rec.Code,
		)
	}
}

func TestHistory(t *testing.T) {

	h := testHandler(t)

	h.store.AddExercise(
		models.ExerciseInput{
			Name: "Bench",
		},
	)

	exercises, _ :=
		h.store.GetExercises()

	req :=
		httptest.NewRequest(
			"GET",
			"/exercise/history",
			nil,
		)

	req.SetPathValue(
		"id",
		strconv.FormatInt(
			exercises[0].Id,
			10,
		),
	)

	rec :=
		httptest.NewRecorder()

	h.HandleHistory(
		rec,
		req,
	)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"got %d",
			rec.Code,
		)
	}
}

func TestHistoryInvalidID(t *testing.T) {

	h := testHandler(t)

	req :=
		httptest.NewRequest(
			"GET",
			"/exercise/history",
			nil,
		)

	req.SetPathValue(
		"id",
		"abc",
	)

	rec :=
		httptest.NewRecorder()

	h.HandleHistory(
		rec,
		req,
	)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf(
			"got %d",
			rec.Code,
		)
	}
}
