package utils

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/lukas-arnold/strength-tracker/internal/models"
)

func ConvertExerciseToBytes(exercise models.ExerciseFull) ([]byte, error) {
	bytes, err := json.Marshal(exercise)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func ConvertExercisesToBytes(exercises []models.ExerciseFull) ([]byte, error) {
	bytes, err := json.Marshal(exercises)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func ConvertBytesToExercise(bytes []byte) (models.ExerciseFull, error) {
	var exercise models.ExerciseFull
	err := json.Unmarshal(bytes, &exercise)
	if err != nil {
		return exercise, err
	}
	return exercise, nil
}

func ConvertBytesToAddExercise(bytes []byte) (models.Exercise, error) {
	var exercise models.Exercise
	err := json.Unmarshal(bytes, &exercise)
	if err != nil {
		return exercise, err
	}
	return exercise, nil
}

func ConvertBytesToExercises(bytes []byte) ([]models.ExerciseFull, error) {
	var exercises []models.ExerciseFull
	err := json.Unmarshal(bytes, &exercises)
	if err != nil {
		return exercises, err
	}
	return exercises, nil
}

func ConvertId(id string) (int64, error) {
	id_int, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return -1, err
	}
	return id_int, nil
}

func ConvertTime(time_str string) (time.Time, error) {
	time_conv, err := time.Parse("2006-01-02T15:04", time_str)
	if err != nil {
		return time_conv, err
	}
	return time_conv, nil
}
