package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/amphigourik/longbridge/internal/views"
)

type uploadResponse struct {
	OK      bool   `json:"ok"`
	Skipped bool   `json:"skipped,omitempty"`
	Error   string `json:"error,omitempty"`
}

func uploadDir() string {
	if d := os.Getenv("UPLOAD_DIR"); d != "" {
		return d
	}
	return "/mnt/upload"
}

func UploadPage(w http.ResponseWriter, r *http.Request) {
	views.RenderTemplate(w, "upload.html", nil)
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(uploadResponse{OK: false, Error: "no file"})
		return
	}
	defer file.Close()

	destPath := filepath.Join(uploadDir(), filepath.Base(header.Filename))

	if _, err := os.Stat(destPath); err == nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(uploadResponse{OK: true, Skipped: true})
		return
	}

	dest, err := os.Create(destPath)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(uploadResponse{OK: false, Error: "cannot create file"})
		return
	}
	defer dest.Close()

	if _, err := io.Copy(dest, file); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(uploadResponse{OK: false, Error: "write failed"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(uploadResponse{OK: true})
}
