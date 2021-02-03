package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
	"time"
)

// GenerateNums it is a helper function
// that produces a slice of two float numbers.
func GenerateNums() []float64 {
	// Make generator to produce different
	// numbers each time it runs (make it nondeterministic).
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	a := r.Float64() * float64(r.Intn(10000))
	b := r.Float64() * float64(r.Intn(10000))
	return []float64{a, b}
}

func TestAdd(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name        string
		nums        []float64
		want        float64
		expectedErr bool
	}{
		{name: "Add two positive numbers", nums: []float64{1, 2}, want: 3, expectedErr: false},
		{name: "Add two negative numbers", nums: []float64{-4.5, -5.5}, want: -10, expectedErr: false},
		{name: "Add multiple numbers", nums: []float64{-4.5, -5.5, 10, 2.5}, want: 2.5, expectedErr: false},
		{name: "Add no numbers", nums: []float64{}, want: 0, expectedErr: true},
		{name: "Add one number", nums: []float64{10}, want: 10, expectedErr: true},
	}

	for _, tc := range tt {
		got, err := calculator.Add(tc.nums...)

		if err != nil && !tc.expectedErr {
			t.Errorf("%s Add(%v) should return an error", tc.name, tc.nums)
		}

		if tc.want != got && !tc.expectedErr {
			t.Errorf("%s Add(%v) = %f, want %f", tc.name, tc.nums, got, tc.want)
		}
	}
}

func TestAddRandom(t *testing.T) {
	t.Parallel()

	name := "Add randomly generated numbers"

	for i := 0; i < 100; i++ {
		a := GenerateNums()
		want := a[0] + a[1]

		// Do not check err here as in this test
		// we always pass two params to the function.
		got, _ := calculator.Add(a...)

		if got != want {
			t.Errorf("%s, Add(%v) = %f, want: %f", name, a, got, want)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name        string
		nums        []float64
		want        float64
		expectedErr bool
	}{
		{name: "Substract two positive numbers", nums: []float64{4, 2}, want: 2, expectedErr: false},
		{name: "Substract two negative numbers", nums: []float64{-4.5, -5.5}, want: 1, expectedErr: false},
		{name: "Substract from zero", nums: []float64{0, 5.45}, want: -5.45, expectedErr: false},
		{name: "Substract multiple numbers", nums: []float64{20, 5.5, 4, 3, 1.25}, want: 6.25, expectedErr: false},
		{name: "Substract incorrect number of params", nums: []float64{4}, want: 0, expectedErr: true},
		{name: "Substract incorrect number of params", nums: []float64{}, want: 0, expectedErr: true},
	}

	for _, tc := range tt {
		got, err := calculator.Subtract(tc.nums...)

		if err != nil && !tc.expectedErr {
			t.Errorf("%s, Substract(%v) should return an error", tc.name, tc.nums)
		}

		if tc.want != got && !tc.expectedErr {
			t.Errorf("%s, Substract(%v) = %f, want %f", tc.name, tc.nums, got, tc.want)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name        string
		nums        []float64
		want        float64
		expectedErr bool
	}{
		{name: "Multiply three positive", nums: []float64{3, 2, 2}, want: 12, expectedErr: false},
		{name: "Multiply two positive one negative", nums: []float64{3, 2, -2}, want: -12, expectedErr: false},
		{name: "Multiply one positive two negative", nums: []float64{-3, 2, -2}, want: 12, expectedErr: false},
		{name: "Multiply two positive numbers", nums: []float64{3, 2}, want: 6, expectedErr: false},
		{name: "Multiply two negative numbers", nums: []float64{-4, -5}, want: 20, expectedErr: false},
		{name: "Multiply one negative one positive", nums: []float64{-4, 5}, want: -20, expectedErr: false},
		{name: "Multiply by zero", nums: []float64{-4, 0}, want: 0, expectedErr: false},
		{name: "Multiply zero args", nums: []float64{}, want: 0, expectedErr: true},
		{name: "Multiply one arg", nums: []float64{-4}, want: 0, expectedErr: true},
	}

	for _, tc := range tt {
		got, err := calculator.Multiply(tc.nums...)

		if err != nil && !tc.expectedErr {
			t.Errorf("%s, Multiply(%v) should return an error", tc.name, tc.nums)
		}

		if tc.want != got {
			t.Errorf("%s, want %f, got %f", tc.name, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name        string
		nums        []float64
		want        float64
		expectedErr bool
	}{
		{name: "Divide two positive numbers", nums: []float64{6, 2}, want: 3, expectedErr: false},
		{name: "Divide two negative numbers", nums: []float64{-10, -5}, want: 2, expectedErr: false},
		{name: "Divide two negative fraction numbers", nums: []float64{-10.5, -5}, want: 2.1, expectedErr: false},
		{name: "Divide one positive, one negative number", nums: []float64{-10, 5}, want: -2, expectedErr: false},
		{name: "Divide 0 by a number", nums: []float64{0, 5}, want: 0, expectedErr: false},
		{name: "Divide only one argument", nums: []float64{4}, want: 0, expectedErr: true},
		{name: "Divide zero agruments", nums: []float64{}, want: 0, expectedErr: true},
		{name: "Divide positive number by 0", nums: []float64{4, 0}, want: 0, expectedErr: true},
		{name: "Divide positive number by 0", nums: []float64{4, 3, 0}, want: 0, expectedErr: true},
		{name: "Divide positive number by 0", nums: []float64{4, 3, 5, 0}, want: 0, expectedErr: true},
		{name: "Divide positive number by 0", nums: []float64{4, 3, 0, 1}, want: 0, expectedErr: true},
		{name: "Divide negative number by 0", nums: []float64{-2, 0}, want: 0, expectedErr: true},
	}

	for _, tc := range tt {
		got, err := calculator.Divide(tc.nums...)

		if err != nil && !tc.expectedErr {
			t.Errorf("%s, Divide(%v) should return error", tc.name, tc.nums)
		}

		if got != tc.want && !tc.expectedErr {
			t.Errorf("%s, Divide(%v) = %f, want %f", tc.name, tc.nums, got, tc.want)
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
