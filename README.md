# Go StackOverflow Utilities

[![CI](https://github.com/irahulranjan/go-stackoverflow/actions/workflows/ci.yml/badge.svg)](https://github.com/irahulranjan/go-stackoverflow/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/irahulranjan/go-stackoverflow/branch/master/graph/badge.svg)](https://codecov.io/gh/irahulranjan/go-stackoverflow)
[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/irahulranjan/go-stackoverflow)
[![Go Report Card](https://goreportcard.com/badge/github.com/irahulranjan/go-stackoverflow)](https://goreportcard.com/report/github.com/irahulranjan/go-stackoverflow)

A collection of utility functions for Go, including CSV operations and generic helper functions.

## Features

- **CSV Operations**: Parse and encode CSV data with ease
- **Generic Contains Function**: Check if an item exists in a slice using Go generics
- **Type-Safe**: Leverages Go's type system for compile-time safety
- **Well-Tested**: Comprehensive test coverage with CI/CD pipeline

## Installation

```bash
go get github.com/irahulranjan/go-stackoverflow
```

## Requirements

- Go 1.18+ (required for generics support)

## Quick Start

```go
package main

import (
    "fmt"
    "strings"
    "github.com/irahulranjan/go-stackoverflow/csv"
)

func main() {
    // Using the generic Contains function
    fruits := []string{"apple", "banana", "cherry"}
    fmt.Println(csv.Contains(fruits, "banana")) // true
    fmt.Println(csv.Contains(fruits, "orange")) // false
    
    // Working with CSV data
    csvData := `name,age,city
John,30,New York
Jane,25,Boston`
    
    reader := strings.NewReader(csvData)
    records := csv.CSVToMap(reader)
    fmt.Printf("Parsed %d records\n", len(records))
}
```

## Package Documentation

### csv

Package csv provides utility functions to deal with CSV files and generic slice operations.

#### Functions

##### `EncodeCSV(columns []string, rows []map[string]string, writer io.Writer) error`

Encodes a slice of maps to a CSV file.

**Parameters:**
- `columns`: Column headers for the CSV
- `rows`: Data rows as maps
- `writer`: Writer to output the CSV data

**Example:**
```go
columns := []string{"name", "age"}
rows := []map[string]string{
    {"name": "John", "age": "30"},
    {"name": "Jane", "age": "25"},
}

var buf bytes.Buffer
err := csv.EncodeCSV(columns, rows, &buf)
if err != nil {
    log.Fatal(err)
}
fmt.Println(buf.String())
```

##### `CSVToMap(reader io.Reader) []map[string]string`

Converts a CSV file to a slice of maps using comma as separator.

**Parameters:**
- `reader`: Reader containing CSV data

**Returns:**
- Slice of maps where each map represents a row

**Example:**
```go
csvData := `name,age,city
John,30,New York
Jane,25,Boston`

reader := strings.NewReader(csvData)
records := csv.CSVToMap(reader)
// records[0]["name"] = "John"
// records[0]["age"] = "30"
```

##### `CSVToMapWithSeparator(reader io.Reader, separator rune) []map[string]string`

Converts a CSV file to a slice of maps with a custom separator.

**Parameters:**
- `reader`: Reader containing CSV data
- `separator`: Custom separator character

##### `Contains[T comparable](list []T, item T) bool` ⭐ **New!**

Checks if an item is present in a list of items using generics. Works with any comparable type.

**Type Parameters:**
- `T`: Any comparable type (strings, numbers, booleans, comparable structs)

**Parameters:**
- `list`: Slice to search in
- `item`: Item to search for

**Returns:**
- `true` if item is found, `false` otherwise

**Examples:**
```go
// With strings
fruits := []string{"apple", "banana", "cherry"}
found := csv.Contains(fruits, "banana") // true

// With integers  
numbers := []int{1, 2, 3, 4, 5}
found = csv.Contains(numbers, 3) // true

// With custom types
type Product struct {
    ID   int
    Name string
}

products := []Product{{1, "Laptop"}, {2, "Mouse"}}
found = csv.Contains(products, Product{1, "Laptop"}) // true

// With any comparable type
booleans := []bool{true, false}
found = csv.Contains(booleans, false) // true
```

## Examples

The `examples/` directory contains comprehensive examples:

- **[basic-usage/](examples/basic-usage/)**: Basic usage of the Contains function with different types
- **[csv-with-contains/](examples/csv-with-contains/)**: Combining CSV operations with Contains for data analysis  
- **[advanced-usage/](examples/advanced-usage/)**: Advanced usage patterns, performance considerations, and edge cases

### Running Examples

```bash
# Run basic usage example
cd examples/basic-usage && go run main.go

# Run CSV with Contains example  
cd examples/csv-with-contains && go run main.go

# Run advanced usage example
cd examples/advanced-usage && go run main.go
```

## Performance

The `Contains` function provides O(n) time complexity for searching. For large datasets or frequent lookups, consider using maps for O(1) average-case performance:

```go
// For frequent lookups, convert to map
items := []string{"apple", "banana", "cherry"}
itemSet := make(map[string]bool)
for _, item := range items {
    itemSet[item] = true
}

// O(1) lookup
found := itemSet["banana"]
```

## Testing

Run the test suite:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

*Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)*
