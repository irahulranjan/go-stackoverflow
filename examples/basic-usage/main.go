package main

import (
	"fmt"
	"github.com/irahulranjan/go-stackoverflow/csv"
)

func main() {
	fmt.Println("=== Basic Usage of Contains Function ===")
	fmt.Println()

	// Example 1: Working with strings
	fmt.Println("1. String Lists:")
	fruits := []string{"apple", "banana", "cherry", "date", "elderberry"}
	fmt.Printf("   Fruits: %v\n", fruits)
	
	searchItems := []string{"banana", "grape", "cherry"}
	for _, item := range searchItems {
		found := csv.Contains(fruits, item)
		fmt.Printf("   Contains '%s': %t\n", item, found)
	}
	fmt.Println()

	// Example 2: Working with integers
	fmt.Println("2. Integer Lists:")
	numbers := []int{1, 3, 5, 7, 9, 11, 13}
	fmt.Printf("   Numbers: %v\n", numbers)
	
	searchNumbers := []int{5, 8, 11}
	for _, num := range searchNumbers {
		found := csv.Contains(numbers, num)
		fmt.Printf("   Contains %d: %t\n", num, found)
	}
	fmt.Println()

	// Example 3: Working with floats
	fmt.Println("3. Float Lists:")
	prices := []float64{9.99, 19.99, 29.99, 49.99}
	fmt.Printf("   Prices: %v\n", prices)
	
	searchPrices := []float64{19.99, 39.99}
	for _, price := range searchPrices {
		found := csv.Contains(prices, price)
		fmt.Printf("   Contains $%.2f: %t\n", price, found)
	}
	fmt.Println()

	// Example 4: Working with custom types (comparable structs)
	fmt.Println("4. Custom Types:")
	type Product struct {
		ID   int
		Name string
	}
	
	products := []Product{
		{ID: 1, Name: "Laptop"},
		{ID: 2, Name: "Mouse"},
		{ID: 3, Name: "Keyboard"},
	}
	fmt.Printf("   Products: %v\n", products)
	
	searchProduct := Product{ID: 2, Name: "Mouse"}
	found := csv.Contains(products, searchProduct)
	fmt.Printf("   Contains %v: %t\n", searchProduct, found)
	
	notFoundProduct := Product{ID: 4, Name: "Monitor"}
	found = csv.Contains(products, notFoundProduct)
	fmt.Printf("   Contains %v: %t\n", notFoundProduct, found)
}