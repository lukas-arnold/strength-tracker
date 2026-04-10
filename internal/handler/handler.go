package handler

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/lukas-arnold/strength-tracker/internal/models"
	"github.com/lukas-arnold/strength-tracker/internal/storage"
	"github.com/lukas-arnold/strength-tracker/internal/utils"
)

func errorHandling(w http.ResponseWriter, httpStatusCode int) {
	w.WriteHeader(httpStatusCode)
}

func HandleViewExercises(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
	exercises, err := storage.GetExercises()
	if err != nil {
		errorHandling(w, 404)
		log.Print(err)
	}
	err = tmpl.Execute(w, exercises)
	if err != nil {
		errorHandling(w, 404)
		log.Print(err)
	}
}

func HandleEditExercise(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/edit.html"))
	id_str := r.PathValue("id")
	id, _ := utils.ConvertId(id_str)
	exercise, _ := storage.GetExercise(id)
	tmpl.Execute(w, exercise)
}

func HandleSaveExercise(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Save")
	id_str := r.PathValue("id")
	id, _ := utils.ConvertId(id_str)
	exercise, _ := storage.GetExercise(id)
	fmt.Println(exercise)
	exercise.MuscleGroup = r.FormValue("muscleGroup")
	exercise.Name = r.FormValue("name")
	storage.UpdateExercise(exercise)
	http.Redirect(w, r, "/", http.StatusFound)
}

func HandleAddExercise(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl := template.Must(template.ParseFiles("web/templates/add.html"))
		tmpl.Execute(w, nil)
	case "POST":
		muscleGroup := r.FormValue("muscleGroup")
		name := r.FormValue("name")
		storage.AddExercise(models.Exercise{Name: name, MuscleGroup: muscleGroup})
		http.Redirect(w, r, "/", http.StatusFound)
	}
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

// func HandleAddExercise(w http.ResponseWriter, r *http.Request) {
// 	bytes, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		errorHandling(w, 400)
// 		log.Print(err)
// 	}
// 	exercise, err := utils.ConvertBytesToAddExercise(bytes)
// 	if err != nil {
// 		errorHandling(w, 400)
// 		log.Print(err)
// 	}
// 	err = storage.AddExercise(exercise)
// 	if err != nil {
// 		errorHandling(w, 400)
// 		log.Print(err)
// 	}
// 	w.WriteHeader(201)
// }

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
	name := r.PathValue("id")
	err := storage.DeleteExercise(name)
	if err != nil {
		errorHandling(w, 404)
		log.Print(err)
	}
	// w.WriteHeader(204)
	http.Redirect(w, r, "/", http.StatusFound)
}
