//go:build integration
// +build integration

package main

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/zlounes/fizzbuzz/testutil"
)

func TestServer(t *testing.T) {
	var resp *http.Response
	var err error
	var result string
	data := testutil.BuildFormValues(testutil.TestFizzBuzzInput)

	url := getPostUrl()
	if resp, err = http.PostForm(url, data); err != nil {
		t.Logf("Erreur calling %s : %v", url, err)
		t.Fail()
		return
	}
	if !checkStatus(resp, t) {
		return
	}
	if result, err = testutil.ReadResult(*resp); err != nil {
		t.Logf("Erreur retreiving value : %v", err)
		t.Fail()
		return
	}
	if result != testutil.TestExpectedResult {
		t.Logf("Erreur retreiving value : %v", err)
		t.Fail()
	}

	resp, err = http.Get(getStatUrl())
	if checkStatus(resp, t) {
		return
	}
	result, _ = testutil.ReadResult(*resp)
	bestHint := testutil.DecodeJson(result, t)
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

func getPostUrl() string {
	return fmt.Sprintf("http://localhost:%s/fizzbuzz", getPort())
}

func getStatUrl() string {
	return fmt.Sprintf("http://localhost:%s/fizzbuzz/stat", getPort())
}

func getPort() string {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func checkStatus(resp *http.Response, t *testing.T) bool {
	if resp.StatusCode != http.StatusOK {
		t.Logf("Unexpected status : %d", resp.StatusCode)
		t.Fail()
		return false
	}
	return true
}
