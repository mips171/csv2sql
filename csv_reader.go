package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

// ReadCsv opens and reads the given CSV file and returns its records
func ReadCsv(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return records, nil
}
