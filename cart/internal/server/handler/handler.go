package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"
)

type ErrResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func WriteResponse(w http.ResponseWriter, v any) {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("Failed response encode: ", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func WriteResponseError(w http.ResponseWriter, code int, msg string) {
	response := ErrResponse{Code: code, Message: msg}
	w.WriteHeader(http.StatusUnprocessableEntity)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed write error response: ", err)
	}
}

func WriteInternalError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	log.Printf("Internal error: %s (%s)", err.Error(), debug.Stack())
}
