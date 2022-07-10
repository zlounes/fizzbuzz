package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	. "github.com/zlounes/fizzbuzz/config"
)

var (
	SEPARATOR = []byte(",")
)

func calculateFizzBuzz(w http.ResponseWriter, inputData InputData) {
	var value string
	var fizzBuzz = fmt.Sprintf("%s%s", inputData.String1, inputData.String2)
	for i := 1; i <= inputData.Limit; i++ {
		if i > 1 {
			w.Write(SEPARATOR)
		}
		fizz := (i % inputData.Int1) == 0
		buzz := (i % inputData.Int2) == 0
		if fizz && buzz {
			value = fizzBuzz
		} else if fizz {
			value = inputData.String1
		} else if buzz {
			value = inputData.String2
		} else {
			value = strconv.Itoa(i)
		}
		w.Write([]byte(value))
	}
}
