package storage

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/lukas-arnold/strength-tracker/internal/models"
)

func testStorage(t *testing.T) *Storage {
	t.Helper()

	return New(filepath.Join(t.TempDir(), "storage.json"))
}

func TestStorageCreatesFile(t *testing.T) {
	store := testStorage(t)
	_, _ = store.readStorage()

	if _, err := os.Stat(store.file); os.IsNotExist(err) {
		t.Fatal("expected storage file to be created")
	}
}

func TestStorageCreatesMissingDirectory(t *testing.T) {
	dir := t.TempDir()
	store := New(filepath.Join(dir, "nested", "storage.json"))

	err := store.checkStorage()
	if err != nil {
		t.Fatal(err)
	}

	_, err = os.Stat(filepath.Join(dir, "nested"))
	if err != nil {
		t.Fatal("expected directory to be created")
	}
}

func TestSaveStoragePersistsData(t *testing.T) {
	store := testStorage(t)

	exercises := []models.Exercise{{
		ExerciseInput: models.ExerciseInput{Name: "Bench", MuscleGroup: "Chest"},
	}}

	if err := store.saveStorage(exercises); err != nil {
		t.Fatalf("failed to save: %v", err)
	}

	result, err := store.GetExercises()
	if err != nil {
		t.Fatalf("failed to get exercises: %v", err)
	}

	if len(result) != 1 || result[0].Name != "Bench" {
		t.Errorf("expected 1 exercise 'Bench', got %+v", result)
	}
}

func TestSaveStorageSortsExercises(t *testing.T) {
	store := testStorage(t)
	input := []models.Exercise{
		{ExerciseInput: models.ExerciseInput{Name: "Squat", MuscleGroup: "Legs"}},
		{ExerciseInput: models.ExerciseInput{Name: "Bench", MuscleGroup: "Chest"}},
	}

	if err := store.saveStorage(input); err != nil {
		t.Fatal(err)
	}

	result, _ := store.GetExercises()
	if result[0].Name != "Bench" {
		t.Errorf("expected sorted result starting with 'Bench', got %s", result[0].Name)
	}
}

func TestSaveStorageSortsStrengthHistory(t *testing.T) {
	store := testStorage(t)
	input := []models.Exercise{{
		StrengthHistory: []models.Strength{
			{StrengthInput: models.StrengthInput{Date: "2026-01-01"}},
			{StrengthInput: models.StrengthInput{Date: "2026-02-01"}},
		},
	}}

	if err := store.saveStorage(input); err != nil {
		t.Fatal(err)
	}

	result, _ := store.GetExercises()
	// Note: Check your sortStrengths logic; typical DESC order puts newest first
	if result[0].StrengthHistory[0].Date != "2026-02-01" {
		t.Errorf("expected newest date first, got %s", result[0].StrengthHistory[0].Date)
	}
}

func TestReadStorageCreatesEmptyStorage(t *testing.T) {
	store := testStorage(t)
	data, err := store.readStorage()

	if err != nil || len(data) == 0 {
		t.Fatalf("expected initialized empty storage, got err: %v, data len: %d", err, len(data))
	}
}

func TestGetExercisesEmpty(t *testing.T) {
	store := testStorage(t)
	result, err := store.GetExercises()

	if err != nil {
		t.Fatal(err)
	}
	if len(result) != 0 {
		t.Errorf("expected 0 exercises, got %d", len(result))
	}
}
