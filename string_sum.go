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
	const errorMsg = "invalid input"

	trimmed := strings.ReplaceAll(input, " ", "")

	if trimmed == "" {
		return "", fmt.Errorf(errorMsg+" %w", errorEmptyInput)
	}

	if _, err := strconv.Atoi(string(trimmed[0])); err == nil {
		trimmed = "+" + trimmed
	}
	fmt.Println(trimmed)

	var indexOfOperands []int
	for i := 0; i < len(trimmed); i++ {
		char := string(trimmed[i])

		if char == "+" || char == "-" {
			indexOfOperands = append(indexOfOperands, i)
			continue
		}

		if _, err := strconv.Atoi(char); err != nil && char != " " {
			return "", fmt.Errorf(errorMsg+" %w", err)
		}
	}

	numberOfOperands := len(indexOfOperands)

	if numberOfOperands != 2 {
		return "", fmt.Errorf(errorMsg+" %w", errorNotTwoOperands)
	}

	point := indexOfOperands[1]

	firstOperand, err := strconv.Atoi(trimmed[0:point])
	secondOperand, err := strconv.Atoi(trimmed[point:])

	result := firstOperand + secondOperand

	return strconv.Itoa(result), nil
}
