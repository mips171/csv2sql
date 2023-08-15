package main

import (
	"fmt"
	"strings"
)

func GenerateInsertStatement(tableName string, columnOrder []string, records [][]string, fields []FieldMapping, customerIDCounter *int, customerIdMapping map[string]int) []string {
	var statements []string

	fieldIndexMap := make(map[string]int)
	for i, name := range records[0] {
		fieldIndexMap[name] = i
	}

	for _, record := range records[1:] {
		var rowValues []string
		if customerIDCounter != nil {
			rowValues = append(rowValues, fmt.Sprintf("%d", *customerIDCounter))
			customerIdMapping[record[fieldIndexMap["Email Address"]]] = *customerIDCounter
			(*customerIDCounter)++
		} else {
			customerId := customerIdMapping[record[fieldIndexMap["Email Address"]]]
			rowValues = append(rowValues, fmt.Sprintf("%d", customerId))
		}

		email := record[fieldIndexMap["Email Address"]]

		for _, col := range columnOrder[1:] {
			for _, field := range fields {
				if col == field.DbColumnName {
					value := field.Transformation(record[fieldIndexMap[field.CsvFieldName]], email)
					rowValues = append(rowValues, fmt.Sprintf("'%s'", strings.ReplaceAll(value, "'", "''")))
				}
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
