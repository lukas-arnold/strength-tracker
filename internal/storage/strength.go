package storage

import (
	"slices"
	"sort"

	"github.com/lukas-arnold/strength-tracker/internal/models"
	"github.com/lukas-arnold/strength-tracker/internal/utils"
)

func (s *Storage) AddStrength(exerciseId int64, input models.StrengthInput) error {
	exercise, err := s.GetExercise(exerciseId)
	if err != nil {
		return err
	}

	strength := models.Strength{
		Id:            utils.Id(),
		StrengthInput: input,
	}

	exercise.StrengthHistory = append(exercise.StrengthHistory, strength)
	return s.UpdateExercise(exercise)
}

func (s *Storage) GetStrength(id int64) (models.Strength, error) {
	exercises, err := s.GetExercises()
	if err != nil {
		return models.Strength{}, err
	}

	for _, exercise := range exercises {
		for _, strength := range exercise.StrengthHistory {
			if strength.Id == id {
				return strength, nil
			}
		}
	}

	return models.Strength{}, nil
}

func (s *Storage) GetLastStrength(exerciseId int64) (models.Strength, error) {
	exercise, err := s.GetExercise(exerciseId)
	if err != nil {
		return models.Strength{}, err
	}

	if len(exercise.StrengthHistory) == 0 {
		return models.Strength{}, nil
	}

	return exercise.StrengthHistory[0], nil
}

func (s *Storage) UpdateStrength(strength models.Strength) error {
	exercises, err := s.GetExercises()
	if err != nil {
		return err
	}

	for i := range exercises {
		for j, st := range exercises[i].StrengthHistory {
			if st.Id == strength.Id {
				exercises[i].StrengthHistory[j] = strength
				return s.saveStorage(exercises)
			}
		}
	}

	return nil
}

func (s *Storage) DeleteStrength(id int64) error {
	exercises, err := s.GetExercises()
	if err != nil {
		return err
	}

	for i := range exercises {
		for j, st := range exercises[i].StrengthHistory {
			if st.Id == id {
				exercises[i].StrengthHistory = slices.Delete(exercises[i].StrengthHistory, j, j+1)
				return s.saveStorage(exercises)
			}
		}
	}

	return nil
}

func sortStrengths(strengths []models.Strength) []models.Strength {
	sort.Slice(strengths, func(i, j int) bool {
		return strengths[j].Date < strengths[i].Date
	})
	return strengths
}
