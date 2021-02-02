package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name string
		a    float64
		b    float64
		want float64
	}{
		{name: "Add two positive numbers", a: 4, b: 5, want: 9},
		{name: "Add two negative numbers", a: -4.5, b: -5.5, want: -10},
		{name: "Add zero to positive number", a: 3, b: 0, want: 3},
	}

	for _, tc := range tt {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s, want %f, got %f", tc.name, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name string
		a    float64
		b    float64
		want float64
	}{
		{name: "Substract two positive numbers", a: 4, b: 5, want: -1},
		{name: "Substract two negative numbers", a: -4.5, b: -5.5, want: 1},
		{name: "Substract from zero", a: 0, b: 5.45, want: -5.45},
	}

	for _, tc := range tt {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s, want %f, got %f", tc.name, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name string
		a    float64
		b    float64
		want float64
	}{
		{name: "Multiply two positive numbers", a: 3, b: 2, want: 6},
		{name: "Multiply two negative numbers", a: -4, b: -5, want: 20},
		{name: "Multiply one negative one positive", a: -4, b: 5, want: -20},
		{name: "Multiply by zero", a: -4, b: 0, want: 0},
	}

	for _, tc := range tt {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s, want %f, got %f", tc.name, tc.want, got)
		}
	}
}
