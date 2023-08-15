package main

// import (
// 	"encoding/csv"
// 	"fmt"
// 	"os"
// 	"strings"
// 	"time"
// )

// type FieldMapping struct {
// 	CsvFieldName   string
// 	DbColumnName   string
// 	Transformation func(string, string) string
// }

// type TableMapping struct {
// 	TableName   string
// 	ColumnOrder []string
// 	Fields      []FieldMapping
// }

// func TransformIdentity(value string) string {
// 	return value
// }

// func GetUserGroupID(groupName string) string {
// 	switch groupName {
// 	case "Retail":
// 		return "1"
// 	case "Trade":
// 		return "2"
// 	default:
// 		return "1"
// 	}
// }

// func GetNewsletterStatus(subscriberStatus string) string {
// 	if subscriberStatus == "y" {
// 		return "1"
// 	}
// 	return "0"
// }

// func GetStatus(activeStatus string) string {
// 	if activeStatus == "y" {
// 		return "1"
// 	}
// 	return "0"
// }

// func GetDateAdded() string {
// 	dateAdded := time.Now().Format("2006-01-02 15:04:05")
// 	return dateAdded
// }

// func GetFirstName(value string, email string) string {
// 	if value != "" {
// 		return value
// 	}
// 	parts := strings.Split(email, "@")
// 	if len(parts) > 0 {
// 		return parts[0]
// 	}
// 	return "N/A"
// }

// func GetLastName(value string, email string) string {
// 	if value != "" {
// 		return value
// 	}
// 	parts := strings.Split(email, "@")
// 	if len(parts) > 1 {
// 		domainParts := strings.Split(parts[1], ".")
// 		if len(domainParts) > 0 {
// 			return domainParts[0]
// 		}
// 	}
// 	return "N/A"
// }

// func GenerateInsertStatement(tableName string, columnOrder []string, records [][]string, fields []FieldMapping, customerIDCounter *int, customerIdMapping map[string]int) []string {
// 	var statements []string

// 	fieldIndexMap := make(map[string]int)
// 	for i, name := range records[0] {
// 		fieldIndexMap[name] = i
// 	}

// 	for _, record := range records[1:] {
// 		var rowValues []string
// 		if customerIDCounter != nil {
// 			rowValues = append(rowValues, fmt.Sprintf("%d", *customerIDCounter))
// 			customerIdMapping[record[fieldIndexMap["Email Address"]]] = *customerIDCounter
// 			(*customerIDCounter)++
// 		} else {
// 			customerId := customerIdMapping[record[fieldIndexMap["Email Address"]]]
// 			rowValues = append(rowValues, fmt.Sprintf("%d", customerId))
// 		}

// 		email := record[fieldIndexMap["Email Address"]]

// 		for _, col := range columnOrder[1:] {
// 			for _, field := range fields {
// 				if col == field.DbColumnName {
// 					value := field.Transformation(record[fieldIndexMap[field.CsvFieldName]], email)
// 					rowValues = append(rowValues, fmt.Sprintf("'%s'", strings.ReplaceAll(value, "'", "''")))
// 				}
// 			}
// 		}

// 		backtickedColumns := make([]string, len(columnOrder))
// 		for i, col := range columnOrder {
// 			backtickedColumns[i] = fmt.Sprintf("`%s`", col)
// 		}

// 		statement := fmt.Sprintf("INSERT INTO `%s` (%s) VALUES (%s);", tableName, strings.Join(backtickedColumns, ", "), strings.Join(rowValues, ", "))
// 		statements = append(statements, statement)
// 	}

// 	return statements
// }

// func main() {
// 	// Define mappings for each table here
// 	customerMapping := TableMapping{
// 		TableName:   "oc_customer",
// 		ColumnOrder: []string{"customer_id", "customer_group_id", "firstname", "lastname", "email", "telephone", "newsletter", "status", "date_added"},
// 		Fields: []FieldMapping{
// 			{"User Group", "customer_group_id", func(value string, _ string) string { return GetUserGroupID(value) }},
// 			{"Bill First Name", "firstname", GetFirstName},
// 			{"Bill Last Name", "lastname", GetLastName},
// 			{"Email Address", "email", func(value string, _ string) string { return TransformIdentity(value) }},
// 			{"Bill Phone", "telephone", func(value string, _ string) string { return TransformIdentity(value) }},
// 			{"Newsletter Subscriber", "newsletter", func(value string, _ string) string { return GetNewsletterStatus(value) }},
// 			{"Active", "status", func(value string, _ string) string { return GetStatus(value) }},
// 			{"", "date_added", func(_ string, _ string) string { return GetDateAdded() }},
// 			// Add more fields if necessary
// 		},
// 	}

// 	addressMapping := TableMapping{
// 		TableName:   "oc_address",
// 		ColumnOrder: []string{"customer_id", "firstname", "lastname", "company", "address_1", "address_2", "city", "postcode"},
// 		Fields: []FieldMapping{
// 			{"Bill First Name", "firstname", func(value string, _ string) string { return TransformIdentity(value) }},
// 			{"Bill Last Name", "lastname", func(value string, _ string) string { return TransformIdentity(value) }},
// 			{"Bill Company", "company", func(value string, _ string) string { return TransformIdentity(value) }},
// 			{"Bill Street Address Line 1", "address_1", func(value string, _ string) string { return TransformIdentity(value) }},
// 			{"Bill Street Address Line 2", "address_2", func(value string, _ string) string { return TransformIdentity(value) }},
// 			{"Bill City", "city", func(value string, _ string) string { return TransformIdentity(value) }},
// 			{"Bill Post Code", "postcode", func(value string, _ string) string { return TransformIdentity(value) }},
// 			// Add more fields if necessary
// 		},
// 	}

// 	// Step 1: Read the CSV file
// 	file, err := os.Open("old_data.csv")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	defer file.Close()

// 	reader := csv.NewReader(file)
// 	records, err := reader.ReadAll()
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	// Step 2: Generate the SQL statements based on the old CSV content
// 	sqlFile, err := os.Create("import_data.sql")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	defer sqlFile.Close()

// 	customerIDCounter := 1                       // This will act as our customer_id
// 	var customerIdMapping = make(map[string]int) // Mapping from email to customer_id
// 	// Generate the SQL statements for oc_customer table
// 	customerInsertStatements := GenerateInsertStatement(customerMapping.TableName, customerMapping.ColumnOrder, records, customerMapping.Fields, &customerIDCounter, customerIdMapping)
// 	for _, stmt := range customerInsertStatements {
// 		sqlFile.WriteString(stmt + "\n")
// 	}

// 	// Generate the SQL statements for oc_address table
// 	addressInsertStatements := GenerateInsertStatement(addressMapping.TableName, addressMapping.ColumnOrder, records, addressMapping.Fields, nil, customerIdMapping)
// 	for _, stmt := range addressInsertStatements {
// 		sqlFile.WriteString(stmt + "\n")
// 	}

// 	fmt.Println("SQL file has been generated successfully.")
// }
