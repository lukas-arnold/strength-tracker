package utils

import (
	"testing"

	"github.com/lukas-arnold/strength-tracker/internal/models"
)

func TestConvertExercisesToBytes(t *testing.T) {
	input := []models.Exercise{
		{
			Id: 1,
			ExerciseInput: models.ExerciseInput{
				Name:        "Bench Press",
				MuscleGroup: "Chest",
				Machine:     "Barbell",
			},
			StrengthHistory: []models.Strength{
				{
					Id: 10,
					StrengthInput: models.StrengthInput{
						Date:        "2024-01-01",
						Load:        100,
						Repetitions: 5,
					},
				},
			},
		},
	}

	bytes, err := ConvertExercisesToBytes(input)

	if err != nil {
		t.Fatal(err)
	}

	if len(bytes) == 0 {
		t.Fatal("expected json bytes")
	}
}

func TestConvertBytesToExercises(t *testing.T) {
	json := []byte(`
	[
		{
			"Id":1,
			"Name":"Bench Press",
			"MuscleGroup":"Chest",
			"Machine":"Barbell",
			"StrengthHistory":[
				{
					"Id":10,
					"Date":"2024-01-01",
					"Load":100,
					"Repetitions":5
				}
			]
		}
	]
	`)

	exercises, err := ConvertBytesToExercises(json)

	if err != nil {
		t.Fatal(err)
	}

	if len(exercises) != 1 {
		t.Fatalf(
			"got %d exercises",
			len(exercises),
		)
	}

	exercise := exercises[0]

	if exercise.Name != "Bench Press" {
		t.Fatalf(
			"got %s",
			exercise.Name,
		)
	}

	if exercise.StrengthHistory[0].Load != 100 {
		t.Fatalf(
			"got %f",
			exercise.StrengthHistory[0].Load,
		)
	}
}

func TestConvertBytesToExercisesInvalid(t *testing.T) {
	_, err := ConvertBytesToExercises(
		[]byte("not json"),
	)

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestConvertToInt(t *testing.T) {
	got, err := ConvertToInt("123")

	if err != nil {
		t.Fatal(err)
	}

	if got != 123 {
		t.Fatalf(
			"got %d want %d",
			got,
			123,
		)
	}
}

func TestConvertToIntInvalid(t *testing.T) {
	_, err := ConvertToInt("abc")

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestConvertLoad(t *testing.T) {
	got, err := ConvertLoad("42.5")

	if err != nil {
		t.Fatal(err)
	}

	if got != 42.5 {
		t.Fatalf(
			"got %f want %f",
			got,
			42.5,
		)
	}
}

func TestConvertLoadInvalid(t *testing.T) {
	_, err := ConvertLoad("abc")

	if err == nil {
		t.Fatal("expected error")
	}
}
