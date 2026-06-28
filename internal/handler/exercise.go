package handler

import (
	"net/http"

	"github.com/lukas-arnold/strength-tracker/internal/models"
	"github.com/lukas-arnold/strength-tracker/internal/utils"
)

func (h *Handler) HandleAddExerciseGet(
	w http.ResponseWriter,
	r *http.Request,
) {
	h.renderTemplate(
		w,
		"templates/exercise/add.html",
		nil,
	)
}

func (h *Handler) HandleAddExercisePost(
	w http.ResponseWriter,
	r *http.Request,
) {

	err := h.store.AddExercise(
		models.ExerciseInput{
			Name:        r.FormValue("name"),
			MuscleGroup: r.FormValue("muscleGroup"),
			Machine:     r.FormValue("machine"),
		},
	)

	if err != nil {
		handleError(w, err, 500)
		return
	}

	http.Redirect(
		w,
		r,
		"/",
		http.StatusFound,
	)
}

func (h *Handler) HandleEditExercise(
	w http.ResponseWriter,
	r *http.Request,
) {
	id, err := utils.ConvertToInt(
		r.PathValue("id"),
	)

	if err != nil {
		handleError(w, err, 500)
		return
	}

	exercise, err := h.store.GetExercise(id)

	if err != nil {
		handleError(w, err, 404)
		return
	}

	h.renderTemplate(
		w,
		"templates/exercise/edit.html",
		exercise,
	)
}

func (h *Handler) HandleSaveExercise(
	w http.ResponseWriter,
	r *http.Request,
) {

	id, err :=
		utils.ConvertToInt(
			r.PathValue("id"),
		)

	if err != nil {
		handleError(w, err, 500)
		return
	}

	exercise, err :=
		h.store.GetExercise(id)

	if err != nil {
		handleError(w, err, 404)
		return
	}

	exercise.Name =
		r.FormValue("name")

	exercise.MuscleGroup =
		r.FormValue("muscleGroup")

	exercise.Machine =
		r.FormValue("machine")

	err =
		h.store.UpdateExercise(
			exercise,
		)

	if err != nil {
		handleError(w, err, 500)
		return
	}

	http.Redirect(
		w,
		r,
		"/",
		http.StatusFound,
	)
}

func (h *Handler) HandleDeleteExercise(
	w http.ResponseWriter,
	r *http.Request,
) {

	id, err :=
		utils.ConvertToInt(
			r.PathValue("id"),
		)

	if err != nil {
		handleError(w, err, 500)
		return
	}

	err =
		h.store.DeleteExercise(id)

	if err != nil {
		handleError(w, err, 404)
		return
	}

	http.Redirect(
		w,
		r,
		"/",
		http.StatusFound,
	)
}

func (h *Handler) HandleHistory(
	w http.ResponseWriter,
	r *http.Request,
) {
	id, err := utils.ConvertToInt(
		r.PathValue("id"),
	)

	if err != nil {
		handleError(w, err, 500)
		return
	}

	exercise, err := h.store.GetExerciseForHistoryChart(id)

	if err != nil {
		handleError(w, err, 404)
		return
	}

	h.renderTemplate(
		w,
		"templates/exercise/history.html",
		exercise,
	)
}
