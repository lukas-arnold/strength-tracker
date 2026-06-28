package storage

import (
	"slices"
	"sort"

	"github.com/lukas-arnold/strength-tracker/internal/models"
	"github.com/lukas-arnold/strength-tracker/internal/utils"
)

func (s *Storage) AddExercise(input models.ExerciseInput) error {
	exercises, err := s.GetExercises()
	if err != nil {
		return err
	}

	exercise := models.Exercise{
		Id:              utils.Id(),
		ExerciseInput:   input,
		StrengthHistory: []models.Strength{},
	}

	exercises = append(exercises, exercise)

	return s.saveStorage(exercises)
}

func (s *Storage) GetExercises() ([]models.Exercise, error) {
	bytes, err := s.readStorage()
	if err != nil {
		return nil, err
	}

	return utils.ConvertBytesToExercises(bytes)
}

func (s *Storage) GetExercisesWithLastStrength() ([]models.ExerciseWithLastStrength, error) {
	exercises, err := s.GetExercises()
	if err != nil {
		return nil, err
	}

	result := make([]models.ExerciseWithLastStrength, 0)

	for _, exercise := range exercises {
		last, err := s.GetLastStrength(exercise.Id)

		if err != nil {
			return nil, err
		}

		result = append(
			result,
			models.ExerciseWithLastStrength{
				Exercise:     exercise,
				LastStrength: last,
			},
		)
	}

	for i := range result {
		if result[i].LastStrength.Load == 0 {
			result[i].LastStrength.Load = -1
			result[i].LastStrength.Date = "1970-01-01"
		}
	}

	return result, nil
}

func (s *Storage) GetExercise(id int64) (models.Exercise, error) {
	exercises, err := s.GetExercises()

	if err != nil {
		return models.Exercise{}, err
	}

	for _, exercise := range exercises {
		if exercise.Id == id {
			return exercise, nil
		}
	}

	return models.Exercise{}, nil
}

func (s *Storage) GetExerciseForHistoryChart(id int64) (models.ExerciseForHistoryChart, error) {
	exercise, err := s.GetExercise(id)

	if err != nil {
		return models.ExerciseForHistoryChart{}, err
	}

	dates := []string{}
	loads := []float64{}
	repetitions := []int64{}

	// reverse StrengthHistory
	for i := len(exercise.StrengthHistory) - 1; i >= 0; i-- {
		strength := exercise.StrengthHistory[i]

		dates = append(dates, strength.Date)
		loads = append(loads, strength.Load)
		repetitions = append(repetitions, strength.Repetitions)
	}

	return models.ExerciseForHistoryChart{
		Exercise:    exercise,
		Dates:       dates,
		Loads:       loads,
		Repetitions: repetitions,
	}, nil
}

func (s *Storage) UpdateExercise(exercise models.Exercise) error {
	exercises, err := s.GetExercises()

	if err != nil {
		return err
	}

	for i := range exercises {
		if exercises[i].Id == exercise.Id {
			exercises[i] = exercise
			break
		}
	}

	return s.saveStorage(exercises)
}

func (s *Storage) DeleteExercise(id int64) error {
	exercises, err := s.GetExercises()

	if err != nil {
		return err
	}

	for i := range exercises {
		if exercises[i].Id == id {
			exercises = slices.Delete(exercises, i, i+1)
			return s.saveStorage(exercises)
		}
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

		return exercises[i].Machine < exercises[j].Machine
	})

	return exercises
}
