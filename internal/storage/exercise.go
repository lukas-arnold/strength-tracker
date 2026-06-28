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

	exercises = append(exercises, models.Exercise{
		Id:              utils.Id(),
		ExerciseInput:   input,
		StrengthHistory: []models.Strength{},
	})

	return s.saveStorage(exercises)
}

func (s *Storage) GetExercises() ([]models.Exercise, error) {
	data, err := s.readStorage()
	if err != nil {
		return nil, err
	}
	return utils.ConvertBytesToExercises(data)
}

func (s *Storage) GetExercisesWithLastStrength() ([]models.ExerciseWithLastStrength, error) {
	exercises, err := s.GetExercises()
	if err != nil {
		return nil, err
	}

	result := make([]models.ExerciseWithLastStrength, len(exercises))
	for i, ex := range exercises {
		last, err := s.GetLastStrength(ex.Id)
		if err != nil {
			return nil, err
		}

		if last.Load == 0 {
			last.Load = -1
			last.Date = "1970-01-01"
		}

		result[i] = models.ExerciseWithLastStrength{
			Exercise:     ex,
			LastStrength: last,
		}
	}
	return result, nil
}

func (s *Storage) GetExercise(id int64) (models.Exercise, error) {
	exercises, err := s.GetExercises()
	if err != nil {
		return models.Exercise{}, err
	}

	for _, ex := range exercises {
		if ex.Id == id {
			return ex, nil
		}
	}
	return models.Exercise{}, nil
}

func (s *Storage) GetExerciseForHistoryChart(id int64) (models.ExerciseForHistoryChart, error) {
	ex, err := s.GetExercise(id)
	if err != nil {
		return models.ExerciseForHistoryChart{}, err
	}

	n := len(ex.StrengthHistory)
	chart := models.ExerciseForHistoryChart{
		Exercise:    ex,
		Dates:       make([]string, n),
		Loads:       make([]float64, n),
		Repetitions: make([]int64, n),
	}

	// Reverse StrengthHistory while filling chart slices
	for i, strength := range ex.StrengthHistory {
		idx := n - 1 - i
		chart.Dates[idx] = strength.Date
		chart.Loads[idx] = strength.Load
		chart.Repetitions[idx] = strength.Repetitions
	}

	return chart, nil
}

func (s *Storage) UpdateExercise(ex models.Exercise) error {
	exercises, err := s.GetExercises()
	if err != nil {
		return err
	}

	for i := range exercises {
		if exercises[i].Id == ex.Id {
			exercises[i] = ex
			return s.saveStorage(exercises)
		}
	}
	return nil
}

func (s *Storage) DeleteExercise(id int64) error {
	exercises, err := s.GetExercises()
	if err != nil {
		return err
	}

	idx := slices.IndexFunc(exercises, func(ex models.Exercise) bool { return ex.Id == id })
	if idx != -1 {
		exercises = slices.Delete(exercises, idx, idx+1)
		return s.saveStorage(exercises)
	}
	return nil
}

func sortExercises(exercises []models.Exercise) []models.Exercise {
	sort.Slice(exercises, func(i, j int) bool {
		a, b := exercises[i], exercises[j]
		if a.MuscleGroup != b.MuscleGroup {
			return a.MuscleGroup < b.MuscleGroup
		}
		if a.Name != b.Name {
			return a.Name < b.Name
		}
		return a.Machine < b.Machine
	})
	return exercises
}
