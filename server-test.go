package main

import (
	"testing"
)

func TestServer(t *testing.T) {
	if 3 == 4 {
		t.Fail()
		t.Logf("Could not parse request")
	}
}
