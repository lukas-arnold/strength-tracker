package storage

import (
	"path/filepath"
	"testing"

	"github.com/lukas-arnold/strength-tracker/internal/models"
)

func TestAddExercise(t *testing.T) {

	store := New(
		filepath.Join(
			t.TempDir(),
			"storage.json",
		),
	)

	err := store.AddExercise(
		models.ExerciseInput{
			Name: "Bench",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	exercises, _ := store.GetExercises()

	if len(exercises) != 1 {
		t.Fatal("exercise not added")
	}
}

func TestGetExercise(t *testing.T) {

	store := New(
		filepath.Join(
			t.TempDir(),
			"storage.json",
		),
	)

	store.AddExercise(
		models.ExerciseInput{
			Name: "Bench",
		},
	)

	exercises, _ := store.GetExercises()

	exercise, err := store.GetExercise(
		exercises[0].Id,
	)

	if err != nil {
		t.Fatal(err)
	}

	if exercise.Name != "Bench" {
		t.Fatalf(
			"got %s",
			exercise.Name,
		)
	}
}

func TestGetExerciseNotFound(t *testing.T) {

	store := New(
		filepath.Join(
			t.TempDir(),
			"storage.json",
		),
	)

	exercise, err := store.GetExercise(999)

	if err != nil {
		t.Fatal(err)
	}

	if exercise.Id != 0 {
		t.Fatal("expected empty exercise")
	}
}

func TestGetExercisesWithLastStrength(t *testing.T) {

	store := New(
		filepath.Join(
			t.TempDir(),
			"storage.json",
		),
	)

	store.AddExercise(
		models.ExerciseInput{
			Name: "Bench",
		},
	)

	exercises, _ := store.GetExercises()

	result, err := store.GetExercisesWithLastStrength()

	if err != nil {
		t.Fatal(err)
	}

	if len(result) != 1 {
		t.Fatal("missing exercise")
	}

	if result[0].LastStrength.Load != -1 {
		t.Fatalf(
			"got %f",
			result[0].LastStrength.Load,
		)
	}

	_ = exercises
}

func TestGetExerciseForHistoryChart(t *testing.T) {

	store := New(
		filepath.Join(
			t.TempDir(),
			"storage.json",
		),
	)

	store.AddExercise(
		models.ExerciseInput{
			Name: "Bench",
		},
	)

	exercises, _ := store.GetExercises()

	store.AddStrength(
		exercises[0].Id,
		models.StrengthInput{
			Date:        "2024-01-01",
			Load:        100,
			Repetitions: 5,
		},
	)

	result, err := store.GetExerciseForHistoryChart(
		exercises[0].Id,
	)

	if err != nil {
		t.Fatal(err)
	}

	if len(result.Dates) != 1 {
		t.Fatal("chart data missing")
	}

	if result.Loads[0] != 100 {
		t.Fatalf(
			"got %f",
			result.Loads[0],
		)
	}
}

func TestUpdateExercise(t *testing.T) {

	store := New(
		filepath.Join(
			t.TempDir(),
			"storage.json",
		),
	)

	store.AddExercise(
		models.ExerciseInput{
			Name: "Bench",
		},
	)

	exercises, _ := store.GetExercises()

	exercise := exercises[0]

	exercise.Name = "Squat"

	err := store.UpdateExercise(exercise)

	if err != nil {
		t.Fatal(err)
	}

	updated, _ := store.GetExercise(
		exercise.Id,
	)

	if updated.Name != "Squat" {
		t.Fatalf(
			"got %s",
			updated.Name,
		)
	}
}

func TestDeleteExercise(t *testing.T) {

	store := New(
		filepath.Join(
			t.TempDir(),
			"storage.json",
		),
	)

	store.AddExercise(
		models.ExerciseInput{
			Name: "Bench",
		},
	)

	exercises, _ := store.GetExercises()

	err := store.DeleteExercise(
		exercises[0].Id,
	)

	if err != nil {
		t.Fatal(err)
	}

	remaining, _ := store.GetExercises()

	if len(remaining) != 0 {
		t.Fatal("exercise not deleted")
	}
}

func TestSortExercises(t *testing.T) {

	exercises := []models.Exercise{
		{
			ExerciseInput: models.ExerciseInput{
				Name:        "Squat",
				MuscleGroup: "Legs",
				Machine:     "Barbell",
			},
		},
		{
			ExerciseInput: models.ExerciseInput{
				Name:        "Bench",
				MuscleGroup: "Chest",
				Machine:     "Barbell",
			},
		},
	}

	sorted := sortExercises(exercises)

	if sorted[0].Name != "Bench" {
		t.Fatal("not sorted")
	}
}
