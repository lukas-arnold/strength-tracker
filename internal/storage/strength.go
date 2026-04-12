package storage

import (
	"slices"

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
	var index int
	for i := range exercises {
		for j := range exercises[i].StrengthHistory {
			if exercises[i].StrengthHistory[j].Id == id {
				index = j
			}
		}
		exercises[i].StrengthHistory = slices.Delete(exercises[i].StrengthHistory, index, index+1)
	}
	err = saveStorage(exercises)
	if err != nil {
		return err
	}
	return nil
}
