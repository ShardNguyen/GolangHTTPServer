package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ShardNguyen/GolangCounter/pkg/handler"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/user/{id}", handler.GetUser).Methods("GET")
	r.HandleFunc("/api/users", handler.GetAllUser).Methods("GET")
	r.HandleFunc("/api/user", handler.CreateUser).Methods("POST")
	r.HandleFunc("/api/user/{id}", handler.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/user/{id}", handler.DeleteUser).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
