package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RankRequest struct {
	JobDescription string   `json:"job_description"`
	Resumes        []string `json:"resumes"`
}

type RankedResume struct {
	ResumeText string  `json:"resume_text"`
	Score      float64 `json:"score"`
}

func RankResumes(w http.ResponseWriter, r *http.Request) {
	var req RankRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	payload, _ := json.Marshal(req)

	resp, err := http.Post("http://localhost:5000/rank", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		http.Error(w, "Failed to call NLP service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}
