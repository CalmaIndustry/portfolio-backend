package api

import (
	"encoding/json"
	"net/http"
)

type Knowledge struct {
	ID    int    `json:"ID"`
	Title string `json:"title"`
}

func GetKnowledge(w http.ResponseWriter, req *http.Request) {
	knowledges := []Knowledge{
		{ID: 1, Title: "KNOWKNOW"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(knowledges)

}
