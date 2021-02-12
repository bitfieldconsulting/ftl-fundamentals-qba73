// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
)

// Add takes two or more numbers and returns the result of adding them together.
func Add(a, b float64, c ...float64) float64 {
	result := a + b

	for _, i := range c {
		result += i
	}

	return result
}

// Subtract takes two or more numbers and returns the result
// of subtracting them.
func Subtract(a, b float64, c ...float64) float64 {
	result := a - b

	for _, i := range c {
		result -= i
	}

	return result
}

// Multiply takes two or more numbers and returns the product.
func Multiply(a, b float64, c ...float64) float64 {
	result := a * b

	for _, i := range c {
		if i == 0 {
			return 0
		}
		result *= i
	}

	return result
}

// Divide takes two or more numbers and return the result of division.
// It returns division by zero error if the n+1 argument is equal 0.
func Divide(a, b float64, c ...float64) (float64, error) {
	result := a / b

	for _, i := range c {
		result /= i
	}

	if math.IsInf(result, 0) {
		return 0, fmt.Errorf("invalid input Divide(%f, %f, %v), division by zero is undefined", a, b, c)
	}

	return result, nil
}

// Sqrt takes a number and returns square root of that number.
// If the argument value is < 0 the function returns an error.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("invalid input Sqrt(%f)", a)
	}
	return math.Sqrt(a), nil
}

// Compute takes a string representing numbers and operation (+, -, /, *)
// on them and returns result of the operation. It returns an error
// if the computation sign is not recognized or attempted division by zero.
//
// Examples:
//
// Compute("3 - 2") should return 1
// Compute("2 & 2") should return error
func Compute(expression string) (float64, error) {
	var a, b float64
	var operator string
	var err error

	n, err := fmt.Sscanf(expression, "%f %s %f", &a, &operator, &b)
	if err != nil {
		return 0, err
	}

	if n != 3 {
		return 0, fmt.Errorf("invalid number of parsed text fields %v", n)
	}

	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("invalid opertor used %v", operator)
	}
}
