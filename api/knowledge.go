package api

import (
	"encoding/json"
	"net/http"
)

type Knowledge struct {
	ID    int    `json:"ID"`
	Title string `json:"title"`
	Image string `json:"image"`
}

func GetKnowledge(w http.ResponseWriter, req *http.Request) {
	knowledges := []Knowledge{
		{ID: 1, Title: "cka", Image: "/images/cka.png"},
		{ID: 2, Title: "cks", Image: "/images/cks.png"},
		{ID: 3, Title: "sap", Image: "/images/sap.png"},
		{ID: 4, Title: "sysops", Image: "/images/sysops.png"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(knowledges)

}
