package server

import (
	"fmt"
	"net/http"

	"github.com/zlounes/fizzbuzz/calculator"
	. "github.com/zlounes/fizzbuzz/config"
)

const (
	FORM_CONTENT = `<form action="/fizzbuzz" method="POST" enctype="application/x-www-form-urlencoded">
    <label for="int1">int1:</label><br>
    <input type="text" id="int1" name="int1" value="3"><br>
    <label for="int2">int2:</label><br>
    <input type="text" id="lname" name="int2" value="5"><br>
    <label for="limit">limit:</label><br>
    <input type="text" id="limit" name="limit" value="30"><br>
    <label for="string1">string1:</label><br>
    <input type="text" id="string1" name="string1" value="fizz"><br>
    <label for="string2">string1:</label><br>
    <input type="text" id="string2" name="string2" value="buzz"><br><br>
    <input type="submit" value="Submit">
  </form>`
)

type fizzBuzzHandler struct {
	http.Handler
	channelInputStat chan<- FizzBuzzInput
}

//HAnle the fiizbuzz post call returning text content and GET returing the fom
func (handler *fizzBuzzHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		handler.processFizzBuzzRequest(w, req)
	case "GET":
		setHtmlContent(w, FORM_CONTENT)
	default:
		methodNotAllowed(w, "POST, GET")
	}
}

func (handler *fizzBuzzHandler) processFizzBuzzRequest(w http.ResponseWriter, req *http.Request) {
	var inputData *FizzBuzzInput
	var err error
	if inputData, err = parseForm(req); err != nil {
		returnError(w, http.StatusBadRequest, err)
		return
	}
	result := calculator.CalculateFizzBuzz(*inputData)
	if _, err := w.Write([]byte(result)); err != nil {
		returnError(w, http.StatusInternalServerError, err)
		return
	}
	//send the input data to statistic calculator
	handler.channelInputStat <- *inputData
}

func setHtmlContent(w http.ResponseWriter, text string) {
	w.Header().Set("Content-Type", "text/html")
	if _, err := w.Write([]byte(text)); err != nil {
		returnError(w, http.StatusInternalServerError, err)
		return
	}
}

func returnError(w http.ResponseWriter, statusCode int, err error) {
	returnHttpError(w, statusCode, fmt.Sprintf("Could not process the request : %v", err))
}
