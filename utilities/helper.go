package utilities

import (
	"GolangHTTPServer/models/entity"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetIDFromRequest(request *http.Request) (int, error) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		return 0, err
	}

	return id, nil
}

func DecodeJsonFromRequest(request *http.Request) (*entity.User, error) {
	var u entity.User
	err := json.NewDecoder(request.Body).Decode(&u)

	if err != nil {
		return nil, err
	}

	return &u, nil
}
