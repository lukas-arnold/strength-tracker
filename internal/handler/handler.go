package handler

import (
	"io"
	"log"
	"net/http"

	"github.com/lukas-arnold/strength-tracker/internal/storage"
	"github.com/lukas-arnold/strength-tracker/internal/utils"
)

func errorHandling(w http.ResponseWriter, httpStatusCode int) {
	w.WriteHeader(httpStatusCode)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("web/templates"))
}

func HandleGetExercises(w http.ResponseWriter, r *http.Request) {
	bytes, err := storage.GetExercisesBytes()
	if err != nil {
		errorHandling(w, 404)
		log.Print(err)
	}
	w.Write(bytes)
}

func HandleGetExercise(w http.ResponseWriter, r *http.Request) {
	id_str := r.PathValue("id")
	id, err := utils.ConvertId(id_str)
	if err != nil {
		errorHandling(w, 404)
		log.Print(err)
	}
	exercise, err := storage.GetExercise(id)
	if err != nil {
		errorHandling(w, 404)
		log.Print(err)
	}
	bytes, err := utils.ConvertExerciseToBytes(exercise)
	if err != nil {
		errorHandling(w, 404)
		log.Print(err)
	}
	w.Write(bytes)
}

func HandleAddExercise(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		errorHandling(w, 400)
		log.Print(err)
	}
	exercise, err := utils.ConvertBytesToAddExercise(bytes)
	if err != nil {
		errorHandling(w, 400)
		log.Print(err)
	}
	err = storage.AddExercise(exercise)
	if err != nil {
		errorHandling(w, 400)
		log.Print(err)
	}
	w.WriteHeader(201)
}

func HandleUpdateExercise(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		errorHandling(w, 400)
		log.Print(err)
	}
	exercise, err := utils.ConvertBytesToExercise(bytes)
	if err != nil {
		errorHandling(w, 400)
		log.Print(err)
	}
	err = storage.UpdateExercise(exercise)
	if err != nil {
		errorHandling(w, 400)
		log.Print(err)
	}
	w.WriteHeader(201)
}

func HandleDeleteExercise(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	err := storage.DeleteExercise(name)
	if err != nil {
		errorHandling(w, 404)
		log.Print(err)
	}
	w.WriteHeader(204)
}
