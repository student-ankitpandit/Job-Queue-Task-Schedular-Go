package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /task", func(w http.ResponseWriter, r *http.Request) {
		
	})

	http.ListenAndServe(":8080", mux)
}

