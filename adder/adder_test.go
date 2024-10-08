package adder

import "testing"

func TestAdder(t *testing.T) {
	t.Run("add two numbers", func(t *testing.T) {
		sum := Add(2, 3)
		expected := 5
		assertCorrectSum(t, sum, expected)
	})

	t.Run("add five numbers", func(t *testing.T) {
		sum := Add(1, 4, 6, 7, 5)
		expected := 23
		assertCorrectSum(t, sum, expected)
	})
}

func assertCorrectSum(t *testing.T, sum, expected int) {
	if sum != expected {
		t.Errorf("expected %q but got %q", expected, sum)
	}
}
