package main

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

func main() {

	products()

	fmt.Println("SQL file has been generated successfully.")
}

func products() {
	productMapping := GetProductMapping()

	// Open the file
	file, err := os.OpenFile("./data/product_cleaned.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Decode the CSV data
	var products []ProductRecord
	if err := gocsv.UnmarshalFile(file, &products); err != nil {
		fmt.Println("Error:", err)
		return
	}

	sqlFile, err := os.Create("./data/import_products.sql")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer sqlFile.Close()

	productIdMapping := make(map[string]int)
	for index, product := range products {
		productIdMapping[product.Model] = index + 1
	}

	entities := make([]Entity, len(products))
	for i, v := range products {
		entities[i] = v
	}

	// Use the helper function for each mapping
	processTable(productMapping, entities, productIdMapping, sqlFile)
	processTable(GetProductDescriptionMapping(productIdMapping), entities, productIdMapping, sqlFile)
	processTable(GetProductSpecialMapping(productIdMapping), entities, productIdMapping, sqlFile)
	processTable(GetProductToStoreMapping(productIdMapping), entities, productIdMapping, sqlFile)
	processTable(GetProductToCostMapping(productIdMapping), entities, productIdMapping, sqlFile)
}

func processTable(tableMapping TableMapping, entities []Entity, productIdMapping map[string]int, sqlFile *os.File) {
	sqlFile.WriteString("TRUNCATE TABLE `" + tableMapping.TableName + "`;\n")
	insertStatements := GenerateInsertStatement(tableMapping.TableName, tableMapping.ColumnOrder, entities, tableMapping.Fields, "Model", productIdMapping)
	for _, stmt := range insertStatements {
		sqlFile.WriteString(stmt + "\n")
	}
}
