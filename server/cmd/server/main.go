package main

import (
	"animal-sound-recognizer/internal/audio"
	"animal-sound-recognizer/internal/export"
	"animal-sound-recognizer/internal/file_storage"
	"animal-sound-recognizer/internal/projects"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	projects.InitController(r)
	file_storage.InitController(r)
	audio.InitController(r)
	export.InitController(r)

	return http.ListenAndServe(":8080", r)
}
