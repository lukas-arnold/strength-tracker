package handler

import (
	"html/template"
	"net/http"

	"github.com/lukas-arnold/strength-tracker/internal/configs"
	"github.com/lukas-arnold/strength-tracker/internal/language"
	"github.com/lukas-arnold/strength-tracker/internal/models"
	"github.com/lukas-arnold/strength-tracker/internal/storage"
	"github.com/lukas-arnold/strength-tracker/internal/utils"
)

func HandleAddExerciseGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(
		template.New("base.html").Funcs(template.FuncMap{
			"T": func(key string) string {
				return language.T(configs.GetLanguage(), key)
			},
		}).ParseFS(configs.GetWebFiles(), "templates/base.html", "templates/exercise/add.html"),
	)
	err := tmpl.Execute(w, nil)
	if err != nil {
		handleError(w, err, 500)
		return
	}
}

func HandleAddExercisePost(w http.ResponseWriter, r *http.Request) {
	err := storage.AddExercise(models.ExerciseInput{Name: r.FormValue("name"), MuscleGroup: r.FormValue("muscleGroup")})
	if err != nil {
		handleError(w, err, 500)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func HandleEditExercise(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(
		template.New("base.html").Funcs(template.FuncMap{
			"T": func(key string) string {
				return language.T(configs.GetLanguage(), key)
			},
		}).ParseFS(configs.GetWebFiles(), "templates/base.html", "templates/exercise/edit.html"),
	)
	id, err := utils.ConvertId(r.PathValue("id"))
	if err != nil {
		handleError(w, err, 500)
		return
	}
	exercise, err := storage.GetExercise(id)
	if err != nil {
		handleError(w, err, 404)
		return
	}
	err = tmpl.Execute(w, exercise)
	if err != nil {
		handleError(w, err, 500)
		return
	}
}

func HandleSaveExercise(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ConvertId(r.PathValue("id"))
	if err != nil {
		handleError(w, err, 500)
		return
	}
	exercise, err := storage.GetExercise(id)
	if err != nil {
		handleError(w, err, 404)
		return
	}
	exercise.Name = r.FormValue("name")
	exercise.MuscleGroup = r.FormValue("muscleGroup")
	err = storage.UpdateExercise(exercise)
	if err != nil {
		handleError(w, err, 500)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func HandleDeleteExercise(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ConvertId(r.PathValue("id"))
	if err != nil {
		handleError(w, err, 500)
		return
	}
	err = storage.DeleteExercise(id)
	if err != nil {
		handleError(w, err, 404)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func HandleHistory(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(
		template.New("base.html").Funcs(template.FuncMap{
			"T": func(key string) string {
				return language.T(configs.GetLanguage(), key)
			},
		}).ParseFS(configs.GetWebFiles(), "templates/base.html", "templates/exercise/history.html"),
	)
	id, err := utils.ConvertId(r.PathValue("id"))
	if err != nil {
		handleError(w, err, 500)
		return
	}
	exercise, err := storage.GetExerciseForHistoryChart(id)
	if err != nil {
		handleError(w, err, 404)
		return
	}
	err = tmpl.Execute(w, exercise)
	if err != nil {
		handleError(w, err, 500)
		return
	}
}
