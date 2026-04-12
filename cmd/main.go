package main

import (
	"log"
	"net/http"

	"github.com/lukas-arnold/strength-tracker/internal/configs"
	"github.com/lukas-arnold/strength-tracker/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handler.HandleView)

	mux.HandleFunc("GET /exercise/add", handler.HandleAddExerciseGet)
	mux.HandleFunc("POST /exercise/add", handler.HandleAddExercisePost)
	mux.HandleFunc("GET /exercise/edit/{id}", handler.HandleEditExercise)
	mux.HandleFunc("POST /exercise/save/{id}", handler.HandleSaveExercise)
	mux.HandleFunc("GET /exercise/delete/{id}", handler.HandleDeleteExercise)
	mux.HandleFunc("GET /exercise/history/{id}", handler.HandleHistory)

	mux.HandleFunc("GET /strength/add/{exerciseId}", handler.HandleAddStrengthGet)
	mux.HandleFunc("POST /strength/add/{exerciseId}", handler.HandleAddStrengthPost)
	mux.HandleFunc("GET /strength/edit/{id}", handler.HandleEditStrength)
	mux.HandleFunc("POST /strength/save/{id}", handler.HandleSaveStrength)
	mux.HandleFunc("GET /strength/delete/{id}", handler.HandleDeleteStrength)

	log.Fatal(http.ListenAndServe(configs.PORT, mux))
}
