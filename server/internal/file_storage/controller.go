package file_storage

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

func InitController(router *chi.Mux) {

	router.Route("/file-storage", func(router chi.Router) {

		// Upload File
		router.Post("/", func(w http.ResponseWriter, r *http.Request) {
			file, header, _ := r.FormFile("file")
			defer file.Close()
			fileInByte, _ := io.ReadAll(file)

			id, err := SaveFile(header.Filename, fileInByte)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			jsonData, err := json.Marshal(SaveFileResponse{
				Id: id,
			})

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			w.Write(jsonData)
		})

		// Get file
		router.Get("/{fileId}", func(w http.ResponseWriter, r *http.Request) {
			fileID := chi.URLParam(r, "fileId")

			file, filename, err := GetFile(fileID)
			if err != nil {
				http.Error(w, fmt.Errorf("file not found").Error(), http.StatusNotFound)
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
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		})

	})
}
