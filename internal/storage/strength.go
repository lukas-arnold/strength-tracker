package storage

import (
	"slices"
	"sort"

	"github.com/lukas-arnold/strength-tracker/internal/models"
	"github.com/lukas-arnold/strength-tracker/internal/utils"
)

func AddStrength(exerciseId int64, strengthInput models.StrengthInput) error {
	exercise, err := GetExercise(exerciseId)
	if err != nil {
		return err
	}
	strengthId := utils.Id()
	strength := models.Strength{Id: strengthId, StrengthInput: strengthInput}
	strengtHistory := append(exercise.StrengthHistory, strength)
	exercise.StrengthHistory = strengtHistory
	err = UpdateExercise(exercise)
	if err != nil {
		return err
	}
	return nil
}

func GetStrength(id int64) (models.Strength, error) {
	exercises, err := GetExercises()
	if err != nil {
		return models.Strength{}, err
	}
	var strength models.Strength
	for _, exercise := range exercises {
		for _, value := range exercise.StrengthHistory {
			if value.Id == id {
				strength = value
			}
		}
	}
	return strength, nil
}

func GetLastStrength(exerciseId int64) (models.Strength, error) {
	exercises, err := GetExercises()
	if err != nil {
		return models.Strength{}, err
	}
	var exercise models.Exercise
	for _, value := range exercises {
		if value.Id == exerciseId {
			exercise = value
		}
	}
	strengthHistoryLength := len(exercise.StrengthHistory)
	var lastStrength models.Strength
	if strengthHistoryLength >= 1 {
		lastStrength = exercise.StrengthHistory[len(exercise.StrengthHistory)-1]
		return lastStrength, nil
	}
	return models.Strength{}, nil
}

func UpdateStrength(strength models.Strength) error {
	exercises, err := GetExercises()
	if err != nil {
		return err
	}
	for i := range exercises {
		for j := range exercises[i].StrengthHistory {
			if exercises[i].StrengthHistory[j].Id == strength.Id {
				exercises[i].StrengthHistory[j].Date = strength.Date
				exercises[i].StrengthHistory[j].Load = strength.Load
				exercises[i].StrengthHistory[j].Repetitions = strength.Repetitions
			}
		}
	}
	err = saveStorage(exercises)
	if err != nil {
		return err
	}
	return nil
}

func DeleteStrength(id int64) error {
	exercises, err := GetExercises()
	if err != nil {
		return err
	}
	for i := range exercises {
		for j := range exercises[i].StrengthHistory {
			if exercises[i].StrengthHistory[j].Id == id {
				exercises[i].StrengthHistory = slices.Delete(exercises[i].StrengthHistory, j, j+1)
				return saveStorage(exercises)
			}
		}
	}
	return nil
}

func sortStrengths(strengths []models.Strength) []models.Strength {
	sort.Slice(strengths, func(i, j int) bool {
		return strengths[i].Date < strengths[j].Date
	})
	return strengths
}
