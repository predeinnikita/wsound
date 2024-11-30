package audio

import (
	"animal-sound-recognizer/internal/recognizer"
	"animal-sound-recognizer/internal/rest"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func InitController(router *chi.Mux) {

	router.Route("/audio", func(r chi.Router) {

		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			var newAudio Audio

			err := json.NewDecoder(r.Body).Decode(&newAudio)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusBadRequest,
					Message: "invalid request body",
					Detail:  err.Error(),
				})
				return
			}

			result := recognizer.ProcessAudio(newAudio.StorageID)
			if result.Result[0].IsWolf {
				newAudio.Status = WolfStatus
			}

			audio, err := Create(newAudio)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusBadRequest,
					Message: "create audio failed",
					Detail:  err.Error(),
				})
				return
			}

			jsonData, err := json.Marshal(audio)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusInternalServerError,
					Message: "internal server error",
					Detail:  err.Error(),
				})
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonData)
		})

		r.Get("/{audioID}", func(w http.ResponseWriter, r *http.Request) {
			audioId, err := strconv.ParseUint(chi.URLParam(r, "audioID"), 10, 64)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusBadRequest,
					Message: "get audioId error",
					Detail:  err.Error(),
				})
				return
			}

			audio, err := GetAudio(audioId)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusBadRequest,
					Message: "get audio failed",
					Detail:  err.Error(),
				})
				return
			}

			jsonData, err := json.Marshal(audio)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusInternalServerError,
					Message: "internal server error",
					Detail:  err.Error(),
				})
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonData)
		})

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {

			projectId, err := strconv.ParseUint(r.URL.Query().Get("projectId"), 10, 64)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusBadRequest,
					Message: "get projectId error",
					Detail:  err.Error(),
				})
				return
			}

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
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusInternalServerError,
					Message: "get audios failed",
					Detail:  err.Error(),
				})
				return
			}

			jsonData, err := json.Marshal(audio)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusInternalServerError,
					Message: "internal server error",
					Detail:  err.Error(),
				})
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonData)
		})

		r.Patch("/{audioID}", func(w http.ResponseWriter, r *http.Request) {
			audioId, err := strconv.ParseUint(chi.URLParam(r, "audioID"), 10, 64)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusBadRequest,
					Message: "get audioId error",
					Detail:  err.Error(),
				})
				return
			}

			var editAudioRequest EditAudioRequest
			err = json.NewDecoder(r.Body).Decode(&editAudioRequest)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusInternalServerError,
					Message: "internal server error",
					Detail:  err.Error(),
				})
				return
			}

			err = EditAudio(audioId, editAudioRequest)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusInternalServerError,
					Message: "internal server error",
					Detail:  err.Error(),
				})
				return
			}

			w.WriteHeader(http.StatusOK)
		})

		r.Delete("/{audioID}", func(w http.ResponseWriter, r *http.Request) {
			audioId, err := strconv.ParseUint(chi.URLParam(r, "audioID"), 10, 64)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusBadRequest,
					Message: "get audioId error",
					Detail:  err.Error(),
				})
				return
			}

			err = DeleteAudio(audioId)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusBadRequest,
					Message: "delete audio failed",
					Detail:  err.Error(),
				})
				return
			}
		})
	})

}
