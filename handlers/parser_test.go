package handlers

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/zlounes/fizzbuzz/config"
)

func TestParser(t *testing.T) {
	resultExpected := config.InputData{Int1: 3, Int2: 5, Limit: 9, String1: "fizz", String2: "buzz"}
	data := buildValues(resultExpected)
	req, _ := http.NewRequest("POST", "https://whatever:8080", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	inputData, err := parseForm(req)
	if err != nil {
		t.Fail()
		t.Logf("Could not parse request %v", err)
	} else if resultExpected != *inputData {
		t.Fail()
		t.Logf("Unpexted paring result, expecting %v received %v", resultExpected, inputData)
		return
	}
}

func buildValues(inputData config.InputData) url.Values {
	data := url.Values{}
	data.Add("int1", strconv.Itoa(inputData.Int1))
	data.Add("int2", strconv.Itoa(inputData.Int2))
	data.Add("limit", strconv.Itoa(inputData.Limit))
	data.Add("string1", inputData.String1)
	data.Add("string2", inputData.String2)
	return data
}
