package storage

import (
	"testing"

	"github.com/lukas-arnold/strength-tracker/internal/models"
)

func TestAddExercise(t *testing.T) {
	store := testStorage(t)

	err := store.AddExercise(models.ExerciseInput{Name: "Bench"})
	if err != nil {
		t.Fatalf("could not add exercise: %v", err)
	}

	exercises, _ := store.GetExercises()
	if len(exercises) != 1 {
		t.Errorf("expected 1 exercise, got %d", len(exercises))
	}
}

func TestGetExercise(t *testing.T) {
	store := testStorage(t)
	store.AddExercise(models.ExerciseInput{Name: "Bench"})

	exercises, _ := store.GetExercises()
	exercise, err := store.GetExercise(exercises[0].Id)
	if err != nil {
		t.Fatalf("error getting exercise: %v", err)
	}

	if exercise.Name != "Bench" {
		t.Errorf("expected Bench, got %s", exercise.Name)
	}
}

func TestGetExerciseNotFound(t *testing.T) {
	store := testStorage(t)

	exercise, err := store.GetExercise(999)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if exercise.Id != 0 {
		t.Errorf("expected empty exercise, got id %d", exercise.Id)
	}
}

func TestGetExercisesWithLastStrength(t *testing.T) {
	store := testStorage(t)

	store.AddExercise(models.ExerciseInput{Name: "Bench"})

	result, err := store.GetExercisesWithLastStrength()
	if err != nil {
		t.Fatalf("error getting exercises: %v", err)
	}

	if len(result) != 1 {
		t.Fatalf("expected 1 result, got %d", len(result))
	}

	if result[0].LastStrength.Load != 0 {
		t.Errorf("expected default load 0, got %f", result[0].LastStrength.Load)
	}

	if result[0].LastStrength.Date != "" {
		t.Errorf("expected empty date, got %s", result[0].LastStrength.Date)
	}
}

func TestGetExercisesWithLastStrengthExisting(t *testing.T) {
	store := testStorage(t)

	store.AddExercise(models.ExerciseInput{Name: "Bench"})

	exercises, _ := store.GetExercises()

	store.AddStrength(exercises[0].Id, models.StrengthInput{
		Date:        "2026-01-01",
		Load:        100,
		Repetitions: 5,
	})

	result, err := store.GetExercisesWithLastStrength()
	if err != nil {
		t.Fatalf("error getting exercises: %v", err)
	}

	if result[0].LastStrength.Load != 100 {
		t.Errorf("expected load 100, got %f", result[0].LastStrength.Load)
	}

	if result[0].LastStrength.Date != "2026-01-01" {
		t.Errorf("expected date 2026-01-01, got %s", result[0].LastStrength.Date)
	}
}

func TestGetExerciseByStrength(t *testing.T) {
	store := testStorage(t)

	err := store.AddExercise(models.ExerciseInput{Name: "Bench"})
	if err != nil {
		t.Fatalf("error adding exercise: %v", err)
	}

	exercises, err := store.GetExercises()
	if err != nil {
		t.Fatalf("error getting exercises: %v", err)
	}

	exercise := exercises[0]

	err = store.AddStrength(exercise.Id, models.StrengthInput{
		Date:        "2026-06-28",
		Load:        100,
		Repetitions: 5,
	})
	if err != nil {
		t.Fatalf("error adding strength: %v", err)
	}

	exercises, err = store.GetExercises()
	if err != nil {
		t.Fatalf("error getting exercises: %v", err)
	}

	strength := exercises[0].StrengthHistory[0]

	result, err := store.GetExerciseByStrength(strength)
	if err != nil {
		t.Fatalf("error getting exercise by strength: %v", err)
	}

	if result.Id != exercise.Id {
		t.Errorf("expected exercise id %d, got %d", exercise.Id, result.Id)
	}

	if result.Name != "Bench" {
		t.Errorf("expected Bench, got %s", result.Name)
	}
}

func TestGetExerciseForHistoryChart(t *testing.T) {
	store := testStorage(t)
	store.AddExercise(models.ExerciseInput{Name: "Bench"})

	exercises, _ := store.GetExercises()
	store.AddStrength(exercises[0].Id, models.StrengthInput{
		Date:        "2024-01-01",
		Load:        100,
		Repetitions: 5,
	})

	result, err := store.GetExerciseForHistoryChart(exercises[0].Id)
	if err != nil {
		t.Fatalf("error getting chart data: %v", err)
	}

	if len(result.Dates) != 1 || result.Loads[0] != 100 {
		t.Errorf("chart data missing or incorrect; got loads: %v", result.Loads)
	}
}

func TestUpdateExercise(t *testing.T) {
	store := testStorage(t)
	store.AddExercise(models.ExerciseInput{Name: "Bench"})

	exercises, _ := store.GetExercises()
	exercise := exercises[0]
	exercise.Name = "Squat"

	if err := store.UpdateExercise(exercise); err != nil {
		t.Fatalf("error updating exercise: %v", err)
	}

	updated, _ := store.GetExercise(exercise.Id)
	if updated.Name != "Squat" {
		t.Errorf("expected Squat, got %s", updated.Name)
	}
}

func TestDeleteExercise(t *testing.T) {
	store := testStorage(t)
	store.AddExercise(models.ExerciseInput{Name: "Bench"})

	exercises, _ := store.GetExercises()
	if err := store.DeleteExercise(exercises[0].Id); err != nil {
		t.Fatalf("error deleting exercise: %v", err)
	}

	remaining, _ := store.GetExercises()
	if len(remaining) != 0 {
		t.Errorf("expected 0 exercises, got %d", len(remaining))
	}
}

func TestSortExercises(t *testing.T) {
	exercises := []models.Exercise{
		{ExerciseInput: models.ExerciseInput{Name: "Squat", MuscleGroup: "Legs", Machine: "Barbell"}},
		{ExerciseInput: models.ExerciseInput{Name: "Bench", MuscleGroup: "Chest", Machine: "Barbell"}},
	}

	sorted := sortExercises(exercises)
	if sorted[0].Name != "Bench" {
		t.Errorf("expected Bench first, got %s", sorted[0].Name)
	}
}
