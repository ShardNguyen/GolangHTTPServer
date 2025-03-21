package main

import (
	"net/http"

	"github.com/ShardNguyen/GolangCounter/pkg/data"
	"github.com/ShardNguyen/GolangCounter/pkg/handler"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	var db data.Database
	var h handler.BaseHandler

	db = data.GetMapDatabaseInstance()
	h = handler.NewUserHandler(db)

	r.HandleFunc("/api/user/{id}", h.Get).Methods("GET")
	r.HandleFunc("/api/users", h.GetAll).Methods("GET")
	r.HandleFunc("/api/user", h.Create).Methods("POST")
	r.HandleFunc("/api/user/{id}", h.Update).Methods("PUT")
	r.HandleFunc("/api/user/{id}", h.Delete).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
