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

	var valuesGroup []string
	for _, entity := range entities {
		var rowValues []string

		for _, col := range columnOrder {
			if field, exists := fieldMap[col]; exists {
				value := field.MappingFunction(entity)
				switch v := value.(type) {
				case string:
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
		valuesGroup = append(valuesGroup, fmt.Sprintf("(%s)", strings.Join(rowValues, ", ")))
	}

	backtickedColumns := make([]string, len(columnOrder))
	for i, col := range columnOrder {
		backtickedColumns[i] = fmt.Sprintf("`%s`", col)
	}

	// LOCK and DISABLE KEYS statements
	statements = append(statements, fmt.Sprintf("LOCK TABLES `%s` WRITE;", tableName))
	statements = append(statements, fmt.Sprintf("/*!40000 ALTER TABLE `%s` DISABLE KEYS */;", tableName))

	// Main INSERT statement
	statements = append(statements, fmt.Sprintf("INSERT INTO `%s` (%s) VALUES\n%s;", tableName, strings.Join(backtickedColumns, ", "), strings.Join(valuesGroup, ",\n")))

	// ENABLE KEYS and UNLOCK TABLES statements
	statements = append(statements, fmt.Sprintf("/*!40000 ALTER TABLE `%s` ENABLE KEYS */;", tableName))
	statements = append(statements, "UNLOCK TABLES;")

	return statements
}
