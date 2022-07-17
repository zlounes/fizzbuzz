package calculator

import (
	"testing"

	"github.com/zlounes/fizzbuzz/testutil"
)

type ExpectedResult struct {
}

func TestCalculator(t *testing.T) {
	result := CalculateFizzBuzz(testutil.TestFizzBuzzInput)
	if testutil.TestExpectedResult != result {
		t.Fail()
		t.Logf("Unexpected result : \n%s\nvs\n%s", testutil.TestExpectedResult, result)
	}
}
