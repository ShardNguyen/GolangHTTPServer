// TO DO LIST
// Figure out how to make an error logging for wrong password? (This one definitely got me not figuring it out for so long)

package main

import (
	"fmt"
	"net/http"

	"GolangHTTPServer/pkg/data"
	"GolangHTTPServer/pkg/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	var db data.Database
	var h handler.BaseHandler

	db, err := data.GetPostgreSQLInstance()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Connection failed.")
		return
	}

	h = handler.NewUserHandler(&db)

	r.HandleFunc("/api/user/{id}", h.Get).Methods("GET")
	r.HandleFunc("/api/users", h.GetAll).Methods("GET")
	r.HandleFunc("/api/user", h.Create).Methods("POST")
	r.HandleFunc("/api/user/{id}", h.Update).Methods("PUT")
	r.HandleFunc("/api/user/{id}", h.Delete).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
