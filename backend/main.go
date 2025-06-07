package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Rounakshrivastava5/Resume-Ranking-App/backend/handlers"
)

func main() {
	// Add the resume upload handler
	http.HandleFunc("/api/rank", handlers.RankResumes)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Resume Ranking API is running!")
	})

	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
