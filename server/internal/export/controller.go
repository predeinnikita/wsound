package export

import (
	"animal-sound-recognizer/internal/rest"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func InitController(router *chi.Mux) {

	router.Route("/export", func(r chi.Router) {

		r.Get("/excel", func(w http.ResponseWriter, r *http.Request) {
			file, filename := createExcelAllData()

			w.Header().Set("Content-Disposition", "attachment; filename="+filename)
			w.Header().Set("Content-Length", string(len(file)))
			w.Write(file)
		})

		r.Get("/excel/{projectID}", func(w http.ResponseWriter, r *http.Request) {
			projectID, err := strconv.ParseUint(chi.URLParam(r, "projectID"), 10, 64)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusBadRequest,
					Message: "get projectId error",
					Detail:  err.Error(),
				})
				return
			}

			file, filename := createExcelDataByProjectId(projectID)

			w.Header().Set("Content-Disposition", "attachment; filename="+filename)
			w.Header().Set("Content-Length", string(len(file)))
			w.Write(file)
		})

	})
}
