package csv

import (
	"testing"
)

func TestContains(t *testing.T) {
	// Test with strings
	t.Run("strings", func(t *testing.T) {
		stringList := []string{"apple", "banana", "cherry"}

		if !Contains(stringList, "apple") {
			t.Error("Expected to find 'apple' in string list")
		}

		if !Contains(stringList, "banana") {
			t.Error("Expected to find 'banana' in string list")
		}

		if Contains(stringList, "orange") {
			t.Error("Expected NOT to find 'orange' in string list")
		}
	})

	// Test with integers
	t.Run("integers", func(t *testing.T) {
		intList := []int{1, 2, 3, 4, 5}

		if !Contains(intList, 3) {
			t.Error("Expected to find 3 in int list")
		}

		if Contains(intList, 6) {
			t.Error("Expected NOT to find 6 in int list")
		}
	})

	// Test with booleans
	t.Run("booleans", func(t *testing.T) {
		boolList := []bool{true, false, true}

		if !Contains(boolList, true) {
			t.Error("Expected to find true in bool list")
		}

		if !Contains(boolList, false) {
			t.Error("Expected to find false in bool list")
		}
	})

	// Test with empty list
	t.Run("empty list", func(t *testing.T) {
		emptyList := []string{}

		if Contains(emptyList, "anything") {
			t.Error("Expected NOT to find anything in empty list")
		}
	})

	// Test with float64
	t.Run("floats", func(t *testing.T) {
		floatList := []float64{1.1, 2.2, 3.3}

		if !Contains(floatList, 2.2) {
			t.Error("Expected to find 2.2 in float list")
		}

		if Contains(floatList, 4.4) {
			t.Error("Expected NOT to find 4.4 in float list")
		}
	})
}
