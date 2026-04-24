package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/lukas-arnold/strength-tracker/internal/storage"
)

func errorHandling(w http.ResponseWriter, httpStatusCode int) {
	w.WriteHeader(httpStatusCode)
}

func HandleView(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
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

func HandleFiles(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/web/", http.FileServer(http.Dir("web"))).ServeHTTP(w, r)
}
