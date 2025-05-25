package test

import "testing"

func TestAddTableDriven(t *testing.T) {
	// Define the "table" of test cases
	testCases := []struct {
		name string // A name for the test case
		a    int
		b    int
		want int
	}{
		{name: "Positive numbers", a: 2, b: 3, want: 5},
		{name: "Negative numbers", a: -5, b: -10, want: -15},
		{name: "Adding zero", a: 100, b: 0, want: 100},
	}

	// Loop through the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Add(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
