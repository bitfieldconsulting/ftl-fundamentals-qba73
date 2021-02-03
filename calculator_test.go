package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
	"time"
)

// GenerateNums it is a helper function
// that produces two float numbers.
func GenerateNums() (float64, float64) {
	// Make generator to produce different
	// numbers each time it runs (make it nondeterministic).
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	a := r.Float64() * float64(r.Intn(10000))
	b := r.Float64() * float64(r.Intn(10000))
	return a, b
}

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

func TestAddRandom(t *testing.T) {
	t.Parallel()

	name := "Add randomly generated numbers"

	for i := 0; i < 100; i++ {
		a, b := GenerateNums()
		want := a + b

		got := calculator.Add(a, b)

		if got != want {
			t.Errorf("%s, Add(%f, %f) = %f, want: %f", name, a, b, got, want)
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

func TestDivide(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name        string
		a           float64
		b           float64
		want        float64
		expectedErr bool
	}{
		{name: "Divide two positive numbers", a: 6, b: 2, want: 3, expectedErr: false},
		{name: "Divide two negative numbers", a: -10, b: -5, want: 2, expectedErr: false},
		{name: "Divide two negative fraction numbers", a: -10.5, b: -5, want: 2.1, expectedErr: false},
		{name: "Divide one positive, one negative number", a: -10, b: 5, want: -2, expectedErr: false},
		{name: "Divide 0 by a number", a: 0, b: 5, want: 0, expectedErr: false},
		{name: "Divide positive number by 0", a: 4, b: 0, want: 0, expectedErr: true},
		{name: "Divide negative number by 0", a: -2, b: 0, want: 0, expectedErr: true},
	}

	for _, tc := range tt {
		got, err := calculator.Divide(tc.a, tc.b)

		if err != nil && !tc.expectedErr {
			t.Errorf("%s, Divide(%f, %f) should return error", tc.name, tc.a, tc.b)
		}

		if got != tc.want && !tc.expectedErr {
			t.Errorf("%s, Divide(%f, %f) = %f, want %f", tc.name, tc.a, tc.b, got, tc.want)
		}
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
		{name: "Sqrt negative number", a: -4, want: 0, expectedErr: true},
		{name: "Sqrt negative number", a: -16.5, want: 0, expectedErr: true},
	}

	for _, tc := range tt {
		got, err := calculator.Sqrt(tc.a)

		if err != nil && !tc.expectedErr {
			t.Errorf("%s, Sqrt(%f) should return error", tc.name, tc.a)
		}

		if got != tc.want && !tc.expectedErr {
			t.Errorf("%s, Sqrt(%f) = %f, want %f", tc.name, tc.a, got, tc.want)
		}
	}
}
