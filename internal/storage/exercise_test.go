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

	if result[0].LastStrength.Load != -1 {
		t.Errorf("expected default load -1, got %f", result[0].LastStrength.Load)
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
