package app_http

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	FORM_FILE_KEY = "uploadfile"
)

type FileSaver interface {
	Save(id, filename string, fileData io.Reader) error
}

type Handlers struct {
	logger *log.Logger
	files  FileSaver
}

func New(logger *log.Logger, files FileSaver) *Handlers {
	return &Handlers{logger, files}
}

func (h *Handlers) FileUploader(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Failed to parse request multipart data", http.StatusInternalServerError)
		h.logger.Println("Error parsing form:", err)
		return
	}

	file, header, err := r.FormFile(FORM_FILE_KEY)

	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		h.logger.Println("Error reading form:", err)
		return
	}

	defer file.Close()

	if err := h.files.Save(id, header.Filename, file); err != nil {
		http.Error(w, "Не удалось сохранить файл", http.StatusInternalServerError)
		h.logger.Println("Error saving file:", err)
		return
	}
}
