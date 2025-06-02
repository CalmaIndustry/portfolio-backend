package main

import (
	"fmt"
	"net/http"

	"github.com/CalmaIndustry/portfolio-backend/api/skill"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/api/skill", skill.Skill)
	if err := http.ListenAndServe(":8090", mux); err != nil {
		fmt.Printf("Cannot start http server: %v", err)
	}

}
