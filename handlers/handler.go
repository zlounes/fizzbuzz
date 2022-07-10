package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	. "github.com/zlounes/fizzbuzz/config"
)

type FizzBuzzHandler struct {
}

func (handler *FizzBuzzHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		post(w, req)
	default:
		returnHttpError(w, http.StatusMethodNotAllowed, "Sorry, only POST and GET methods are supported.")
	}
}

func post(w http.ResponseWriter, req *http.Request) {
	var inputData *InputData
	var err error
	if inputData, err = parseForm(req); err != nil {
		returnHttpError(w, http.StatusBadRequest, fmt.Sprintf("Could not process the request : %v", err))
		return
	}
	calculateFizzBuzz(w, *inputData)
}

func returnHttpError(w http.ResponseWriter, statusCode int, reasonStr string) {
	reason := []byte(reasonStr)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(reason)))
	w.WriteHeader(statusCode)
	w.Write(reason)
}

func returnHttpSuccess(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	//w.Header().Set("Content-Length", strconv.Itoa(len(reason)))
	w.WriteHeader(http.StatusAccepted)
}
