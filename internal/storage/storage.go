package storage

import (
	"os"
	"slices"

	"github.com/lukas-arnold/strength-tracker/internal/configs"
	"github.com/lukas-arnold/strength-tracker/internal/models"
	"github.com/lukas-arnold/strength-tracker/internal/utils"
)

func saveStorage(exercises []models.ExerciseFull) error {
	bytes, err := utils.ConvertExercisesToBytes(exercises)
	if err != nil {
		return err
	}
	err = os.WriteFile(configs.STORAGE_FILE, []byte(bytes), 0666)
	if err != nil {
		return err
	}
	return nil
}

func readStorage() ([]byte, error) {
	checkStorage()
	bytes, err := os.ReadFile(configs.STORAGE_FILE)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func checkStorage() {
	_, err := os.ReadFile(configs.STORAGE_FILE)
	if err != nil {
		saveStorage([]models.ExerciseFull{})
	}
}

func AddExercise(exercise models.Exercise) error {
	exercises, err := GetExercises()
	if err != nil {
		return err
	}
	new_exercise := models.ExerciseAdd{Id: utils.Id(), Exercise: exercise}
	exercises = append(exercises, models.ExerciseFull{ExerciseAdd: new_exercise, StrengthHistory: []models.Strength{}})
	err = saveStorage(exercises)
	if err != nil {
		return err
	}
	return nil
}

func GetExercise(id int64) (models.ExerciseFull, error) {
	exercises, err := GetExercises()
	if err != nil {
		return models.ExerciseFull{}, err
	}
	var exercise models.ExerciseFull
	for _, value := range exercises {
		if value.Id == id {
			exercise = value
		}
	}
	return exercise, nil
}

func GetExercises() ([]models.ExerciseFull, error) {
	bytes, err := readStorage()
	if err != nil {
		return nil, err
	}
	exercises, err := utils.ConvertBytesToExercises(bytes)
	if err != nil {
		return nil, err
	}
	return exercises, nil
}

func GetExercisesBytes() ([]byte, error) {
	bytes, err := readStorage()
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func UpdateExercise(exercise models.ExerciseFull) error {
	exercises, err := GetExercises()
	if err != nil {
		return err
	}
	for i := range exercises {
		if exercises[i].Id == exercise.Id {
			exercises[i].Name = exercise.Name
			exercises[i].MuscleGroup = exercise.MuscleGroup
			exercises[i].StrengthHistory = exercise.StrengthHistory
		}
	}
	err = saveStorage(exercises)
	if err != nil {
		return err
	}
	return nil
}

func DeleteExercise(name string) error {
	exercises, err := GetExercises()
	if err != nil {
		return err
	}
	var index int
	for i := range exercises {
		if exercises[i].Name == name {
			index = i
		}
	}
	exercises = slices.Delete(exercises, index, index+1)
	err = saveStorage(exercises)
	if err != nil {
		return err
	}
	return nil
}
