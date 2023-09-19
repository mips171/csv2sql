package main

import (
	"fmt"
	"strings"
)

func GenerateInsertStatement(tableName string, columnOrder []string, records []map[string]string, fields []FieldMapping, uniqueIdentifier string, uniqueIdMapping map[string]int) []string {
	var statements []string

	// Create a map for faster look-up
	fieldMap := make(map[string]FieldMapping)
	for _, field := range fields {
		fieldMap[field.DbColumnName] = field
	}

	for _, record := range records {
		var rowValues []string

		// If a uniqueIdMapping is provided, use it.
		if uniqueIdMapping != nil {
			if id, ok := uniqueIdMapping[record[uniqueIdentifier]]; ok {
				rowValues = append(rowValues, fmt.Sprintf("%d", id))
			} else {
				continue // Skip the record if no unique ID is found
			}
		}

		for _, col := range columnOrder {
			if field, exists := fieldMap[col]; exists {
				value := field.Transformation(record[field.CsvFieldName], record[uniqueIdentifier])
				rowValues = append(rowValues, fmt.Sprintf("'%s'", strings.ReplaceAll(value, "'", "''")))
			} else {
				rowValues = append(rowValues, "NULL") // Handle fields not present in the CSV
			}
		}

		backtickedColumns := make([]string, len(columnOrder))
		for i, col := range columnOrder {
			backtickedColumns[i] = fmt.Sprintf("`%s`", col)
		}

		statement := fmt.Sprintf("INSERT INTO `%s` (%s) VALUES (%s);", tableName, strings.Join(backtickedColumns, ", "), strings.Join(rowValues, ", "))
		statements = append(statements, statement)
	}

	return statements
}
