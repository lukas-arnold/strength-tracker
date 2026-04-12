package utils

import (
	"encoding/json"
	"strconv"

	"github.com/lukas-arnold/strength-tracker/internal/models"
)

func ConvertExercisesToBytes(exercises []models.Exercise) ([]byte, error) {
	bytes, err := json.Marshal(exercises)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func ConvertBytesToExercises(bytes []byte) ([]models.Exercise, error) {
	var exercises []models.Exercise
	err := json.Unmarshal(bytes, &exercises)
	if err != nil {
		return exercises, err
	}
	return exercises, nil
}

func ConvertId(idStr string) (int64, error) {
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func ConvertLoad(loadStr string) (float64, error) {
	load, err := strconv.ParseFloat(loadStr, 64)
	if err != nil {
		return -1, err
	}
	return load, nil
}
