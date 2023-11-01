package main

import (
	"fmt"
	"strings"
)

const batchSize = 10000

func GenerateInsertStatement(tableName string, columnOrder []string, entities []Entity, fields []FieldMapping) []string {
	var statements []string

	// Create a map for faster look-up
	fieldMap := make(map[string]FieldMapping)
	for _, field := range fields {
		fieldMap[field.DBColumnName] = field
	}

	backtickedColumns := make([]string, len(columnOrder))
	for i, col := range columnOrder {
		backtickedColumns[i] = fmt.Sprintf("`%s`", col)
	}

	for i := 0; i < len(entities); i += batchSize {
		var allValues []string
		end := i + batchSize

		if end > len(entities) {
			end = len(entities)
		}

		for j := i; j < end; j++ {
			var rowValues []string
			for _, col := range columnOrder {
				if field, exists := fieldMap[col]; exists {
					value := field.MappingFunction(entities[j])
					switch v := value.(type) {
					case string:
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
			allValues = append(allValues, fmt.Sprintf("(%s)", strings.Join(rowValues, ", ")))
		}

		statement := fmt.Sprintf("INSERT IGNORE INTO `%s` (%s) VALUES\n%s;", tableName, strings.Join(backtickedColumns, ", "), strings.Join(allValues, ",\n"))
		statements = append(statements, statement)
	}

	return statements
}
