package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenerateNums it is a helper function that produces a slice
// of minimum two and maximum 100 float numbers.
func GenerateNums() []float64 {
	// Max & Min represent maximum and minimum numbers to generate.
	min, max := 2, 100
	// Make sure we have at least 0 + min numbers!
	n := random.Intn(max-min) + min

	nums := make([]float64, n)

	// Populate slice with float numbers.
	for i := 0; i < n; i++ {
		fl := random.Float64() * float64(random.Intn(10000))
		nums[i] = fl
	}

	return nums
}

func TestAddRandom(t *testing.T) {
	t.Parallel()

	name := "Add randomly generated numbers"

	for i := 0; i < 100; i++ {
		a := GenerateNums()
		var want float64
		var args []float64

		for _, i := range a {
			want += i
		}

		// If we generated more than required minimum params
		// then we will pass them to Add() func as optional args.
		if len(a) > 2 {
			args = a[2:]
		}

		got := calculator.Add(a[0], a[1], args...)

		if got != want {
			t.Errorf("%s; Add() = %f; want: %f", name, got, want)
		}
	}
}

func TestCalculator(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name string
		fn   func(a, b float64, c ...float64) float64
		a    float64
		b    float64
		c    []float64
		want float64
	}{
		{name: "Add two positive numbers", fn: calculator.Add, a: 1, b: 2, want: 3},
		{name: "Add two negative numbers", fn: calculator.Add, a: -4.5, b: -5.5, want: -10},
		{name: "Add three negative numbers", fn: calculator.Add, a: -4.5, b: -5.5, c: []float64{-4.5}, want: -14.5},
		{name: "Add multiple numbers", fn: calculator.Add, a: 2, b: 3.4, c: []float64{-4.5, -5.5, 10, 2.5}, want: 7.9},

		{name: "Multiply three positive", fn: calculator.Multiply, a: 3, b: 2, c: []float64{2}, want: 12},
		{name: "Multiply two positive one negative", fn: calculator.Multiply, a: 3, b: 2, c: []float64{-2}, want: -12},
		{name: "Multiply one positive two negative", fn: calculator.Multiply, a: -3, b: 2, c: []float64{-2.5}, want: 15},
		{name: "Multiply multiple numbers", fn: calculator.Multiply, a: 3, b: 2, c: []float64{2, -1}, want: -12},
		{name: "Multiply including negative numbers", fn: calculator.Multiply, a: 2.5, b: 1.5, c: []float64{-4, -5}, want: 75},
		{name: "Multiply by zero", fn: calculator.Multiply, a: 3, b: 4, c: []float64{0, -4}, want: 0},

		{name: "Subtract two positive numbers", fn: calculator.Subtract, a: 4, b: 2, want: 2},
		{name: "Subtract two negative numbers", fn: calculator.Subtract, a: -4.5, b: -5.5, want: 1},
		{name: "Subtract from zero", fn: calculator.Subtract, a: 0, b: 5.45, want: -5.45},
		{name: "Subtract multiple numbers", fn: calculator.Subtract, a: 10, b: 3, c: []float64{2, 1, 4, 3}, want: -3},
		{name: "Subtract multiple numbers with negative", fn: calculator.Subtract, a: 10.5, b: 3, c: []float64{20, 5.5, 4, -3, 1.25}, want: -20.25},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.fn(tc.a, tc.b, tc.c...)
			if got != tc.want {
				t.Errorf("%s, got: %v, want: %v", tc.name, got, tc.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name        string
		a, b        float64
		c           []float64
		want        float64
		expectedErr bool
	}{
		{name: "Divide two positive numbers", a: 6, b: 2, c: []float64{}, want: 3, expectedErr: false},
		{name: "Divide two negative numbers", a: -10, b: -5, want: 2, expectedErr: false},
		{name: "Divide two negative fraction numbers", a: -10.5, b: -5, want: 2.1, expectedErr: false},
		{name: "Divide one positive, one negative number", a: -10, b: 2.5, want: -4, expectedErr: false},
		{name: "Divide 0 by a number", a: 0, b: 5, want: 0, expectedErr: false},
		{name: "Divide by zero", a: 2.5, b: 0.5, c: []float64{2, 1.25, 1}, want: 2, expectedErr: false},

		// Test cases with expected errors.
		{name: "Divide by zero", a: 2, b: 0, want: 0, expectedErr: true},
		{name: "Divide by zero", a: 2, b: 1, c: []float64{2, 0, 1}, want: 0, expectedErr: true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			got, err := calculator.Divide(tc.a, tc.b, tc.c...)

			if (err != nil) != tc.expectedErr {
				t.Fatalf("%s; Divide(%f, %f, %v) unexpected error status: %v", tc.name, tc.a, tc.b, tc.c, err)
			}

			if !tc.expectedErr && got != tc.want {
				t.Errorf("%s; Divide(%f, %f, %v) = %f; want %f", tc.name, tc.a, tc.b, tc.c, got, tc.want)
			}
		})
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name        string
		a           float64
		want        float64
		expectedErr bool
	}{
		{name: "Sqrt positive number", a: 9, want: 3, expectedErr: false},
		{name: "Sqrt 0 number", a: 0, want: 0, expectedErr: false},

		// Test cases with expected error.
		{name: "Sqrt negative number", a: -4, want: 0, expectedErr: true},
		{name: "Sqrt negative number", a: -16.5, want: 0, expectedErr: true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := calculator.Sqrt(tc.a)

			if (err != nil) != tc.expectedErr {
				t.Errorf("%s; Sqrt(%f) should return error", tc.name, tc.a)
			}

			if !tc.expectedErr && got != tc.want {
				t.Errorf("%s; Sqrt(%f) = %f; want %f", tc.name, tc.a, got, tc.want)
			}
		})
	}
}

func TestCompute(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name        string
		expression  string
		want        float64
		expectedErr bool
	}{
		{name: "Compute sum", expression: "2 + 3", want: 5, expectedErr: false},
		{name: "Compute sum, no float", expression: "2 + 4", want: 6, expectedErr: false},
		{name: "Compute substract no float", expression: "5 - 2", want: 3, expectedErr: false},
		{name: "Compute multiply no float", expression: "4 * -4", want: -16, expectedErr: false},
		{name: "Compute multiply no float", expression: "3 * -2", want: -6, expectedErr: false},
		{name: "Compute multiply no float", expression: "-3 * 2", want: -6, expectedErr: false},
		{name: "Compute divide no float", expression: "10 / 2", want: 5, expectedErr: false},
		{name: "Compute additional spaces", expression: "2  + 4 ", want: 6, expectedErr: false},
		{name: "Compute additional spaces", expression: " 2  +  4  ", want: 6, expectedErr: false},
		{name: "Compute additional spaces", expression: "  2    +  -4  ", want: -2, expectedErr: false},

		// Test cases with expected error.
		{name: "Compute incorrect operator", expression: " 10 & 2   ", want: 0, expectedErr: true},
		{name: "Compute incorrect operator", expression: " 10  #  2", want: 0, expectedErr: true},
		{name: "Compute incorrect operator", expression: "10 & 2", want: 0, expectedErr: true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := calculator.Compute(tc.expression)

			if (err != nil) != tc.expectedErr {
				t.Errorf("%s; Compute(%s) = %v; expected error", tc.name, tc.expression, err)
			}

			if !tc.expectedErr && got != tc.want {
				t.Errorf("%s; Compute(%s) = %f; want %f", tc.name, tc.expression, got, tc.want)
			}
		})
	}
}
