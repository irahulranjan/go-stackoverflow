package main

import (
	"fmt"
	"strings"
	"github.com/irahulranjan/go-stackoverflow/csv"
)

func main() {
	fmt.Println("=== CSV Operations with Contains Function ===")
	fmt.Println()

	// Sample CSV data as string
	csvData := `name,age,city,department
John Doe,30,New York,Engineering
Jane Smith,25,San Francisco,Marketing
Bob Johnson,35,Chicago,Sales
Alice Brown,28,Boston,Engineering
Charlie Wilson,32,Seattle,Marketing`

	// Parse CSV data
	reader := strings.NewReader(csvData)
	records := csv.CSVToMap(reader)

	fmt.Printf("Parsed %d records from CSV:\n", len(records))
	for i, record := range records {
		fmt.Printf("  %d. %v\n", i+1, record)
	}
	fmt.Println()

	// Extract unique values for different fields
	fmt.Println("1. Finding unique departments:")
	var departments []string
	for _, record := range records {
		dept := record["department"]
		if !csv.Contains(departments, dept) {
			departments = append(departments, dept)
		}
	}
	fmt.Printf("   Unique departments: %v\n", departments)
	fmt.Println()

	// Check if specific values exist
	fmt.Println("2. Checking for specific values:")
	
	var cities []string
	for _, record := range records {
		cities = append(cities, record["city"])
	}
	
	searchCities := []string{"New York", "Los Angeles", "Chicago"}
	for _, city := range searchCities {
		found := csv.Contains(cities, city)
		fmt.Printf("   City '%s' exists in data: %t\n", city, found)
	}
	fmt.Println()

	// Find people in specific department
	fmt.Println("3. Finding people in Engineering department:")
	engineeringPeople := []string{}
	for _, record := range records {
		if record["department"] == "Engineering" {
			engineeringPeople = append(engineeringPeople, record["name"])
		}
	}
	fmt.Printf("   Engineering team: %v\n", engineeringPeople)
	
	// Check if specific person is in engineering
	checkPeople := []string{"John Doe", "Jane Smith", "David Lee"}
	for _, person := range checkPeople {
		isInEngineering := csv.Contains(engineeringPeople, person)
		fmt.Printf("   Is '%s' in Engineering: %t\n", person, isInEngineering)
	}
	fmt.Println()

	// Age-based analysis
	fmt.Println("4. Age-based filtering:")
	var ages []string
	for _, record := range records {
		ages = append(ages, record["age"])
	}
	
	targetAges := []string{"25", "30", "40"}
	for _, age := range targetAges {
		found := csv.Contains(ages, age)
		fmt.Printf("   Someone is %s years old: %t\n", age, found)
	}
}