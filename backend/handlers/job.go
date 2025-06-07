package handlers

import (
	"encoding/json"
	"net/http"
)

type JobDescription struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func UploadJobDescription(w http.ResponseWriter, r *http.Request) {
	var jd JobDescription

	err := json.NewDecoder(r.Body).Decode(&jd)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// In a real app, you'd save this to MongoDB or process it
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Job description received",
		"title":   jd.Title,
	})
}
