package api

import (
	"encoding/json"
	"net/http"
)

type Skill struct {
	ID    int    `json:"id"	`
	Title string `json:"title"`
}

func GetSkill(w http.ResponseWriter, req *http.Request) {
	skills := []Skill{
		{ID: 1, Title: "Go"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(skills)
}
