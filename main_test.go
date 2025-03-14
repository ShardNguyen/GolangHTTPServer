package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/ShardNguyen/GolangCounter/api"
)

func TestMain(t *testing.T) {
	server := api.NewServer()
	mux := http.NewServeMux()
	// These handle functions takes in a pattern as string and a callback function
	// This callback function must have http.ResponseWriter and *http.Request
	mux.HandleFunc("/", server.HandleRoot)
	mux.HandleFunc("POST /users", server.CreateUser)
	mux.HandleFunc("GET /users/{id}", server.GetUser)
	mux.HandleFunc("DELETE /user/{id}", server.DeleteUser)

	fmt.Println("Server listening to :8080")
	// http.ListenAndServe(":8080", mux)
}
