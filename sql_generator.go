package main

import (
	"fmt"
	"strings"
)

func GenerateInsertStatement(tableName string, columnOrder []string, entities []Entity, fields []FieldMapping, uniqueIdentifier string, uniqueIdMapping map[string]int) []string {
	var statements []string

	// Create a map for faster look-up
	fieldMap := make(map[string]FieldMapping)
	for _, field := range fields {
		fieldMap[field.DBColumnName] = field
	}

	for _, entity := range entities {
		var rowValues []string

		for _, col := range columnOrder {
			if field, exists := fieldMap[col]; exists {
				value := field.MappingFunction(entity)
				switch v := value.(type) {
				case string:
					// replace all \ with ""
					v = strings.ReplaceAll(v, "\\", "")
					rowValues = append(rowValues, fmt.Sprintf("'%s'", strings.ReplaceAll(v, "'", "''")))
				case int:
					rowValues = append(rowValues, fmt.Sprintf("%d", v))
				default:
					// Handle other types or raise an error if needed
				}
			} else {
				rowValues = append(rowValues, "NULL") // Handle fields not present in the CSV
			}
		}

		backtickedColumns := make([]string, len(columnOrder))
		for i, col := range columnOrder {
			backtickedColumns[i] = fmt.Sprintf("`%s`", col)
		}

		statement := fmt.Sprintf("INSERT IGNORE INTO `%s` (%s) VALUES (%s);", tableName, strings.Join(backtickedColumns, ", "), strings.Join(rowValues, ", "))
		statements = append(statements, statement)
	}

	return statements
}
