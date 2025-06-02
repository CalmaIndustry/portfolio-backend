package api

import (
	"encoding/json"
	"net/http"
)

type Project struct {
	ID          int    `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

func GetProject(w http.ResponseWriter, req *http.Request) {
	projects := []Project{
		{ID: 1, Title: "Test", Description: "testtest"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}
