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

	store := storage.New(
		filepath.Join(
			t.TempDir(),
			"test.json",
		),
	)

	return New(store)
}

func TestHandleViewEmpty(t *testing.T) {

	h := testHandler(t)

	req :=
		httptest.NewRequest(
			"GET",
			"/",
			nil,
		)

	rec :=
		httptest.NewRecorder()

	h.HandleView(
		rec,
		req,
	)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected 200 got %d",
			rec.Code,
		)
	}
}

func TestHandleViewWithExercises(t *testing.T) {

	h := testHandler(t)

	err := h.store.AddExercise(
		models.ExerciseInput{
			Name:        "Bench Press",
			MuscleGroup: "Chest",
			Machine:     "Barbell",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	req :=
		httptest.NewRequest(
			"GET",
			"/",
			nil,
		)

	rec :=
		httptest.NewRecorder()

	h.HandleView(
		rec,
		req,
	)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected 200 got %d",
			rec.Code,
		)
	}

	if !strings.Contains(
		rec.Body.String(),
		"Bench Press",
	) {
		t.Fatal(
			"exercise not rendered",
		)
	}
}

func TestTemplateFuncs(t *testing.T) {

	funcs := getTemplateFuncs()

	if funcs == nil {
		t.Fatal("expected template funcs")
	}

	for name, fn := range funcs {

		if fn == nil {
			t.Fatalf(
				"template func %s is nil",
				name,
			)
		}
	}
}

func TestTemplateFuncsFormat(t *testing.T) {

	funcs := getTemplateFuncs()

	load :=
		funcs["formatLoad"].(func(float64) string)

	reps :=
		funcs["formatRepetitions"].(func(int64) string)

	if load(100.5) != "100.5" {
		t.Fatal(
			"formatLoad failed",
		)
	}

	if reps(8) != "8" {
		t.Fatal(
			"formatRepetitions failed",
		)
	}
}

func TestTemplateTranslationFunc(t *testing.T) {

	funcs := getTemplateFuncs()

	tFunc :=
		funcs["T"].(func(string) string)

	result := tFunc("something")

	if result == "" {
		t.Fatal(
			"expected translation result",
		)
	}
}

func TestRenderTemplate(t *testing.T) {

	h := testHandler(t)

	rec :=
		httptest.NewRecorder()

	h.renderTemplate(
		rec,
		"templates/exercise/add.html",
		nil,
	)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected 200 got %d",
			rec.Code,
		)
	}
}

func TestRenderTemplateWithData(t *testing.T) {

	h := testHandler(t)

	rec :=
		httptest.NewRecorder()

	h.renderTemplate(
		rec,
		"templates/exercise/edit.html",
		models.Exercise{},
	)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected 200 got %d",
			rec.Code,
		)
	}
}
