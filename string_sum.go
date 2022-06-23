package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	trimmed := strings.ReplaceAll(input, " ", "")

	if trimmed == "" {
		return "", errorEmptyInput
	}
	var indexOfOperands []int
	for i := 0; i < len(trimmed); i++ {
		char := string(trimmed[i])

		if char == "+" || char == "-" {
			indexOfOperands = append(indexOfOperands, i)
			continue
		}

		if _, err := strconv.Atoi(char); err != nil && char != " " {
			return "", fmt.Errorf("invalid input %w", err)
		}
	}

	if len(indexOfOperands) > 2 {
		return "", errorNotTwoOperands
	}

	firstOperand, err := strconv.Atoi(trimmed[0:indexOfOperands[1]])
	secondOperand, err := strconv.Atoi(trimmed[indexOfOperands[1]:])

	result := firstOperand + secondOperand

	return strconv.Itoa(result), nil
}
