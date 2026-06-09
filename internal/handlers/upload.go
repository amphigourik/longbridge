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
	w.Header().Set("Content-Type", "application/json")

	mr, err := r.MultipartReader()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(uploadResponse{OK: false, Error: "invalid multipart"})
		return
	}

	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(uploadResponse{OK: false, Error: "read error"})
			return
		}
		if part.FormName() != "file" {
			part.Close()
			continue
		}

		filename := filepath.Base(part.FileName())
		if filename == "" || filename == "." {
			part.Close()
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(uploadResponse{OK: false, Error: "no filename"})
			return
		}

		destPath := filepath.Join(uploadDir(), filename)

		if _, err := os.Stat(destPath); err == nil {
			part.Close()
			json.NewEncoder(w).Encode(uploadResponse{OK: true, Skipped: true})
			return
		}

		dest, err := os.Create(destPath)
		if err != nil {
			part.Close()
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(uploadResponse{OK: false, Error: "cannot create file"})
			return
		}

		if _, err := io.Copy(dest, part); err != nil {
			dest.Close()
			part.Close()
			os.Remove(destPath)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(uploadResponse{OK: false, Error: "write failed"})
			return
		}

		dest.Close()
		part.Close()
		json.NewEncoder(w).Encode(uploadResponse{OK: true})
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(uploadResponse{OK: false, Error: "no file"})
}
