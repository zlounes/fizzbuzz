package server

import (
	"fmt"
	"net/http"
	"strconv"

	. "github.com/zlounes/fizzbuzz/config"
)

const (
	MAX_STRING_LEN = 50
)

func parseForm(req *http.Request) (*FizzBuzzInput, error) {
	var intVals []int
	var strVals []string
	var err error
	if err = req.ParseForm(); err != nil {
		return nil, fmt.Errorf("ParseForm() err: %v", err)
	}
	if intVals, err = parseInts(req, "int1", "int2", "limit"); err != nil {
		return nil, err
	}
	if strVals, err = parseStrings(req, "string1", "string2"); err != nil {
		return nil, err
	}
	return &FizzBuzzInput{Int1: intVals[0], Int2: intVals[1], Limit: intVals[2],
		String1: strVals[0], String2: strVals[1]}, nil
}

func parseInts(req *http.Request, strKeys ...string) ([]int, error) {
	result := make([]int, 0)
	var strVal string
	var err error
	var val int
	for _, strKey := range strKeys {
		if strVal, err = parseRequiredArg(req, strKey); err != nil {
			return nil, err
		}
		if val, err = strconv.Atoi(strVal); err != nil {
			return nil, fmt.Errorf("The post argument %s should be an integer", strKey)
		}
		if val <= 0 {
			return nil, fmt.Errorf("The post argument %s should be gretaer than 0", strKey)
		}
		result = append(result, val)
	}
	return result, nil
}

func parseStrings(req *http.Request, strKeys ...string) ([]string, error) {
	result := make([]string, 0)
	var strVal string
	var err error
	for _, strKey := range strKeys {
		if strVal, err = parseRequiredArg(req, strKey); err != nil {
			return nil, err
		}
		if len(strVal) > MAX_STRING_LEN {
			return nil, fmt.Errorf("The lenght for %s should lower than %d", strKey, MAX_STRING_LEN)
		}
		result = append(result, strVal)
	}
	return result, nil
}

func parseRequiredArg(req *http.Request, key string) (string, error) {
	strVal := req.FormValue(key)
	if len(strVal) == 0 {
		return "", fmt.Errorf("The post argument %s is required", key)
	}
	return strVal, nil
}
