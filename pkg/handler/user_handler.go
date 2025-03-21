package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ShardNguyen/GolangCounter/pkg/data"
	"github.com/ShardNguyen/GolangCounter/pkg/entity"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	db data.Database
}

func NewUserHandler(db data.Database) *UserHandler {
	uh := new(UserHandler)
	uh.db = db
	return uh
}

func (uh UserHandler) SetDB(db data.Database) {
	uh.db = db
}

func (uh UserHandler) Get(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request) // Get variables from path and stores it into params
	id, err := strconv.Atoi(params["id"])

	// Error handling: When ID is not converted to int
	if err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
		return
	}

	// Find user with said ID
	user, err := uh.db.GetUser(id)

	if err != nil {
		responseWithJson(writer, http.StatusNotFound, map[string]string{"message": "User not found"})
		return
	}

	// Convert user data to user response data
	ur, err := user.ConvertToResponse()
	if err != nil {
		responseWithJson(writer, http.StatusInternalServerError, map[string]string{"message": "There's an error in converting user to user response"})
		return
	}

	responseWithJson(writer, http.StatusOK, ur)
}

func (uh UserHandler) GetAll(writer http.ResponseWriter, request *http.Request) {
	urSlice := []entity.UserResponse{}
	userData, err := uh.db.GetAllUsers()

	if err != nil {
		responseWithJson(writer, http.StatusInternalServerError, map[string]string{"message": "Cannot get user data"})
	}

	// Convert all user in map to user responses
	for _, user := range userData {
		ur, err := user.ConvertToResponse()

		if err != nil {
			continue
		}

		urSlice = append(urSlice, ur)
	}

	// Response with OK Status and everything in the user slice
	responseWithJson(writer, http.StatusOK, urSlice)
}

func (uh UserHandler) Create(writer http.ResponseWriter, request *http.Request) {
	var ur entity.UserResponse
	// Read the Json file from the requested side
	// Decode the said Json file and assign the variables into the public user response for converting later
	err := json.NewDecoder(request.Body).Decode(&ur)

	// Error handling: When the Json file from the requested side cannot be assigned into newUser
	if err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid requested body"})
		return
	}

	// Tell the database to create a user
	if err := uh.db.CreateUser(&ur); err != nil {
		responseWithJson(writer, http.StatusInternalServerError, map[string]string{"message": "Cannot convert to user"})
		return
	}

	responseWithJson(writer, http.StatusCreated, ur)
}

// Basically the same as GetUser but adding a delete function
func (uh UserHandler) Delete(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	// Error handling: When ID is not converted to int
	if err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
		return
	}

	err = uh.db.DeleteUser(id)
	if err != nil {
		responseWithJson(writer, http.StatusNotFound, map[string]string{"message": "User not found"})
		return
	}

	responseWithJson(writer, http.StatusOK, map[string]string{"message": "User is deleted"})
}

func (uh UserHandler) Update(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	// Error handling: When ID is not converted to int
	if err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
		return
	}

	var ur entity.UserResponse
	err = json.NewDecoder(request.Body).Decode(&ur)

	// Error handling: When the Json file from the requested side cannot be assigned into updateUser
	if err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	err = uh.db.UpdateUser(id, &ur)

	if err.Error() == "user not found" {
		responseWithJson(writer, http.StatusNotFound, map[string]string{"message": "User not found"})
		return
	}

	if err.Error() == "cannot edit this user" {
		responseWithJson(writer, http.StatusInternalServerError, map[string]string{"message": "Cannot edit this user"})
		return
	}

	responseWithJson(writer, http.StatusOK, ur)
}

func responseWithJson(writer http.ResponseWriter, status int, object interface{}) {
	writer.Header().Set("Content-Type", "application/json") // ResponseWriter sets "Content-Type: application/json" in the HTTP header
	writer.WriteHeader(status)                              // ResponseWriter writes the status of the response into the header
	json.NewEncoder(writer).Encode(object)                  // ResponseWriter encodes the object into a Json file and adds to the response
}
