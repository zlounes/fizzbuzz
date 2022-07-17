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

//Return a string with  value from 1 to inputData.Limit, separated by comma
//with inputData.string1 if the value is multiple of inputData.Int1
//with inputData.string2 if multiple of inputData.Int2
//with string1, string2 cocatened if lultiple of Int1 and Int2
//with value in other cases
func CalculateFizzBuzz(inputData FizzBuzzInput) string {
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

func calculateIndex(index int, inputData FizzBuzzInput, concatenedStrings string) string {
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
