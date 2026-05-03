package storage

import (
	"os"

	"github.com/lukas-arnold/strength-tracker/internal/configs"
	"github.com/lukas-arnold/strength-tracker/internal/models"
	"github.com/lukas-arnold/strength-tracker/internal/utils"
)

func saveStorage(exercises []models.Exercise) error {
	exercises = sortStorage(exercises)
	bytes, err := utils.ConvertExercisesToBytes(exercises)
	if err != nil {
		return err
	}
	err = os.WriteFile(configs.GetStorageFile(), []byte(bytes), 0666)
	if err != nil {
		return err
	}
	return nil
}

func readStorage() ([]byte, error) {
	checkStorage()
	bytes, err := os.ReadFile(configs.GetStorageFile())
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func checkStorage() {
	_, err := os.ReadFile(configs.GetStorageFile())
	if err != nil {
		saveStorage([]models.Exercise{})
	}
}

func sortStorage(exercises []models.Exercise) []models.Exercise {
	exercises = sortExercises(exercises)
	for i := range exercises {
		exercises[i].StrengthHistory = sortStrengths(exercises[i].StrengthHistory)
	}
	return exercises
}
