//go:build integration
// +build integration

package main

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/zlounes/fizzbuzz/config"
	"github.com/zlounes/fizzbuzz/testutil"
)

var (
	inputData = config.InputData{
		Int1:    3,
		Int2:    5,
		Limit:   30,
		String1: "fizz",
		String2: "buzz",
	}
)

func TestServer(t *testing.T) {
	var resp *http.Response
	var err error
	var result string
	expectedResult := "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,17,fizz,19,buzz,fizz,22,23,fizz,buzz,26,fizz,28,29,fizzbuzz"
	data := testutil.BuildFormValues(inputData)
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	url := fmt.Sprintf("http://localhost:%s/fizzbuzz", port)
	if resp, err = http.PostForm(url, data); err != nil {
		t.Logf("Erreur calling %s : %v", url, err)
		t.Fail()
		return
	}
	if resp.StatusCode != 200 {
		t.Logf("Unexpected status : %d", resp.StatusCode)
		t.Fail()
		return
	}
	if result, err = testutil.ReadResult(*resp); err != nil {
		t.Logf("Erreur retreiving value : %v", err)
		t.Fail()
		return
	}
	if result != expectedResult {
		t.Logf("Erreur retreiving value : %v", err)
		t.Fail()
	}
}
