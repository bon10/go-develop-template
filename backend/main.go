package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	r.HandleFunc("/greet/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		fmt.Fprintf(w, "Hello, %s!", name)
	})

	port := "8080"
	server := &http.Server{
		Addr:              ":" + port,
		Handler:           r,
		ReadHeaderTimeout: 3 * time.Second,
	}

	fmt.Printf("Starting server on port %s...\n", port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
