package handler

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("GET /", h.HandleView)

	mux.HandleFunc("GET /web/", h.HandleFiles)
	mux.HandleFunc("GET /service-worker.js", h.HandleServiceWorker)

	mux.HandleFunc("GET /exercise/add", h.HandleAddExerciseGet)
	mux.HandleFunc("POST /exercise/add", h.HandleAddExercisePost)
	mux.HandleFunc("GET /exercise/edit/{id}", h.HandleEditExercise)
	mux.HandleFunc("POST /exercise/save/{id}", h.HandleSaveExercise)
	mux.HandleFunc("GET /exercise/delete/{id}", h.HandleDeleteExercise)
	mux.HandleFunc("GET /exercise/history/{id}", h.HandleHistory)

	mux.HandleFunc("GET /strength/add/{exerciseId}", h.HandleAddStrengthGet)
	mux.HandleFunc("POST /strength/add/{exerciseId}", h.HandleAddStrengthPost)
	mux.HandleFunc("GET /strength/edit/{id}", h.HandleEditStrength)
	mux.HandleFunc("POST /strength/save/{id}", h.HandleSaveStrength)
	mux.HandleFunc("GET /strength/delete/{id}", h.HandleDeleteStrength)
}
