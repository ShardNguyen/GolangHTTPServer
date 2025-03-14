package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/ShardNguyen/GolangCounter/entity"
)

type Server struct {
	mu      sync.RWMutex
	userMap map[int]entity.User
}

func NewServer() *Server {
	server := new(Server)
	server.userMap = make(map[int]entity.User)

	return server
}

func (s *Server) HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!!!")
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user) // Read requested JSON's file and decoded it into the user struct

	// Error handling: When there's an error in decoding
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Error handling: When the username is empty
	if user.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	// Add user to the database
	s.mu.Lock()
	s.userMap[len(s.userMap)+1] = user
	s.mu.Unlock()

	// Writing the response code from the website
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	// Retrieving ID from the entered path
	id, err := strconv.Atoi(r.PathValue("id"))

	// Error handling: When there's an error in retrieving the ID
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.mu.RLock()
	user, ok := s.userMap[id] // Search if the user is inside the database
	s.mu.RUnlock()

	// Error handling: When the user is not found in the database
	if !ok {
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}

	// Create a header for the HTTP Reply and set the corresponding keys and values inside that header
	w.Header().Set("Content-type", "application/json")
	j, err := json.Marshal(user) // Converting user struct into a JSON file

	// Error handling: When conversion fails
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return ok status and writes out the JSON file as an HTTP reply?
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Retrieve ID from path
	id, err := strconv.Atoi(r.PathValue("id"))

	// Error handling: When there's an error in retrieving the ID
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	// Error handling: When the user is not found in the database
	if _, ok := s.userMap[id]; !ok {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}
	// Otherwise just delete the user out of the database
	delete(s.userMap, id)
	s.mu.Unlock()

	// Return status
	w.WriteHeader(http.StatusNoContent)
}
