package export

import (
	"animal-sound-recognizer/internal/rest"
	"github.com/go-chi/chi/v5"
	"io"
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

		r.Get("/excel/project/{projectID}", func(w http.ResponseWriter, r *http.Request) {
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

		r.Get("/excel/audio/{audioID}", func(w http.ResponseWriter, r *http.Request) {
			audioID, err := strconv.ParseUint(chi.URLParam(r, "audioID"), 10, 64)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusBadRequest,
					Message: "get audioID error",
					Detail:  err.Error(),
				})
				return
			}

			file, filename := createExcelDataByAudioId(audioID)

			w.Header().Set("Content-Disposition", "attachment; filename="+filename)
			w.Header().Set("Content-Length", string(len(file)))
			w.Write(file)
		})

		r.Get("/zip/project/{projectID}", func(w http.ResponseWriter, r *http.Request) {
			projectID, err := strconv.ParseUint(chi.URLParam(r, "projectID"), 10, 64)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusBadRequest,
					Message: "get projectId error",
					Detail:  err.Error(),
				})
				return
			}

			archive, err := GetZipArchiveIntervalsByProject(projectID)
			defer archive.Close()
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusInternalServerError,
					Message: "Creating archive error",
					Detail:  err.Error(),
				})
				return
			}

			var contentLength int64
			info, err := archive.Stat()
			if err != nil {
				contentLength = 0
			} else {
				contentLength = info.Size()
			}

			w.Header().Set("Content-Disposition", "attachment; filename="+archive.Name())
			w.Header().Set("Content-Length", strconv.FormatInt(contentLength, 10))
			_, err = io.Copy(w, archive)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusBadRequest,
					Message: "Receive file error",
					Detail:  err.Error(),
				})
				return
			}
		})

		r.Get("/zip/audio/{audioID}", func(w http.ResponseWriter, r *http.Request) {
			audioID, err := strconv.ParseUint(chi.URLParam(r, "audioID"), 10, 64)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusBadRequest,
					Message: "get audioID error",
					Detail:  err.Error(),
				})
				return
			}

			archive, err := GetZipArchiveIntervalsByAudio(audioID)
			defer archive.Close()

			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusInternalServerError,
					Message: "Creating archive error",
					Detail:  err.Error(),
				})
				return
			}

			var contentLength int64
			info, err := archive.Stat()
			if err != nil {
				contentLength = 0
			} else {
				contentLength = info.Size()
			}

			w.Header().Set("Content-Disposition", "attachment; filename="+archive.Name())
			w.Header().Set("Content-Length", strconv.FormatInt(contentLength, 10))
			_, err = io.Copy(w, archive)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusBadRequest,
					Message: "Receive file error",
					Detail:  err.Error(),
				})
				return
			}
		})
	})
}
