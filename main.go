package main

import (
	"GolangHTTPServer/controllers/handlers"
	"GolangHTTPServer/models/database"
	"GolangHTTPServer/models/database/postgresql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	var db database.Database
	var h handlers.Handler

	db, err := postgresql.GetPostgresqlInstance()

	if err != nil {
		fmt.Println(err)
		fmt.Println("Connection failed.")
		return
	}

	defer db.CloseConnection()
	h = handlers.NewUserHandler(&db)

	h = handlers.NewUserHandler(&db)

	r.HandleFunc("/api/user/{id}", h.Get).Methods("GET")
	r.HandleFunc("/api/users", h.GetAll).Methods("GET")
	r.HandleFunc("/api/user", h.Create).Methods("POST")
	r.HandleFunc("/api/user/{id}", h.Update).Methods("PUT")
	r.HandleFunc("/api/user/{id}", h.Delete).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
