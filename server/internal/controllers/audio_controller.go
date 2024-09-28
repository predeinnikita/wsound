package controllers

import (
	"animal-sound-recognizer/internal/handlers"
	"encoding/json"
	"fmt"
	"github.com/amanitaverna/go-mp3"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func InitAudioController(router *chi.Mux) {
	router.Post("/audio", func(w http.ResponseWriter, r *http.Request) {
		file, _, _ := r.FormFile()

		decoder, err := mp3.NewDecoder(file)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(400)
			return
		}

		response, err := handlers.GetAudioDuration(decoder)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(400)
			return
		}

		json, err := json.Marshal(response)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(json)
	})
}
