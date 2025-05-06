package handlers

import (
	"GolangHTTPServer/controllers"
	"GolangHTTPServer/models/database"
	"GolangHTTPServer/utilities"
	"net/http"
)

type UserHandler struct {
	db database.Database
}

func NewUserHandler(db *database.Database) *UserHandler {
	uh := new(UserHandler)
	uh.db = *db
	return uh
}

func (uh *UserHandler) Get(writer http.ResponseWriter, request *http.Request) {
	id, err := utilities.GetIDFromRequest(request)

	if err != nil {
		controllers.RespondBadRequest(writer, err.Error())
		return
	}

	u, err := uh.db.GetUserByID(id)

	if err != nil {
		controllers.RespondNotFound(writer, err.Error())
		return
	}

	controllers.RespondOK(writer, *u)
}

func (uh *UserHandler) GetAll(writer http.ResponseWriter, request *http.Request) {
	uSlice, err := uh.db.GetAllUsers()

	if err != nil {
		controllers.RespondInternalServerError(writer, err.Error())
		return
	}

	controllers.RespondOK(writer, uSlice)
}

func (uh *UserHandler) Create(writer http.ResponseWriter, request *http.Request) {
	u, err := utilities.DecodeJsonFromRequest(request)

	if err != nil {
		controllers.RespondBadRequest(writer, err.Error())
		return
	}

	err = uh.db.CreateUser(u)

	if err != nil {
		controllers.RespondInternalServerError(writer, err.Error())
		return
	}

	controllers.RespondOK(writer, *u)
}

func (uh *UserHandler) Delete(writer http.ResponseWriter, request *http.Request) {
	id, err := utilities.GetIDFromRequest(request)

	if err != nil {
		controllers.RespondBadRequest(writer, err.Error())
		return
	}

	err = uh.db.DeleteUserByID(id)

	if err != nil {
		controllers.RespondBadRequest(writer, err.Error())
		return
	}

	controllers.RespondOK(writer, "User is deleted!")
}

func (uh *UserHandler) Update(writer http.ResponseWriter, request *http.Request) {
	id, err := utilities.GetIDFromRequest(request)

	if err != nil {
		controllers.RespondBadRequest(writer, err.Error())
		return
	}

	_, err = uh.db.GetUserByID(id)

	if err != nil {
		controllers.RespondNotFound(writer, err.Error())
		return
	}

	u, err := utilities.DecodeJsonFromRequest(request)

	if err != nil {
		controllers.RespondBadRequest(writer, err.Error())
		return
	}

	err = uh.db.UpdateUserByID(id, u)
	if err != nil {
		controllers.RespondInternalServerError(writer, err.Error())
		return
	}

	controllers.RespondOK(writer, *u)
}
