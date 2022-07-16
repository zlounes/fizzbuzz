package server

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	. "github.com/zlounes/fizzbuzz/config"
	"github.com/zlounes/fizzbuzz/metrics"
	"github.com/zlounes/fizzbuzz/testutil"
)

var (
	inputData = InputData{
		Int1:    3,
		Int2:    5,
		Limit:   30,
		String1: "fizz",
		String2: "buzz",
	}
)

func TestPost(t *testing.T) {
	expectedResult := "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,17,fizz,19,buzz,fizz,22,23,fizz,buzz,26,fizz,28,29,fizzbuzz"
	server := NewServer(ServerConfig{Port: 7878})
	defer server.Stop()
	w, _, result := sendFizzBuzz(inputData, *server)
	//status could be tested here because managed in the http.resonse
	if w.Header().Get("Content-Type") != "text/plain; charset=utf-8" {
		t.Logf("Unexpected post content-type output : %v", w.Header().Get("Content-Type"))
		t.Fail()
	}
	if result != expectedResult {
		t.Fail()
		t.Logf("Unexpected fizzbuzz result, expecting %v received %v", expectedResult, result)
	}
}

func TestGet(t *testing.T) {
	server := NewServer(ServerConfig{Port: 7878})
	defer server.Stop()
	sendFizzBuzz(inputData, *server)
	req, _ := http.NewRequest("GET", "/fizzbuzz/stat", nil)
	w := testutil.NewResponseMock()
	server.ServeHTTP(w, req)
	if w.Header().Get("Content-Type") != "application/json; charset=utf-8" {
		t.Logf("Unexpected post content-type output : %v", w.Header().Get("Content-Type"))
		t.Fail()
		return
	}
	bestHint := metrics.BestHint{}
	decoder := json.NewDecoder(strings.NewReader(w.GetResult()))
	if err := decoder.Decode(&bestHint); err != nil {
		t.Fail()
		t.Logf("Could not decode the result :  %v", err)
		return
	}
	if bestHint.NbCalls != 1 {
		t.Fail()
		t.Logf("unexpected nbCalls, expected 1 received :  %d", bestHint.NbCalls)
		return
	}
	entry := bestHint.Entry
	if entry != inputData {
		t.Fail()
		t.Logf("unexpected result, expected %v received :  %v", inputData, entry)
	}
}

func sendFizzBuzz(inputData InputData, server FizzBuzzServer) (http.ResponseWriter, int, string) {
	data := testutil.BuildFormValues(inputData)
	req, _ := http.NewRequest("POST", "/fizzbuzz", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := testutil.NewResponseMock()
	server.ServeHTTP(w, req)
	return w, w.GetStatus(), w.GetResult()
}
