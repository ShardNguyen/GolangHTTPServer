package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestMain(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)

	mux.HandleFunc("POST /users", createUser)
	mux.HandleFunc("GET /users/{id}", getUser)
	mux.HandleFunc("DELETE /user/{id}", deleteUser)

	fmt.Println("Server listening to :8080")
	// http.ListenAndServe(":8080", mux)
}
