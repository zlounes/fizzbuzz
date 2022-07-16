package testutil

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/zlounes/fizzbuzz/config"
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
	fmt.Println("Hello")
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

func BuildFormValues(inputData config.InputData) url.Values {
	data := url.Values{}
	data.Add("int1", strconv.Itoa(inputData.Int1))
	data.Add("int2", strconv.Itoa(inputData.Int2))
	data.Add("limit", strconv.Itoa(inputData.Limit))
	data.Add("string1", inputData.String1)
	data.Add("string2", inputData.String2)
	return data
}
