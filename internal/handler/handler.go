package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/lukas-arnold/strength-tracker/internal/configs"
	"github.com/lukas-arnold/strength-tracker/internal/language"
	"github.com/lukas-arnold/strength-tracker/internal/storage"
)

func handleError(w http.ResponseWriter, err error, statusCode int) {
	log.Printf("HTTP %d: %v", statusCode, err)
	http.Error(w, http.StatusText(statusCode), statusCode)
}

func HandleServiceWorker(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, configs.GetWebFiles(), "service-worker.js")
}

func HandleFiles(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/web/", http.FileServerFS(configs.GetWebFiles())).ServeHTTP(w, r)
}

func HandleView(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(
		template.New("base.html").Funcs(template.FuncMap{
			"T": func(key string) string {
				return language.T(configs.GetLanguage(), key)
			},
		}).ParseFS(configs.GetWebFiles(), "templates/base.html", "templates/index.html"),
	)
	exercises, err := storage.GetExercisesWithLastStrength()
	if err != nil {
		handleError(w, err, 404)
		return
	}
	err = tmpl.Execute(w, exercises)
	if err != nil {
		handleError(w, err, 500)
		return
	}
}
