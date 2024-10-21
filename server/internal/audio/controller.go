package audio

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func InitController(router *chi.Mux) {

	router.Route("/audio", func(r chi.Router) {

		// Create audio
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			var newAudio AudioEntity

			err := json.NewDecoder(r.Body).Decode(&newAudio)
			fmt.Println(newAudio)

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			audio, err := Create(newAudio)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			jsonData, err := json.Marshal(audio)
			if err != nil {
				http.Error(w, "Не удалось преобразовать данные в JSON", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonData)
		})

		// Get audio
		r.Get("/{audioID}", func(w http.ResponseWriter, r *http.Request) {
			audioId, err := strconv.ParseUint(chi.URLParam(r, "audioID"), 10, 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			audio, err := GetAudio(audioId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			jsonData, err := json.Marshal(audio)
			if err != nil {
				http.Error(w, "Не удалось преобразовать данные в JSON", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonData)
		})

		// Get audios
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {

			projectId, _ := strconv.ParseUint(r.URL.Query().Get("projectId"), 10, 64)
			pageStr := r.URL.Query().Get("page")
			limitStr := r.URL.Query().Get("limit")
			page, err := strconv.Atoi(pageStr)
			if err != nil || page < 1 {
				page = 1
			}

			limit, err := strconv.Atoi(limitStr)
			if err != nil || limit < 1 {
				limit = 10
			}

			offset := (page - 1) * limit

			audio, err := GetAudios(projectId, limit, offset)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			jsonData, err := json.Marshal(audio)
			if err != nil {
				http.Error(w, "Не удалось преобразовать данные в JSON", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonData)
		})

		// Edit audio
		r.Patch("/{audioID}", func(w http.ResponseWriter, r *http.Request) {

		})

		// Delete audio
		r.Delete("/{audioID}", func(w http.ResponseWriter, r *http.Request) {

		})
	})

}
