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

func TestAddStrengthGet(t *testing.T) {

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
			"/strength/add",
			nil,
		)

	req.SetPathValue(
		"exerciseId",
		strconv.FormatInt(
			exercises[0].Id,
			10,
		),
	)

	rec :=
		httptest.NewRecorder()

	h.HandleAddStrengthGet(
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

func TestAddStrengthGetInvalidID(t *testing.T) {

	h := testHandler(t)

	req :=
		httptest.NewRequest(
			"GET",
			"/strength/add",
			nil,
		)

	req.SetPathValue(
		"exerciseId",
		"abc",
	)

	rec :=
		httptest.NewRecorder()

	h.HandleAddStrengthGet(
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

func TestAddStrengthPost(t *testing.T) {

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
		"date",
		"2024-01-01",
	)

	form.Set(
		"load",
		"100",
	)

	form.Set(
		"repetitions",
		"5",
	)

	req :=
		httptest.NewRequest(
			"POST",
			"/strength/add",
			strings.NewReader(
				form.Encode(),
			),
		)

	req.Header.Set(
		"Content-Type",
		"application/x-www-form-urlencoded",
	)

	req.SetPathValue(
		"exerciseId",
		strconv.FormatInt(
			exercises[0].Id,
			10,
		),
	)

	rec :=
		httptest.NewRecorder()

	h.HandleAddStrengthPost(
		rec,
		req,
	)

	if rec.Code != http.StatusFound {
		t.Fatalf(
			"got %d",
			rec.Code,
		)
	}

	exercise, _ :=
		h.store.GetExercise(
			exercises[0].Id,
		)

	if len(exercise.StrengthHistory) != 1 {
		t.Fatal(
			"strength missing",
		)
	}

	if exercise.StrengthHistory[0].Load != 100 {
		t.Fatalf(
			"got %f",
			exercise.StrengthHistory[0].Load,
		)
	}
}

func TestAddStrengthPostInvalidLoad(t *testing.T) {

	h := testHandler(t)

	req :=
		httptest.NewRequest(
			"POST",
			"/strength/add",
			strings.NewReader(
				"load=test&repetitions=5",
			),
		)

	req.Header.Set(
		"Content-Type",
		"application/x-www-form-urlencoded",
	)

	req.SetPathValue(
		"exerciseId",
		"1",
	)

	rec :=
		httptest.NewRecorder()

	h.HandleAddStrengthPost(
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

func TestEditStrength(t *testing.T) {

	h := testHandler(t)

	h.store.AddExercise(
		models.ExerciseInput{
			Name: "Bench",
		},
	)

	exercises, _ :=
		h.store.GetExercises()

	h.store.AddStrength(
		exercises[0].Id,
		models.StrengthInput{
			Date: "2024",
			Load: 50,
		},
	)

	exercise, _ :=
		h.store.GetExercise(
			exercises[0].Id,
		)

	id :=
		exercise.StrengthHistory[0].Id

	req :=
		httptest.NewRequest(
			"GET",
			"/strength/edit",
			nil,
		)

	req.SetPathValue(
		"id",
		strconv.FormatInt(
			id,
			10,
		),
	)

	rec :=
		httptest.NewRecorder()

	h.HandleEditStrength(
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

func TestEditStrengthInvalidID(t *testing.T) {

	h := testHandler(t)

	req :=
		httptest.NewRequest(
			"GET",
			"/strength/edit",
			nil,
		)

	req.SetPathValue(
		"id",
		"abc",
	)

	rec :=
		httptest.NewRecorder()

	h.HandleEditStrength(
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

func TestSaveStrength(t *testing.T) {

	h := testHandler(t)

	h.store.AddExercise(
		models.ExerciseInput{
			Name: "Bench",
		},
	)

	exercises, _ :=
		h.store.GetExercises()

	h.store.AddStrength(
		exercises[0].Id,
		models.StrengthInput{
			Date: "2024",
			Load: 50,
		},
	)

	exercise, _ :=
		h.store.GetExercise(
			exercises[0].Id,
		)

	id :=
		exercise.StrengthHistory[0].Id

	form := url.Values{}

	form.Set(
		"date",
		"2025",
	)

	form.Set(
		"load",
		"100",
	)

	form.Set(
		"repetitions",
		"10",
	)

	req :=
		httptest.NewRequest(
			"POST",
			"/strength/save",
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
			id,
			10,
		),
	)

	rec :=
		httptest.NewRecorder()

	h.HandleSaveStrength(
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
		h.store.GetStrength(id)

	if updated.Load != 100 {
		t.Fatalf(
			"got %f",
			updated.Load,
		)
	}
}

func TestSaveStrengthInvalidID(t *testing.T) {

	h := testHandler(t)

	form := url.Values{}

	form.Set(
		"date",
		"2025",
	)

	form.Set(
		"load",
		"100",
	)

	form.Set(
		"repetitions",
		"10",
	)

	req :=
		httptest.NewRequest(
			"POST",
			"/strength/save",
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

	h.HandleSaveStrength(
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

func TestDeleteStrength(t *testing.T) {

	h := testHandler(t)

	h.store.AddExercise(
		models.ExerciseInput{
			Name: "Bench",
		},
	)

	exercises, _ :=
		h.store.GetExercises()

	h.store.AddStrength(
		exercises[0].Id,
		models.StrengthInput{
			Date: "2024",
		},
	)

	exercise, _ :=
		h.store.GetExercise(
			exercises[0].Id,
		)

	id :=
		exercise.StrengthHistory[0].Id

	req :=
		httptest.NewRequest(
			"GET",
			"/strength/delete",
			nil,
		)

	req.SetPathValue(
		"id",
		strconv.FormatInt(
			id,
			10,
		),
	)

	rec :=
		httptest.NewRecorder()

	h.HandleDeleteStrength(
		rec,
		req,
	)

	if rec.Code != http.StatusFound {
		t.Fatalf(
			"got %d",
			rec.Code,
		)
	}

	strength, _ :=
		h.store.GetStrength(id)

	if strength.Id != 0 {
		t.Fatal(
			"strength not deleted",
		)
	}
}

func TestDeleteStrengthInvalidID(t *testing.T) {

	h := testHandler(t)

	req :=
		httptest.NewRequest(
			"GET",
			"/strength/delete",
			nil,
		)

	req.SetPathValue(
		"id",
		"abc",
	)

	rec :=
		httptest.NewRecorder()

	h.HandleDeleteStrength(
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
