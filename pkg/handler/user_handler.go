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
func generateId(users []entity.User) int {
	var maxId int
	for _, user := range users {
		if user.Id > maxId {
			maxId = user.Id
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
	for _, user := range data.UserTestData {
		if user.Id == id {
			// Response with OK Status and the user
			responseWithJson(writer, http.StatusOK, user)
			return
		}
	}

	// Response with Not Found Status and not found message
	responseWithJson(writer, http.StatusNotFound, map[string]string{"message": "User not found"})
}

func GetAllUser(writer http.ResponseWriter, request *http.Request) {
	// Response with OK Status and everything in the user slice
	responseWithJson(writer, http.StatusOK, data.UserTestData)
}

func CreateUser(writer http.ResponseWriter, request *http.Request) {
	var newUser entity.User
	// Read the Json file from the requested side
	// Decode the said Json file and assign the variables into newUser
	err := json.NewDecoder(request.Body).Decode(&newUser)

	// Error handling: When the Json file from the requested side cannot be assigned into newUser
	if err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	// Add user into the user slice
	newUser.Id = generateId(data.UserTestData)
	data.UserTestData = append(data.UserTestData, newUser)

	// Resposne with Created Status and new user to the requesting side
	responseWithJson(writer, http.StatusCreated, newUser)
}

// Basically the same as GetUser but adding a delete function
func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
		return
	}

	for i, user := range data.UserTestData {
		if user.Id == id {
			// Delete data
			data.UserTestData = append(data.UserTestData[:i], data.UserTestData[i+1:]...)
			responseWithJson(writer, http.StatusOK, map[string]string{"message": "User is deleted"})
		}
	}
	responseWithJson(writer, http.StatusNotFound, map[string]string{"message": "User not found"})
}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	// Error handling: When ID is not converted to int
	if err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
		return
	}

	var updateUser entity.User
	err = json.NewDecoder(request.Body).Decode(&updateUser)
	// Error handling: When the Json file from the requested side cannot be assigned into updateUser
	if err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}
	updateUser.Id = id

	// Find ID of the user to apply the changes
	for i, user := range data.UserTestData {
		if user.Id == id {
			data.UserTestData[i] = updateUser
			responseWithJson(writer, http.StatusOK, updateUser)
			return
		}
	}

	responseWithJson(writer, http.StatusNotFound, map[string]string{"message": "User not found"})
}
