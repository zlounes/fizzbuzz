package handlers

import (
	"net/http"
	"testing"

	. "github.com/zlounes/fizzbuzz/config"
)

type requestWriterMock struct {
	http.ResponseWriter
	result string
}

func newMock() *requestWriterMock {
	return &requestWriterMock{result: ""}
}

func (mock *requestWriterMock) Write(value []byte) (int, error) {
	mock.result = mock.result + string(value)
	return len(value), nil
}

type ExpectedResult struct {
}

func TestCalculator(t *testing.T) {
	mock := newMock()
	inputData := InputData{
		Int1:    3,
		Int2:    5,
		Limit:   30,
		String1: "fizz",
		String2: "buzz",
	}
	calculateFizzBuzz(mock, inputData)
	expectedResult := "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,17,fizz,19,buzz,fizz,22,23,fizz,buzz,26,fizz,28,29,fizzbuzz"
	if expectedResult != mock.result {
		t.Fail()
		t.Logf("Unexpected result : \n%s\nvs\n%s", expectedResult, mock.result)
	}
}
