// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
)

var (
	// errNotEnoughNumbers is returned by functions that expects
	// minimum of two arguments.
	errNotEnoughNumbers = errors.New("minimum two numbers are required")

	errDivisionByZero     = errors.New("division by zero")
	errSqrtNumberNegative = errors.New("sqrt from negative number")
)

// Add takes 2 or more numbers and returns the result of adding them together.
// If the number of arguments is less than two the function will return an error.
func Add(a ...float64) (float64, error) {
	if len(a) < 2 {
		return 0, errNotEnoughNumbers
	}
	var result float64

	for _, i := range a {
		result = result + i
	}

	return result, nil
}

// Subtract takes two or more numbers and returns the result
// of subtracting them or error, if the number of arguments
// is less than two.
func Subtract(a ...float64) (float64, error) {
	if len(a) < 2 {
		return 0, errNotEnoughNumbers
	}
	result := a[0]

	for i := 1; i < len(a); i++ {
		result = result - a[i]
	}

	return result, nil
}

// Multiply takes min two numbers and returns the product
// or error, if the number of arguments is less than two.
func Multiply(a ...float64) (float64, error) {
	if len(a) < 2 {
		return 0, errNotEnoughNumbers
	}
	result := a[0]

	for i := 1; i < len(a); i++ {
		result = result * a[i]
	}

	return result, nil
}

// Divide takes n numbers and return the result of division.
// It returns error if the n+1 argument is equal 0 (division by zero)
// or if the numer of arguments is less than 2.
func Divide(a ...float64) (float64, error) {
	if len(a) < 2 {
		return 0, errNotEnoughNumbers
	}

	result := a[0]

	for i := 1; i < len(a); i++ {
		if a[i] == 0 {
			return 0, errDivisionByZero
		}
		result = result / a[i]
	}

	return result, nil
}

// Sqrt takes a number and returns square root if the number.
// If the number is < 0 the function returns an error.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errSqrtNumberNegative
	}
	return math.Sqrt(a), nil
}
