package models

type ExerciseInput struct {
	Name        string `json:"Name"`
	MuscleGroup string `json:"MuscleGroup"`
	Machine     string `json:"Machine"`
}

type Exercise struct {
	Id int64 `json:"Id"`
	ExerciseInput
	StrengthHistory []Strength `json:"StrengthHistory"`
}

type ExerciseWithLastStrength struct {
	Exercise
	LastStrength Strength
}

type ExerciseForHistoryChart struct {
	Exercise
	Dates       []string
	Loads       []float64
	Repetitions []int64
}
