package storage

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/lukas-arnold/strength-tracker/internal/models"
)

func newTestStorage(t *testing.T) *Storage {
	t.Helper()

	return New(
		filepath.Join(
			t.TempDir(),
			"storage.json",
		),
	)
}

func TestStorageCreatesFile(t *testing.T) {

	store := newTestStorage(t)

	err := store.checkStorage()

	if err != nil {
		t.Fatal(err)
	}

	data, err := store.readStorage()

	if err != nil {
		t.Fatal(err)
	}

	if len(data) == 0 {
		t.Fatal("expected storage file")
	}
}

func TestStorageCreatesMissingDirectory(t *testing.T) {

	dir := t.TempDir()

	store := New(
		filepath.Join(
			dir,
			"nested",
			"storage.json",
		),
	)

	err := store.checkStorage()

	if err != nil {
		t.Fatal(err)
	}

	_, err = os.Stat(
		filepath.Join(
			dir,
			"nested",
		),
	)

	if err != nil {
		t.Fatal("expected directory to be created")
	}
}

func TestSaveStoragePersistsData(t *testing.T) {

	store := newTestStorage(t)

	err := store.saveStorage([]models.Exercise{
		{
			ExerciseInput: models.ExerciseInput{
				Name:        "Bench",
				MuscleGroup: "Chest",
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	result, err := store.GetExercises()

	if err != nil {
		t.Fatal(err)
	}

	if len(result) != 1 {
		t.Fatalf(
			"expected 1 exercise got %d",
			len(result),
		)
	}

	if result[0].Name != "Bench" {
		t.Fatalf(
			"wrong exercise %s",
			result[0].Name,
		)
	}
}

func TestSaveStorageSortsExercises(t *testing.T) {

	store := newTestStorage(t)

	err := store.saveStorage([]models.Exercise{
		{
			ExerciseInput: models.ExerciseInput{
				Name:        "Squat",
				MuscleGroup: "Legs",
			},
		},
		{
			ExerciseInput: models.ExerciseInput{
				Name:        "Bench",
				MuscleGroup: "Chest",
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	result, err := store.GetExercises()

	if err != nil {
		t.Fatal(err)
	}

	if result[0].Name != "Bench" {
		t.Fatalf(
			"expected sorted exercise got %s",
			result[0].Name,
		)
	}
}

func TestSaveStorageSortsStrengthHistory(t *testing.T) {

	store := newTestStorage(t)

	err := store.saveStorage([]models.Exercise{
		{
			StrengthHistory: []models.Strength{
				{
					StrengthInput: models.StrengthInput{
						Date: "2026-01-01",
					},
				},
				{
					StrengthInput: models.StrengthInput{
						Date: "2026-02-01",
					},
				},
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	result, err := store.GetExercises()

	if err != nil {
		t.Fatal(err)
	}

	history := result[0].StrengthHistory

	if history[0].Date != "2026-02-01" {
		t.Fatalf(
			"expected sorted strengths got %s",
			history[0].Date,
		)
	}
}

func TestReadStorageCreatesEmptyStorage(t *testing.T) {

	store := newTestStorage(t)

	data, err := store.readStorage()

	if err != nil {
		t.Fatal(err)
	}

	if len(data) == 0 {
		t.Fatal(
			"expected initialized storage",
		)
	}
}

func TestGetExercisesEmpty(t *testing.T) {

	store := newTestStorage(t)

	result, err := store.GetExercises()

	if err != nil {
		t.Fatal(err)
	}

	if len(result) != 0 {
		t.Fatal(
			"expected no exercises",
		)
	}
}
