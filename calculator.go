// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
)

// Add takes 0 or more numbers and returns the result of adding them together.
func Add(a ...float64) float64 {
	var sum float64

	for _, i := range a {
		sum = sum + i
	}

	return sum
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers and returns the product.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers and return the result of division.
// It returns error if the second number is 0 (division by zero).
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("error: division by zero")
	}
	return a / b, nil
}

// Sqrt takes a number and returns square root if the number
// is >= 0. If number is < 0 the function returns an error.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("error: sqrt from negative number")
	}
	return math.Sqrt(a), nil
}
