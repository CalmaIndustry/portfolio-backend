package main

import (
	"fmt"
	"net/http"

	"portfolio-backend/api"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/skills", api.GetSkill)
	mux.HandleFunc("/api/knowledges", api.GetKnowledge)
	mux.HandleFunc("/api/projects", api.GetProject)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Cannot start http server: %v", err)
	}

}
