package main

import (
	"fmt"
	"os"
)

func main() {
	customerMapping := GetCustomerMapping()
	addressMapping := GetAddressMapping()

	records, err := ReadCsv("./data/old_data.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	sqlFile, err := os.Create("./data/import_data.sql")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer sqlFile.Close()

	customerIDCounter := 1                       // This will act as our customer_id
	var customerIdMapping = make(map[string]int) // Mapping from email to customer_id
	// Generate the SQL statements for oc_customer table
	customerInsertStatements := GenerateInsertStatement(customerMapping.TableName, customerMapping.ColumnOrder, records, customerMapping.Fields, &customerIDCounter, customerIdMapping)
	for _, stmt := range customerInsertStatements {
		sqlFile.WriteString(stmt + "\n")
	}

	// Generate the SQL statements for oc_address table
	addressInsertStatements := GenerateInsertStatement(addressMapping.TableName, addressMapping.ColumnOrder, records, addressMapping.Fields, nil, customerIdMapping)
	for _, stmt := range addressInsertStatements {
		sqlFile.WriteString(stmt + "\n")
	}

	fmt.Println("SQL file has been generated successfully.")
}
