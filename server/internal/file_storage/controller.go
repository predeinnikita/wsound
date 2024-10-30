package file_storage

import (
	"animal-sound-recognizer/internal/rest"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

func InitController(router *chi.Mux) {

	router.Route("/file-storage", func(router chi.Router) {

		// Upload File
		router.Post("/", func(w http.ResponseWriter, r *http.Request) {
			file, header, err := r.FormFile("file")
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusBadRequest,
					Message: "get file error",
					Detail:  err.Error(),
				})
				return
			}
			defer file.Close()
			fileInByte, _ := io.ReadAll(file)

			id, err := SaveFile(header.Filename, fileInByte)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusBadRequest,
					Message: "save file failed",
					Detail:  err.Error(),
				})
				return
			}

			jsonData, err := json.Marshal(SaveFileResponse{
				Id: id,
			})

			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusInternalServerError,
					Message: "internal server error",
					Detail:  err.Error(),
				})
				return
			}

			w.Write(jsonData)
		})

		// Get file
		router.Get("/{fileId}", func(w http.ResponseWriter, r *http.Request) {
			fileID := chi.URLParam(r, "fileId")

			file, filename, err := GetFile(fileID)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusBadRequest,
					Message: "file not found",
					Detail:  err.Error(),
				})
				return
			}

			w.Header().Set("Content-Disposition", "attachment; filename="+filename)
			w.Header().Set("Content-Length", string(len(file)))
			w.Write(file)
		})

		// Delete file
		router.Delete("/{fileId}", func(w http.ResponseWriter, r *http.Request) {
			fileID := chi.URLParam(r, "fileId")

			err := DeleteFile(fileID)
			if err != nil {
				rest.ResponseError(w, rest.Error{
					Status:  http.StatusBadRequest,
					Message: "delete file failed",
					Detail:  err.Error(),
				})
				return
			}
		})

	})
}
