package utils

import (
	"encoding/json"
	"strconv"

	"github.com/lukas-arnold/strength-tracker/internal/models"
)

func ConvertExercisesToBytes(exercises []models.Exercise) ([]byte, error) {
	return json.Marshal(exercises)
}

func ConvertBytesToExercises(bytes []byte) ([]models.Exercise, error) {
	var exercises []models.Exercise
	err := json.Unmarshal(bytes, &exercises)
	return exercises, err
}

func ConvertToInt(intStr string) (int64, error) {
	return strconv.ParseInt(intStr, 10, 64)
}

func ConvertLoad(loadStr string) (float64, error) {
	return strconv.ParseFloat(loadStr, 64)
}
