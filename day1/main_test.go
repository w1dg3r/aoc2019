package main

import (
	"fmt"
	"testing"
)

func TestFuel(t *testing.T) {
	cases := []struct {
		mass     int
		expected int
	}{
		{mass: 12, expected: 2},
		{mass: 14, expected: 2},
		{mass: 1969, expected: 654},
		{mass: 100756, expected: 33583},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Mass_%d", c.mass), func(t *testing.T) {
			actual := calculateFuel(c.mass)

			if c.expected != actual {
				t.Errorf("Wanted: %v got: %v", c.expected, actual)
			}
		})
	}
}
