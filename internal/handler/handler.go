package handler

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/lukas-arnold/strength-tracker/internal/configs"
	"github.com/lukas-arnold/strength-tracker/internal/language"
	"github.com/lukas-arnold/strength-tracker/internal/storage"
)

type Handler struct {
	store *storage.Storage
}

func New(store *storage.Storage) *Handler {
	return &Handler{
		store: store,
	}
}

func handleError(
	w http.ResponseWriter,
	err error,
	statusCode int,
) {
	log.Printf(
		"HTTP %d: %v",
		statusCode,
		err,
	)

	http.Error(
		w,
		http.StatusText(statusCode),
		statusCode,
	)
}

func (h *Handler) renderTemplate(
	w http.ResponseWriter,
	file string,
	data any,
) {
	tmpl := template.Must(
		template.New("base.html").
			Funcs(getTemplateFuncs()).
			ParseFS(
				configs.GetWebFiles(),
				"templates/base.html",
				file,
			),
	)

	err := tmpl.Execute(
		w,
		data,
	)

	if err != nil {
		handleError(
			w,
			err,
			http.StatusInternalServerError,
		)
	}
}

func getTemplateFuncs() template.FuncMap {
	return template.FuncMap{

		"T": func(key string) string {
			return language.T(
				configs.GetLanguage(),
				key,
			)
		},

		"formatLoad": func(value float64) string {
			return strconv.FormatFloat(
				value,
				'f',
				-1,
				64,
			)
		},

		"formatRepetitions": func(value int64) string {
			return strconv.FormatInt(
				value,
				10,
			)
		},
	}
}

func (h *Handler) HandleView(
	w http.ResponseWriter,
	r *http.Request,
) {

	tmpl := template.Must(
		template.New("base.html").
			Funcs(getTemplateFuncs()).
			ParseFS(
				configs.GetWebFiles(),
				"templates/base.html",
				"templates/index.html",
			),
	)

	exercises, err :=
		h.store.GetExercisesWithLastStrength()

	if err != nil {
		handleError(
			w,
			err,
			404,
		)
		return
	}

	err = tmpl.Execute(
		w,
		exercises,
	)

	if err != nil {
		handleError(
			w,
			err,
			500,
		)
	}
}
