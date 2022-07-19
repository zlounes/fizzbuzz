package testutil

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/zlounes/fizzbuzz/config"
	"github.com/zlounes/fizzbuzz/metrics"
)

const (
	TestExpectedResult = "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,17,fizz,19,buzz,fizz,22,23,fizz,buzz,26,fizz,28,29,fizzbuzz"
)

var (
	TestFizzBuzzInput = config.FizzBuzzInput{
		Int1:    3,
		Int2:    5,
		Limit:   30,
		String1: "fizz",
		String2: "buzz",
	}
	TestWrongInput = config.FizzBuzzInput{
		Int1:    3,
		Int2:    -5,
		Limit:   30,
		String1: "fizz",
		String2: "buzz",
	}
)

type requestWriterMock struct {
	http.ResponseWriter
	result     string
	header     http.Header
	statusCode int
}

func NewResponseMock() *requestWriterMock {
	return &requestWriterMock{result: "", header: http.Header{}}
}

func (mock *requestWriterMock) Write(value []byte) (int, error) {
	mock.result = mock.result + string(value)
	return len(value), nil
}

func (mock *requestWriterMock) Header() http.Header {
	return mock.header
}

func (mock *requestWriterMock) WriteHeader(statusCode int) {
	mock.statusCode = statusCode
}

func (mock *requestWriterMock) GetResult() string {
	return mock.result
}

func (mock *requestWriterMock) GetStatus() int {
	return mock.statusCode
}

func ReadResult(resp http.Response) (string, error) {
	b, err := io.ReadAll(resp.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func BuildFormValues(inputData config.FizzBuzzInput) url.Values {
	data := url.Values{}
	data.Add("int1", strconv.Itoa(inputData.Int1))
	data.Add("int2", strconv.Itoa(inputData.Int2))
	data.Add("limit", strconv.Itoa(inputData.Limit))
	data.Add("string1", inputData.String1)
	data.Add("string2", inputData.String2)
	return data
}

func DecodeJson(str string, t *testing.T) *metrics.BestHint {
	bestHint := metrics.BestHint{}
	decoder := json.NewDecoder(strings.NewReader(str))
	if err := decoder.Decode(&bestHint); err != nil {
		t.Fail()
		t.Logf("Could not decode the result :  %v", err)
		return nil
	}
	return &bestHint
}
