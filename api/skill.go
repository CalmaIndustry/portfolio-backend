package api

import (
	"encoding/json"
	"net/http"
)

type Skill struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func GetSkill(w http.ResponseWriter, req *http.Request) {
	skills := []Skill{
		{ID: 1, Title: "Sysops/SRE", Description: "I’m basically a Linux ninja — give me a log file, and I’ll find how many requests came from that pesky IP. OSI model? I know it like the back of my hand. Network flow issues? Consider them debugged before your coffee’s even cooled."},
		{ID: 2, Title: "DevOps", Description: "CI/CD pipelines are my playground, Git is my ride-or-die, and GitOps? That’s my gospel. Need a Docker image for your Node.js app? I got your back"},
		{ID: 3, Title: "DEV ?! (golang)", Description: "Still climbing the Go mountain — interfaces are starting to make sense, but I’m definitely on the “I’ll need a lot more caffeine and project before being a pro."},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(skills)
}
