package router

import (
	"net/http"
	"portfolio-backend/api"
)

func GetRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/skills", api.GetSkill)
	mux.HandleFunc("/api/knowledges", api.GetKnowledge)
	mux.HandleFunc("/api/projects", api.GetProject)

	fs := http.FileServer(http.Dir("./public/images"))
	mux.Handle("/images/", http.StripPrefix("/images/", fs))

	return mux
}
