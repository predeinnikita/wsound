package projects

import (
	"animal-sound-recognizer/internal/rest"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func InitController(router *chi.Mux) {

	router.Route("/projects", func(r chi.Router) {

		r.Post("/", func(w http.ResponseWriter, r *http.Request) {

			var newProject Project

			decodeErr := json.NewDecoder(r.Body).Decode(&newProject)
			if decodeErr != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusBadRequest,
					Message: "invalid request body",
					Detail:  decodeErr.Error(),
				})
				return
			}

			project, createErr := Create(newProject)
			if createErr != nil {
				http.Error(w, decodeErr.Error(), http.StatusBadRequest)
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusBadRequest,
					Message: "create project error",
					Detail:  decodeErr.Error(),
				})
				return
			}

			jsonData, err := json.Marshal(project)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusInternalServerError,
					Message: "project to json error",
					Detail:  decodeErr.Error(),
				})
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonData)
		})

		r.Get("/{projectID}", func(w http.ResponseWriter, r *http.Request) {
			projectID, err := strconv.ParseUint(chi.URLParam(r, "projectID"), 10, 64)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusBadRequest,
					Message: "get projectId error",
					Detail:  err.Error(),
				})
				return
			}

			project, err := GetProject(projectID)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusBadRequest,
					Message: "get project error",
					Detail:  err.Error(),
				})
				return
			}

			jsonData, err := json.Marshal(project)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusBadRequest,
					Message: "project to json error",
					Detail:  err.Error(),
				})
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonData)
		})

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

			project, err := GetProjects(limit, offset)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusBadRequest,
					Message: "get projects error",
					Detail:  err.Error(),
				})
				return
			}

			jsonData, err := json.Marshal(project)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusInternalServerError,
					Message: "projects to json error",
					Detail:  err.Error(),
				})
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonData)
		})

		r.Patch("/{projectID}", func(w http.ResponseWriter, r *http.Request) {
			projectID, err := strconv.ParseUint(chi.URLParam(r, "projectID"), 10, 64)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusBadRequest,
					Message: "get projectId error",
					Detail:  err.Error(),
				})
				return
			}

			var project Project
			decodeErr := json.NewDecoder(r.Body).Decode(&project)
			if decodeErr != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusBadRequest,
					Message: "project to json error",
					Detail:  err.Error(),
				})
				return
			}

			err = Update(projectID, project)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusInternalServerError,
					Message: "update project error",
					Detail:  err.Error(),
				})
				return
			}

			w.WriteHeader(http.StatusOK)
		})

		r.Delete("/{projectID}", func(w http.ResponseWriter, r *http.Request) {
			projectID, err := strconv.ParseUint(chi.URLParam(r, "projectID"), 10, 64)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusBadRequest,
					Message: "get projectId error",
					Detail:  err.Error(),
				})
				return
			}

			err = Delete(projectID)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusInternalServerError,
					Message: "delete project error",
					Detail:  err.Error(),
				})
				return
			}
		})
	})
}
