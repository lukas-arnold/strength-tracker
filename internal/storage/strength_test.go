package storage

import (
	"path/filepath"
	"testing"

	"github.com/lukas-arnold/strength-tracker/internal/models"
)

func TestAddStrength(t *testing.T) {
	store := New(filepath.Join(t.TempDir(), "storage.json"))

	err := store.AddExercise(models.ExerciseInput{Name: "Bench"})
	if err != nil {
		t.Fatal(err)
	}

	exercises, _ := store.GetExercises()

	err = store.AddStrength(
		exercises[0].Id,
		models.StrengthInput{
			Date:        "2024-01-01",
			Load:        100,
			Repetitions: 5,
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	updated, _ := store.GetExercise(exercises[0].Id)

	if len(updated.StrengthHistory) != 1 {
		t.Fatal("strength not added")
	}
}

func TestAddStrengthInvalidExercise(t *testing.T) {
	store := New(filepath.Join(t.TempDir(), "storage.json"))

	err := store.AddStrength(
		999,
		models.StrengthInput{
			Date: "2024",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetStrength(t *testing.T) {
	store := New(filepath.Join(t.TempDir(), "storage.json"))

	store.AddExercise(models.ExerciseInput{Name: "Bench"})
	exercises, _ := store.GetExercises()

	store.AddStrength(
		exercises[0].Id,
		models.StrengthInput{
			Date: "2024-01-01",
			Load: 100,
		},
	)

	exercise, _ := store.GetExercise(exercises[0].Id)
	id := exercise.StrengthHistory[0].Id

	strength, err := store.GetStrength(id)
	if err != nil {
		t.Fatal(err)
	}

	if strength.Load != 100 {
		t.Fatalf("got %f", strength.Load)
	}
}

func TestGetStrengthNotFound(t *testing.T) {
	store := New(filepath.Join(t.TempDir(), "storage.json"))

	strength, err := store.GetStrength(123)
	if err != nil {
		t.Fatal(err)
	}

	if strength.Id != 0 {
		t.Fatal("expected empty strength")
	}
}

func TestGetLastStrength(t *testing.T) {
	store := New(filepath.Join(t.TempDir(), "storage.json"))

	store.AddExercise(models.ExerciseInput{Name: "Bench"})
	exercises, _ := store.GetExercises()

	store.AddStrength(
		exercises[0].Id,
		models.StrengthInput{
			Date: "2024-01-01",
			Load: 50,
		},
	)

	last, err := store.GetLastStrength(exercises[0].Id)
	if err != nil {
		t.Fatal(err)
	}

	if last.Load != 50 {
		t.Fatalf("got %f", last.Load)
	}
}

func TestGetLastStrengthEmpty(t *testing.T) {
	store := New(filepath.Join(t.TempDir(), "storage.json"))

	store.AddExercise(models.ExerciseInput{Name: "Bench"})
	exercises, _ := store.GetExercises()

	last, err := store.GetLastStrength(exercises[0].Id)
	if err != nil {
		t.Fatal(err)
	}

	if last.Id != 0 {
		t.Fatal("expected empty strength")
	}
}

func TestUpdateStrength(t *testing.T) {
	store := New(filepath.Join(t.TempDir(), "storage.json"))

	store.AddExercise(models.ExerciseInput{Name: "Bench"})
	exercises, _ := store.GetExercises()

	store.AddStrength(
		exercises[0].Id,
		models.StrengthInput{
			Date: "2024",
			Load: 50,
		},
	)

	exercise, _ := store.GetExercise(exercises[0].Id)
	strength := exercise.StrengthHistory[0]
	strength.Load = 100

	err := store.UpdateStrength(strength)
	if err != nil {
		t.Fatal(err)
	}

	updated, _ := store.GetStrength(strength.Id)
	if updated.Load != 100 {
		t.Fatalf("got %f", updated.Load)
	}
}

func TestDeleteStrength(t *testing.T) {
	store := New(filepath.Join(t.TempDir(), "storage.json"))

	store.AddExercise(models.ExerciseInput{Name: "Bench"})
	exercises, _ := store.GetExercises()

	store.AddStrength(
		exercises[0].Id,
		models.StrengthInput{
			Date: "2024",
		},
	)

	updated, _ := store.GetExercise(exercises[0].Id)
	strengthID := updated.StrengthHistory[0].Id

	err := store.DeleteStrength(strengthID)
	if err != nil {
		t.Fatal(err)
	}

	updated, _ = store.GetExercise(exercises[0].Id)
	if len(updated.StrengthHistory) != 0 {
		t.Fatal("strength not deleted")
	}
}

func TestSortStrengths(t *testing.T) {
	strengths := []models.Strength{
		{StrengthInput: models.StrengthInput{Date: "2026-01-01"}},
		{StrengthInput: models.StrengthInput{Date: "2026-02-01"}},
	}

	sorted := sortStrengths(strengths)
	if sorted[0].Date != "2026-02-01" {
		t.Fatal("not sorted")
	}
}
