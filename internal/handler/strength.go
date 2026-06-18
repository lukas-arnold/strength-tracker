package handler

import (
	"html/template"
	"net/http"

	"github.com/lukas-arnold/strength-tracker/internal/configs"
	"github.com/lukas-arnold/strength-tracker/internal/models"
	"github.com/lukas-arnold/strength-tracker/internal/storage"
	"github.com/lukas-arnold/strength-tracker/internal/utils"
)

func HandleAddStrengthGet(w http.ResponseWriter, r *http.Request) {
	exerciseId, err := utils.ConvertId(r.PathValue("exerciseId"))
	if err != nil {
		handleError(w, err, 500)
		return
	}
	tmpl := template.Must(
		template.New("base.html").
			Funcs(getTemplateFuncs()).
			ParseFS(configs.GetWebFiles(), "templates/base.html", "templates/strength/add.html"),
	)
	exercise, err := storage.GetExercise(exerciseId)
	if err != nil {
		handleError(w, err, 404)
		return
	}
	err = tmpl.Execute(w, exercise)
}

func HandleAddStrengthPost(w http.ResponseWriter, r *http.Request) {
	exerciseId, err := utils.ConvertId(r.PathValue("exerciseId"))
	if err != nil {
		handleError(w, err, 500)
		return
	}
	date := r.FormValue("date")
	loadForm := r.FormValue("load")
	load, err := utils.ConvertLoad(loadForm)
	if err != nil {
		handleError(w, err, 500)
		return
	}
	storage.AddStrength(exerciseId, models.StrengthInput{Date: date, Load: load})
	http.Redirect(w, r, "/", http.StatusFound)
}

func HandleEditStrength(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(
		template.New("base.html").
			Funcs(getTemplateFuncs()).
			ParseFS(configs.GetWebFiles(), "templates/base.html", "templates/strength/edit.html"),
	)
	id, err := utils.ConvertId(r.PathValue("id"))
	if err != nil {
		handleError(w, err, 500)
		return
	}
	strength, err := storage.GetStrength(id)
	if err != nil {
		handleError(w, err, 404)
		return
	}
	err = tmpl.Execute(w, strength)
	if err != nil {
		handleError(w, err, 500)
		return
	}
}

func HandleSaveStrength(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ConvertId(r.PathValue("id"))
	if err != nil {
		handleError(w, err, 500)
		return
	}
	strength, err := storage.GetStrength(id)
	if err != nil {
		handleError(w, err, 404)
		return
	}
	loadForm := r.FormValue("load")
	load, err := utils.ConvertLoad(loadForm)
	if err != nil {
		handleError(w, err, 500)
		return
	}
	strength.Date = r.FormValue("date")
	strength.Load = load
	err = storage.UpdateStrength(strength)
	if err != nil {
		handleError(w, err, 500)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func HandleDeleteStrength(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ConvertId(r.PathValue("id"))
	if err != nil {
		handleError(w, err, 500)
		return
	}
	err = storage.DeleteStrength(id)
	if err != nil {
		handleError(w, err, 404)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
