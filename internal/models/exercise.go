package models

type ExerciseAdd struct {
	Name        string `json:"Name"`
	MuscleGroup string `json:"MuscleGroup"`
}

type ExerciseGet struct {
	Id int `json:"Id"`
	ExerciseAdd
	StrengthHistory []StrengthGet `json:"StrengtHistory"`
}
