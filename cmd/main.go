package main

import (
	"log"
	"net/http"

	"github.com/lukas-arnold/strength-tracker/internal/configs"
	"github.com/lukas-arnold/strength-tracker/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handler.HomePage)
	mux.HandleFunc("GET /exercises", handler.HandleGetExercises)
	mux.HandleFunc("GET /exercise/{id}", handler.HandleGetExercise)
	mux.HandleFunc("POST /exercise", handler.HandleAddExercise)
	mux.HandleFunc("PUT /exercise/{id}", handler.HandleUpdateExercise)
	mux.HandleFunc("DELETE /exercise/{id}", handler.HandleDeleteExercise)
	log.Fatal(http.ListenAndServe(configs.PORT, mux))
}
