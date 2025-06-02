package main

import (
	"fmt"
	"net/http"
	"portfolio-backend/router"
)

func main() {

	mux := router.GetRouter()
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Cannot start http server: %v", err)
	}

}
