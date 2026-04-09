package models

type Exercise struct {
	Name        string `json:"Name"`
	MuscleGroup string `json:"MuscleGroup"`
}

type ExerciseAdd struct {
	Id int64 `json:"Id"`
	Exercise
}

type ExerciseFull struct {
	ExerciseAdd
	StrengthHistory []Strength `json:"StrengthHistory"`
}
