package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/lukas-arnold/strength-tracker/internal/models"
	"github.com/lukas-arnold/strength-tracker/internal/storage"
	"github.com/lukas-arnold/strength-tracker/internal/utils"
)

func HandleAddStrengthGet(w http.ResponseWriter, r *http.Request) {
	exerciseId, err := utils.ConvertId(r.PathValue("exerciseId"))
	if err != nil {
		errorHandling(w, 500)
		log.Print(err)
	}
	tmpl := template.Must(template.ParseFiles("web/templates/strength/add.html"))
	exercise, err := storage.GetExercise(exerciseId)
	if err != nil {
		errorHandling(w, 404)
		log.Print(err)
	}
	err = tmpl.Execute(w, exercise)
}

func HandleAddStrengthPost(w http.ResponseWriter, r *http.Request) {
	exerciseId, err := utils.ConvertId(r.PathValue("exerciseId"))
	if err != nil {
		errorHandling(w, 500)
		log.Print(err)
	}
	date := r.FormValue("date")
	loadForm := r.FormValue("load")
	load, err := utils.ConvertLoad(loadForm)
	if err != nil {
		errorHandling(w, 500)
		log.Print(err)
	}
	storage.AddStrength(exerciseId, models.StrengthInput{Date: date, Load: load})
	http.Redirect(w, r, "/", http.StatusFound)
}

func HandleEditStrength(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/strength/edit.html"))
	id, err := utils.ConvertId(r.PathValue("id"))
	if err != nil {
		errorHandling(w, 500)
		log.Print(err)
	}
	exercise, err := storage.GetExerciseByStrength(id)
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

func HandleSaveStrength(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ConvertId(r.PathValue("id"))
	if err != nil {
		errorHandling(w, 500)
		log.Print(err)
	}
	strength, err := storage.GetStrength(id)
	if err != nil {
		errorHandling(w, 404)
		log.Print(err)
	}
	loadForm := r.FormValue("load")
	load, err := utils.ConvertLoad(loadForm)
	if err != nil {
		errorHandling(w, 500)
		log.Print(err)
	}
	strength.Date = r.FormValue("date")
	strength.Load = load
	err = storage.UpdateStrength(strength)
	if err != nil {
		errorHandling(w, 500)
		log.Print(err)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func HandleDeleteStrength(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ConvertId(r.PathValue("id"))
	if err != nil {
		errorHandling(w, 500)
		log.Print(err)
	}
	err = storage.DeleteStrength(id)
	if err != nil {
		errorHandling(w, 404)
		log.Print(err)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
