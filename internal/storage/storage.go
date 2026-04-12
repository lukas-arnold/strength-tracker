package storage

import (
	"os"

	"github.com/lukas-arnold/strength-tracker/internal/configs"
	"github.com/lukas-arnold/strength-tracker/internal/models"
	"github.com/lukas-arnold/strength-tracker/internal/utils"
)

func saveStorage(exercises []models.Exercise) error {
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
		saveStorage([]models.Exercise{})
	}
}
