package controllers

import (
	"encoding/json"
	"net/http"
)

func RespondBadRequest(writer http.ResponseWriter, message string) {
	responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": message})
}

func RespondInternalServerError(writer http.ResponseWriter, message string) {
	responseWithJson(writer, http.StatusInternalServerError, map[string]string{"message": message})
}

func RespondNotFound(writer http.ResponseWriter, message string) {
	responseWithJson(writer, http.StatusNotFound, map[string]string{"message": message})
}

func RespondOK(writer http.ResponseWriter, object any) {
	switch object := object.(type) {
	default:
		responseWithJson(writer, http.StatusOK, object)
	case string:
		responseWithJson(writer, http.StatusOK, map[string]string{"message": object})
	}

}

func responseWithJson(writer http.ResponseWriter, status int, object any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(object)
}
