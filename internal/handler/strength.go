package handler

import (
	"net/http"

	"github.com/lukas-arnold/strength-tracker/internal/models"
	"github.com/lukas-arnold/strength-tracker/internal/utils"
)

func (h *Handler) HandleAddStrengthGet(
	w http.ResponseWriter,
	r *http.Request,
) {
	id, err := utils.ConvertToInt(
		r.PathValue("exerciseId"),
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
		"templates/strength/add.html",
		exercise,
	)
}

func (h *Handler) HandleAddStrengthPost(
	w http.ResponseWriter,
	r *http.Request,
) {

	exerciseId, err :=
		utils.ConvertToInt(
			r.PathValue("exerciseId"),
		)

	if err != nil {
		handleError(w, err, 500)
		return
	}

	load, err :=
		utils.ConvertLoad(
			r.FormValue("load"),
		)

	if err != nil {
		handleError(w, err, 500)
		return
	}

	repetitions, err :=
		utils.ConvertToInt(
			r.FormValue("repetitions"),
		)

	if err != nil {
		handleError(w, err, 500)
		return
	}

	err =
		h.store.AddStrength(
			exerciseId,
			models.StrengthInput{
				Date:        r.FormValue("date"),
				Load:        load,
				Repetitions: repetitions,
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

func (h *Handler) HandleEditStrength(
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

	strength, err := h.store.GetStrength(id)

	if err != nil {
		handleError(w, err, 404)
		return
	}

	h.renderTemplate(
		w,
		"templates/strength/edit.html",
		strength,
	)
}

func (h *Handler) HandleSaveStrength(
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

	strength, err :=
		h.store.GetStrength(id)

	if err != nil {
		handleError(w, err, 404)
		return
	}

	load, err :=
		utils.ConvertLoad(
			r.FormValue("load"),
		)

	if err != nil {
		handleError(w, err, 500)
		return
	}

	repetitions, err :=
		utils.ConvertToInt(
			r.FormValue("repetitions"),
		)

	if err != nil {
		handleError(w, err, 500)
		return
	}

	strength.Date =
		r.FormValue("date")

	strength.Load =
		load

	strength.Repetitions =
		repetitions

	err =
		h.store.UpdateStrength(
			strength,
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

func (h *Handler) HandleDeleteStrength(
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
		h.store.DeleteStrength(id)

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
