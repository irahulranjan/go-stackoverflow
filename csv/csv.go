// Package csv provides utility functions to deal with csvs
package csv

import (
	"encoding/csv"
	"io"
	"log"
)

// EncodeCSV encodes slice of maps to a csv file
func EncodeCSV(columns []string, rows []map[string]string, writer io.Writer) (err error) {
	w := csv.NewWriter(writer)
	if err = w.Write(columns); err != nil {
		return
	}
	r := make([]string, len(columns))
	for _, row := range rows {
		for i, column := range columns {
			r[i] = row[column]
		}
		if err = w.Write(r); err != nil {
			log.Println(err)
			return
		}
	}
	w.Flush()
	err = w.Error()
	return
}

// CSVToMap converts a csv file to a slice of maps
func CSVToMap(reader io.Reader) []map[string]string {
	return CSVToMapWithSeparator(reader, ',')
}

func CSVToMapWithSeparator(reader io.Reader, separator rune) []map[string]string {
	r := csv.NewReader(reader)
	r.LazyQuotes = true
	r.FieldsPerRecord = -1
	r.Comma = separator
	rows := []map[string]string{}
	var header []string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if header == nil {
			header = record
		} else {
			dict := map[string]string{}
			for i := range header {
				dict[header[i]] = record[i]
			}
			rows = append(rows, dict)
		}
	}
	return rows
}

// Contains checks if an item is present in a list of items using generics.
// It works with any comparable type (strings, numbers, booleans, etc.).
func Contains[T comparable](list []T, item T) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}
