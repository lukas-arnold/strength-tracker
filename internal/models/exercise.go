package models

type ExerciseInput struct {
	Name        string `json:"name"`
	MuscleGroup string `json:"muscleGroup"`
	Machine     string `json:"machine"`
}

type Exercise struct {
	Id int64 `json:"id"`
	ExerciseInput
	StrengthHistory []Strength `json:"strengthHistory"`
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
