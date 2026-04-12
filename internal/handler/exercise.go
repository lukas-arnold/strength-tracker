package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/lukas-arnold/strength-tracker/internal/models"
	"github.com/lukas-arnold/strength-tracker/internal/storage"
	"github.com/lukas-arnold/strength-tracker/internal/utils"
)

func HandleAddExerciseGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/exercise/add.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		errorHandling(w, 500)
		log.Print(err)
	}
}

func HandleAddExercisePost(w http.ResponseWriter, r *http.Request) {
	err := storage.AddExercise(models.ExerciseInput{Name: r.FormValue("name"), MuscleGroup: r.FormValue("muscleGroup")})
	if err != nil {
		errorHandling(w, 500)
		log.Print(err)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func HandleEditExercise(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/exercise/edit.html"))
	id, err := utils.ConvertId(r.PathValue("id"))
	if err != nil {
		errorHandling(w, 500)
		log.Print(err)
	}
	exercise, err := storage.GetExercise(id)
	if err != nil {
		errorHandling(w, 404)
		log.Print(err)
	}
	err = tmpl.Execute(w, exercise)
	if err != nil {
		errorHandling(w, 500)
		log.Print(err)
	}
}

func HandleSaveExercise(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ConvertId(r.PathValue("id"))
	if err != nil {
		errorHandling(w, 500)
		log.Print(err)
	}
	exercise, err := storage.GetExercise(id)
	if err != nil {
		errorHandling(w, 404)
		log.Print(err)
	}
	exercise.Name = r.FormValue("name")
	exercise.MuscleGroup = r.FormValue("muscleGroup")
	err = storage.UpdateExercise(exercise)
	if err != nil {
		errorHandling(w, 500)
		log.Print(err)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func HandleDeleteExercise(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ConvertId(r.PathValue("id"))
	if err != nil {
		errorHandling(w, 500)
		log.Print(err)
	}
	err = storage.DeleteExercise(id)
	if err != nil {
		errorHandling(w, 404)
		log.Print(err)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
