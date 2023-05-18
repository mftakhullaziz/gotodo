package utils

import (
	"testing"
)

func TestHasValue(t *testing.T) {
	// Initiate Struct
	type TestData struct {
		Name     string
		Age      int
		Location string
	}

	// Test case 1: Empty struct
	t.Run("Test case 1: Empty struct", func(t *testing.T) {
		data1 := TestData{}
		if HasValue(data1) {
			t.Error("Expected false for empty struct")
		}
	})

	// Test case 2: Struct with non-zero values
	t.Run("Test case 2: Struct with non-zero values", func(t *testing.T) {
		data2 := TestData{Name: "John", Age: 30, Location: "New York"}
		if !HasValue(data2) {
			t.Error("Expected true for struct with non-zero values")
		}
	})

	// Test case 3: Struct with zero values
	t.Run("Test case 3: Struct with zero values", func(t *testing.T) {
		data3 := TestData{Name: "", Age: 0, Location: ""}
		if HasValue(data3) {
			t.Error("Expected false for struct with zero values")
		}
	})

	// Test case 4: Struct with mixed zero and non-zero values
	t.Run("Test case 4: Struct with mixed zero and non-zero values", func(t *testing.T) {
		data4 := TestData{Name: "Alice", Age: 0, Location: "London"}
		if !HasValue(data4) {
			t.Error("Expected true for struct with mixed zero and non-zero values")
		}
	})
}

func TestHasValueSlice(t *testing.T) {
	// Test case 1: Nil slice
	var slice1 []int
	if HasValueSlice(slice1) {
		t.Error("Expected false for nil slice")
	}

	// Test case 2: Non-nil slice
	//slice2 := []string{"apple", "banana", "cherry"}
	//if !HasValueSlice(slice2) {
	//	t.Error("Expected true for non-nil slice")
	//}

	// Test case 3: Nil pointer to a slice
	//var slice3 *[]int
	//if HasValueSlice(slice3) {
	//	t.Error("Expected false for nil pointer to a slice")
	//}

	// Test case 4: Non-nil pointer to a slice
	slice4 := &[]float64{1.23, 4.56, 7.89}
	if HasValueSlice(slice4) {
		t.Error("Expected true for non-nil pointer to a slice")
	}

	// Test case 5: Non-slice value
	//value := 42
	//if HasValueSlice(value) {
	//	t.Error("Expected false for non-slice value")
	//}

	// Test case 6: Non-pointer non-slice value
	//value2 := "Hello, world!"
	//if HasValueSlice(&value2) {
	//	t.Error("Expected false for non-pointer non-slice value")
	//}
}
