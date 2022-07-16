package calculator

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/zlounes/fizzbuzz/config"
)

var (
	SEPARATOR = ","
)

func CalculateFizzBuzz(inputData InputData) string {
	var sb strings.Builder
	var concatenedStrings = fmt.Sprintf("%s%s", inputData.String1, inputData.String2)
	for i := 1; i <= inputData.Limit; i++ {
		if i > 1 {
			sb.WriteString(SEPARATOR)
		}
		sb.WriteString(calculateIndex(i, inputData, concatenedStrings))
	}
	return sb.String()
}

func calculateIndex(index int, inputData InputData, concatenedStrings string) string {
	var result string
	fizz := (index % inputData.Int1) == 0
	buzz := (index % inputData.Int2) == 0
	if fizz && buzz {
		result = concatenedStrings
	} else if fizz {
		result = inputData.String1
	} else if buzz {
		result = inputData.String2
	} else {
		result = strconv.Itoa(index)
	}
	return result
}
