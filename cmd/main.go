package main

import (
	"log"
	"net/http"

	"github.com/lukas-arnold/strength-tracker/internal/configs"
	"github.com/lukas-arnold/strength-tracker/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handler.HandleViewExercises)
	mux.HandleFunc("GET /exercise/add", handler.HandleAddExercise)
	mux.HandleFunc("POST /exercise/add", handler.HandleAddExercise)
	mux.HandleFunc("GET /exercise/edit/{id}", handler.HandleEditExercise)
	mux.HandleFunc("POST /exercise/save/{id}", handler.HandleSaveExercise)
	mux.HandleFunc("GET /exercise/delete/{id}", handler.HandleDeleteExercise)
	mux.HandleFunc("GET /strength/add/{id}", handler.HandleAddStrength)
	mux.HandleFunc("POST /strength/add/{id}", handler.HandleAddStrength)
	mux.HandleFunc("GET /strength/delete/{id}/{createdAt}", handler.HandleDeleteStrength)

	mux.HandleFunc("GET /exercises", handler.HandleGetExercises)
	mux.HandleFunc("GET /exercise/{id}", handler.HandleGetExercise)
	mux.HandleFunc("POST /exercise", handler.HandleAddExercise)
	mux.HandleFunc("PUT /exercise/{id}", handler.HandleUpdateExercise)
	mux.HandleFunc("DELETE /exercise/{id}", handler.HandleDeleteExercise)
	log.Fatal(http.ListenAndServe(configs.PORT, mux))
}
