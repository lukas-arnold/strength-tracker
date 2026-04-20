package storage

import (
	"slices"
	"sort"

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

func GetExercisesWithLastStrength() ([]models.ExerciseWithLastStrength, error) {
	bytes, err := readStorage()
	if err != nil {
		return nil, err
	}
	exercises, err := utils.ConvertBytesToExercises(bytes)
	if err != nil {
		return nil, err
	}
	var exercisesWithLastStrength []models.ExerciseWithLastStrength
	for _, exercise := range exercises {
		lastStrength, err := GetLastStrength(exercise.Id)
		if err != nil {
			return nil, err
		}
		exercisesWithLastStrength = append(exercisesWithLastStrength, models.ExerciseWithLastStrength{Exercise: exercise, LastStrength: lastStrength})
	}
	return exercisesWithLastStrength, nil
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

func GetExerciseForHistoryChart(id int64) (models.ExerciseForHistoryChart, error) {
	exercise, err := GetExercise(id)
	if err != nil {
		return models.ExerciseForHistoryChart{}, err
	}
	var dates []string
	var loads []float64
	for _, value := range exercise.StrengthHistory {
		dates = append(dates, value.Date)
		loads = append(loads, value.Load)
	}
	exerciseForHistoryChart := models.ExerciseForHistoryChart{Exercise: models.Exercise{Id: exercise.Id, ExerciseInput: models.ExerciseInput{Name: exercise.Name, MuscleGroup: exercise.MuscleGroup}, StrengthHistory: exercise.StrengthHistory}, Dates: dates, Loads: loads}
	return exerciseForHistoryChart, nil
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

func sortExercises(exercises []models.Exercise) []models.Exercise {
	sort.Slice(exercises, func(i, j int) bool {
		if exercises[i].MuscleGroup != exercises[j].MuscleGroup {
			return exercises[i].MuscleGroup < exercises[j].MuscleGroup
		}
		if exercises[i].Name != exercises[j].Name {
			return exercises[i].Name < exercises[j].Name
		}
		return false
	})
	return exercises
}
