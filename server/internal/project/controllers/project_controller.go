package project_controllers

import (
	"animal-sound-recognizer/internal/project/entities"
	"animal-sound-recognizer/internal/project/repositories"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func InitProjectsController(router *chi.Mux) {

	router.Route("/projects", func(r chi.Router) {

		// Create project
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {

			var newProject project_entities.ProjectEntity

			decodeErr := json.NewDecoder(r.Body).Decode(&newProject)
			if decodeErr != nil {
				http.Error(w, decodeErr.Error(), http.StatusBadRequest)
				return
			}

			project, createErr := project_repository.Create(newProject)
			if createErr != nil {
				http.Error(w, decodeErr.Error(), http.StatusBadRequest)
				return
			}

			jsonData, err := json.Marshal(project)
			if err != nil {
				http.Error(w, "Не удалось преобразовать данные в JSON", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonData)
		})

		// Get project
		r.Get("/{projectID}", func(w http.ResponseWriter, r *http.Request) {
			projectID, err := strconv.ParseUint(chi.URLParam(r, "projectID"), 10, 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			project, err := project_repository.GetProject(projectID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			jsonData, err := json.Marshal(project)
			if err != nil {
				http.Error(w, "Не удалось преобразовать данные в JSON", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonData)
		})

		// Get projects
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {

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

			project, err := project_repository.GetProjects(limit, offset)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			jsonData, err := json.Marshal(project)
			if err != nil {
				http.Error(w, "Не удалось преобразовать данные в JSON", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonData)
		})

		// Edit project
		r.Patch("/{projectID}", func(w http.ResponseWriter, r *http.Request) {
			projectID, err := strconv.ParseUint(chi.URLParam(r, "projectID"), 10, 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			var project project_entities.ProjectEntity
			decodeErr := json.NewDecoder(r.Body).Decode(&project)
			if decodeErr != nil {
				http.Error(w, decodeErr.Error(), http.StatusBadRequest)
				return
			}

			err = project_repository.Update(projectID, project)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusOK)
		})

		// Delete project
		r.Delete("/{projectID}", func(w http.ResponseWriter, r *http.Request) {
			projectID, err := strconv.ParseUint(chi.URLParam(r, "projectID"), 10, 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			err = project_repository.Delete(projectID)

		})
	})

}
