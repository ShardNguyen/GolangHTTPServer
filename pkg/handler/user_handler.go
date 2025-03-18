package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ShardNguyen/GolangCounter/pkg/data"
	"github.com/ShardNguyen/GolangCounter/pkg/entity"
	"github.com/gorilla/mux"
)

func responseWithJson(writer http.ResponseWriter, status int, object interface{}) {
	writer.Header().Set("Content-Type", "application/json") // ResponseWriter sets "Content-Type: application/json" in the HTTP header
	writer.WriteHeader(status)                              // ResponseWriter writes the status of the response into the header
	json.NewEncoder(writer).Encode(object)                  // ResponseWriter encodes the object into a Json file and adds to the response
}

// Basically get the highest ID and add 1 to it
func generateId(users map[int]entity.User) int {
	var maxId int

	for id := range users {
		if id > maxId {
			maxId = id
		}
	}

	return maxId + 1
}

func GetUser(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request) // Get variables from path and stores it into params
	id, err := strconv.Atoi(params["id"])

	// Error handling: When ID is not converted to int
	if err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
		return
	}

	// Find user with said ID
	user, ok := data.UserTestData[id]
	if !ok {
		responseWithJson(writer, http.StatusNotFound, map[string]string{"message": "User not found"})
	} else {
		// Convert user data to user response data
		ur, err := user.ConvertToResponse()
		if err != nil {
			responseWithJson(writer, http.StatusInternalServerError, map[string]string{"message": "There's an error in getting user's info"})
		} else {
			responseWithJson(writer, http.StatusOK, ur)
		}
	}
}

func GetAllUser(writer http.ResponseWriter, request *http.Request) {
	urSlice := []entity.UserResponse{}

	// Convert all user in map to user responses
	for _, user := range data.UserTestData {
		ur, err := user.ConvertToResponse()
		if err != nil {
			continue
		}
		urSlice = append(urSlice, ur)
	}

	// Response with OK Status and everything in the user slice
	responseWithJson(writer, http.StatusOK, urSlice)
}

func CreateUser(writer http.ResponseWriter, request *http.Request) {
	var ur entity.UserResponse
	// Read the Json file from the requested side
	// Decode the said Json file and assign the variables into newUser
	err := json.NewDecoder(request.Body).Decode(&ur)

	// Error handling: When the Json file from the requested side cannot be assigned into newUser
	if err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid requested body"})
		return
	}

	// Get new ID
	newID := generateId(data.UserTestData)
	ur.SetID(newID)

	// Convert response to user data
	newUser, err := ur.ConvertToUser()
	if err != nil {
		responseWithJson(writer, http.StatusInternalServerError, map[string]string{"message": "Cannot convert to user"})
		return
	}

	// Add user to the map and respond back with ok status
	data.UserTestData[newID] = newUser
	responseWithJson(writer, http.StatusCreated, ur)
}

// Basically the same as GetUser but adding a delete function
func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	// Error handling: When ID is not converted to int
	if err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
		return
	}

	// Find if user with said ID exists
	_, ok := data.UserTestData[id]
	if !ok {
		responseWithJson(writer, http.StatusNotFound, map[string]string{"message": "User not found"})
	} else {
		delete(data.UserTestData, id)
		responseWithJson(writer, http.StatusOK, map[string]string{"message": "User is deleted"})
	}
}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
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

	// Get ID for user conversion
	ur.SetID(id)

	// Find ID of the user to apply the changes
	_, ok := data.UserTestData[id]
	if !ok {
		responseWithJson(writer, http.StatusNotFound, map[string]string{"message": "User not found"})
	} else {
		if edittedUser, err := ur.ConvertToUser(); err != nil {
			responseWithJson(writer, http.StatusInternalServerError, map[string]string{"message": "Cannot edit this user"})
		} else {
			data.UserTestData[id] = edittedUser
			responseWithJson(writer, http.StatusOK, ur)
		}
	}
}
