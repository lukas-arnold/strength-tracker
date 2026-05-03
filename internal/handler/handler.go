package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/lukas-arnold/strength-tracker/internal/configs"
	"github.com/lukas-arnold/strength-tracker/internal/language"
	"github.com/lukas-arnold/strength-tracker/internal/storage"
)

func errorHandling(w http.ResponseWriter, httpStatusCode int) {
	w.WriteHeader(httpStatusCode)
}

func HandleServiceWorker(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, configs.GetWebFiles(), "service-worker.js")
}

func HandleFiles(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/web/", http.FileServerFS(configs.GetWebFiles())).ServeHTTP(w, r)
}

func HandleView(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(
		template.New("index.html").Funcs(template.FuncMap{
			"T": func(key string) string {
				return language.T(configs.GetLanguage(), key)
			},
		}).ParseFS(configs.GetWebFiles(), "templates/index.html"),
	)
	exercises, err := storage.GetExercisesWithLastStrength()
	if err != nil {
		errorHandling(w, 404)
		log.Print(err)
	}
	err = tmpl.Execute(w, exercises)
	if err != nil {
		errorHandling(w, 500)
		log.Print(err)
	}
}
