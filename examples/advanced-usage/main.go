package main

import (
	"fmt"
	"time"
	"github.com/irahulranjan/go-stackoverflow/csv"
)

func main() {
	fmt.Println("=== Advanced Usage and Performance Examples ===")
	fmt.Println()

	// Example 1: Edge cases
	fmt.Println("1. Edge Cases:")
	
	// Empty slice
	emptyList := []string{}
	fmt.Printf("   Empty list contains 'anything': %t\n", csv.Contains(emptyList, "anything"))
	
	// Single item
	singleItem := []int{42}
	fmt.Printf("   Single item [42] contains 42: %t\n", csv.Contains(singleItem, 42))
	fmt.Printf("   Single item [42] contains 24: %t\n", csv.Contains(singleItem, 24))
	
	// Nil values in slice (with pointers)
	var nilPtr *string
	validPtr := &[]string{"test"}[0]
	pointerSlice := []*string{nilPtr, validPtr, nilPtr}
	fmt.Printf("   Pointer slice contains nil: %t\n", csv.Contains(pointerSlice, nilPtr))
	fmt.Printf("   Pointer slice contains valid pointer: %t\n", csv.Contains(pointerSlice, validPtr))
	fmt.Println()

	// Example 2: Performance with large datasets
	fmt.Println("2. Performance Testing:")
	
	// Create a large slice
	largeSlice := make([]int, 100000)
	for i := range largeSlice {
		largeSlice[i] = i
	}
	
	// Test performance for item at beginning
	start := time.Now()
	found := csv.Contains(largeSlice, 5)
	duration := time.Since(start)
	fmt.Printf("   Found item at beginning (5): %t in %v\n", found, duration)
	
	// Test performance for item at end
	start = time.Now()
	found = csv.Contains(largeSlice, 99995)
	duration = time.Since(start)
	fmt.Printf("   Found item at end (99995): %t in %v\n", found, duration)
	
	// Test performance for non-existent item
	start = time.Now()
	found = csv.Contains(largeSlice, 200000)
	duration = time.Since(start)
	fmt.Printf("   Searched for non-existent item (200000): %t in %v\n", found, duration)
	fmt.Println()

	// Example 3: Working with different numeric types
	fmt.Println("3. Different Numeric Types:")
	
	int8Slice := []int8{1, 2, 3, 127}
	fmt.Printf("   int8 slice contains 127: %t\n", csv.Contains(int8Slice, int8(127)))
	
	uint64Slice := []uint64{1000000000000, 2000000000000}
	fmt.Printf("   uint64 slice contains large number: %t\n", csv.Contains(uint64Slice, uint64(1000000000000)))
	
	float32Slice := []float32{1.1, 2.2, 3.3}
	fmt.Printf("   float32 slice contains 2.2: %t\n", csv.Contains(float32Slice, float32(2.2)))
	fmt.Println()

	// Example 4: Complex structs
	fmt.Println("4. Complex Comparable Structs:")
	
	type Address struct {
		Street string
		City   string
		Zip    int
	}
	
	type Person struct {
		Name    string
		Age     int
		Address Address
	}
	
	people := []Person{
		{Name: "Alice", Age: 30, Address: Address{Street: "123 Main St", City: "Boston", Zip: 12345}},
		{Name: "Bob", Age: 25, Address: Address{Street: "456 Oak Ave", City: "Seattle", Zip: 67890}},
	}
	
	searchPerson := Person{Name: "Alice", Age: 30, Address: Address{Street: "123 Main St", City: "Boston", Zip: 12345}}
	found = csv.Contains(people, searchPerson)
	fmt.Printf("   Found exact person match: %t\n", found)
	
	// Same person but different age - should not match
	differentPerson := Person{Name: "Alice", Age: 31, Address: Address{Street: "123 Main St", City: "Boston", Zip: 12345}}
	found = csv.Contains(people, differentPerson)
	fmt.Printf("   Found person with different age: %t\n", found)
	fmt.Println()

	// Example 5: Benchmarking helper function
	fmt.Println("5. Benchmarking Different Approaches:")
	testSlice := []string{"apple", "banana", "cherry", "date", "elderberry"}
	
	// Time multiple lookups
	start = time.Now()
	for i := 0; i < 10000; i++ {
		csv.Contains(testSlice, "cherry")
	}
	duration = time.Since(start)
	fmt.Printf("   10,000 lookups using Contains: %v\n", duration)
	
	// Compare with manual loop
	start = time.Now()
	for i := 0; i < 10000; i++ {
		found := false
		for _, item := range testSlice {
			if item == "cherry" {
				found = true
				break
			}
		}
		_ = found
	}
	duration = time.Since(start)
	fmt.Printf("   10,000 lookups using manual loop: %v\n", duration)
}