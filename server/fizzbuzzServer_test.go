package server

import (
	"net/http"
	"strings"
	"testing"

	. "github.com/zlounes/fizzbuzz/config"
	"github.com/zlounes/fizzbuzz/testutil"
)

func TestPost(t *testing.T) {
	server := NewServer(ServerConfig{Port: 7878})
	defer server.Stop()
	_, _, result := sendFizzBuzz(testutil.TestFizzBuzzInput, *server)
	//status and contet-type could be tested here because managed in the http.resonse
	if result != testutil.TestExpectedResult {
		t.Fail()
		t.Logf("Unexpected fizzbuzz result, expecting %v received %v", testutil.TestExpectedResult, result)
	}
}

func TestGet(t *testing.T) {
	server := NewServer(ServerConfig{Port: 7878})
	defer server.Stop()
	sendFizzBuzz(testutil.TestFizzBuzzInput, *server)
	req, _ := http.NewRequest("GET", "/fizzbuzz/stat", nil)
	w := testutil.NewResponseMock()
	server.ServeHTTP(w, req)
	if w.Header().Get("Content-Type") != "application/json; charset=utf-8" {
		t.Logf("Unexpected post content-type output : %v", w.Header().Get("Content-Type"))
		t.Fail()
		return
	}
	bestHint := testutil.DecodeJson(w.GetResult(), t)
	if bestHint == nil {
		return
	}
	if bestHint.NbCalls != 1 {
		t.Fail()
		t.Logf("unexpected nbCalls, expected 1 received :  %d", bestHint.NbCalls)
		return
	}
	entry := bestHint.Entry
	if entry != testutil.TestFizzBuzzInput {
		t.Fail()
		t.Logf("unexpected result, expected %v received :  %v", testutil.TestFizzBuzzInput, entry)
	}
}

func sendFizzBuzz(inputData FizzBuzzInput, server FizzBuzzServer) (http.ResponseWriter, int, string) {
	data := testutil.BuildFormValues(inputData)
	req, _ := http.NewRequest("POST", "/fizzbuzz", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := testutil.NewResponseMock()
	server.ServeHTTP(w, req)
	return w, w.GetStatus(), w.GetResult()
}
