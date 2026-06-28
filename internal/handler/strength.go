package handler

import (
	"net/http"

	"github.com/lukas-arnold/strength-tracker/internal/models"
	"github.com/lukas-arnold/strength-tracker/internal/utils"
)

func (h *Handler) HandleAddStrengthGet(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ConvertToInt(r.PathValue("exerciseId"))
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	exercise, err := h.store.GetExercise(id)
	if err != nil {
		handleError(w, err, http.StatusNotFound)
		return
	}

	h.renderTemplate(w, "templates/strength/add.html", exercise)
}

func (h *Handler) HandleAddStrengthPost(w http.ResponseWriter, r *http.Request) {
	exerciseId, err := utils.ConvertToInt(r.PathValue("exerciseId"))
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	load, err := utils.ConvertLoad(r.FormValue("load"))
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	reps, err := utils.ConvertToInt(r.FormValue("repetitions"))
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	input := models.StrengthInput{
		Date:        r.FormValue("date"),
		Load:        load,
		Repetitions: reps,
	}

	if err := h.store.AddStrength(exerciseId, input); err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Handler) HandleEditStrength(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ConvertToInt(r.PathValue("id"))
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	strength, err := h.store.GetStrength(id)
	if err != nil {
		handleError(w, err, http.StatusNotFound)
		return
	}

	h.renderTemplate(w, "templates/strength/edit.html", strength)
}

func (h *Handler) HandleSaveStrength(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ConvertToInt(r.PathValue("id"))
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	strength, err := h.store.GetStrength(id)
	if err != nil {
		handleError(w, err, http.StatusNotFound)
		return
	}

	load, err := utils.ConvertLoad(r.FormValue("load"))
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	reps, err := utils.ConvertToInt(r.FormValue("repetitions"))
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	strength.Date = r.FormValue("date")
	strength.Load = load
	strength.Repetitions = reps

	if err := h.store.UpdateStrength(strength); err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Handler) HandleDeleteStrength(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ConvertToInt(r.PathValue("id"))
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	if err := h.store.DeleteStrength(id); err != nil {
		handleError(w, err, http.StatusNotFound)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
