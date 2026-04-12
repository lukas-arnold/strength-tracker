package storage

import (
	"slices"

	"github.com/lukas-arnold/strength-tracker/internal/models"
	"github.com/lukas-arnold/strength-tracker/internal/utils"
)

func AddExercise(exercise models.ExerciseInput) error {
	exercises, err := GetExercises()
	if err != nil {
		return err
	}
	newExercise := models.Exercise{Id: utils.Id(), ExerciseInput: exercise, StrengthHistory: []models.Strength{}}
	exercises = append(exercises, newExercise)
	err = saveStorage(exercises)
	if err != nil {
		return err
	}
	return nil
}

func GetExercises() ([]models.Exercise, error) {
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

func GetExercise(id int64) (models.Exercise, error) {
	exercises, err := GetExercises()
	if err != nil {
		return models.Exercise{}, err
	}
	var exercise models.Exercise
	for _, value := range exercises {
		if value.Id == id {
			exercise = value
		}
	}
	return exercise, nil
}

func GetExerciseByStrength(strengthId int64) (models.Exercise, error) {
	exercises, err := GetExercises()
	if err != nil {
		return models.Exercise{}, err
	}
	var exercise models.Exercise
	for _, value := range exercises {
		for _, strength := range value.StrengthHistory {
			if strength.Id == strengthId {
				exercise = value
			}
		}
	}
	return exercise, nil
}

func UpdateExercise(exercise models.Exercise) error {
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

func DeleteExercise(id int64) error {
	exercises, err := GetExercises()
	if err != nil {
		return err
	}
	var index int
	for i := range exercises {
		if exercises[i].Id == id {
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
