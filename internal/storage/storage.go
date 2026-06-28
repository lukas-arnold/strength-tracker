package storage

import (
	"os"
	"path/filepath"

	"github.com/lukas-arnold/strength-tracker/internal/models"
	"github.com/lukas-arnold/strength-tracker/internal/utils"
)

type Storage struct {
	file string
}

func New(file string) *Storage {
	return &Storage{
		file: file,
	}
}

func (s *Storage) saveStorage(exercises []models.Exercise) error {
	exercises = sortStorage(exercises)

	bytes, err := utils.ConvertExercisesToBytes(exercises)
	if err != nil {
		return err
	}

	return os.WriteFile(s.file, bytes, 0666)
}

func (s *Storage) readStorage() ([]byte, error) {
	err := s.checkStorage()
	if err != nil {
		return nil, err
	}

	return os.ReadFile(s.file)
}

func (s *Storage) checkStorage() error {
	_, err := os.ReadFile(s.file)

	if err == nil {
		return nil
	}

	err = os.MkdirAll(filepath.Dir(s.file), 0755)
	if err != nil {
		return err
	}

	return s.saveStorage([]models.Exercise{})
}

func sortStorage(exercises []models.Exercise) []models.Exercise {
	exercises = sortExercises(exercises)

	for i := range exercises {
		exercises[i].StrengthHistory =
			sortStrengths(exercises[i].StrengthHistory)
	}

	return exercises
}
