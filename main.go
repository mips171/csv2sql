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
	file, err := os.OpenFile("./data/product_cleaned_short.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
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

	// Global mapping from SKU to product_id
	var productIdMapping = make(map[string]int)

	var records []map[string]string
	for index, product := range products {
		records = append(records, productToMap(product))
		productIdMapping[product.Model] = index + 1
	}

	// Generate product statements
	sqlFile.WriteString("TRUNCATE TABLE `" + productMapping.TableName + "`;\n")
	productInsertStatements := GenerateInsertStatement(productMapping.TableName, productMapping.ColumnOrder, records, productMapping.Fields, "Model", productIdMapping)
	for _, stmt := range productInsertStatements {
		sqlFile.WriteString(stmt + "\n")
	}

	// Generate description statements
	var descriptionRecords []map[string]string
	descriptionMapping := GetProductDescriptionMapping(productIdMapping)
	sqlFile.WriteString("TRUNCATE TABLE `" + descriptionMapping.TableName + "`;\n")
	for _, product := range products {
		descriptionRecords = append(descriptionRecords, productToMap(product))
	}
	descriptionInsertStatements := GenerateInsertStatement(descriptionMapping.TableName, descriptionMapping.ColumnOrder, descriptionRecords, descriptionMapping.Fields, "Model", productIdMapping)
	for _, stmt := range descriptionInsertStatements {
		sqlFile.WriteString(stmt + "\n")
	}

	// Generate trade price statements
	var tradePriceRecords []map[string]string
	tradePriceMapping := GetProductSpecialMapping(productIdMapping)
	sqlFile.WriteString("TRUNCATE TABLE `" + tradePriceMapping.TableName + "`;\n")
	for _, product := range products {
		tradePriceRecords = append(tradePriceRecords, productToMap(product))
	}
	tradePriceInsertStatements := GenerateInsertStatement(tradePriceMapping.TableName, tradePriceMapping.ColumnOrder, tradePriceRecords, tradePriceMapping.Fields, "Model", productIdMapping)
	for _, stmt := range tradePriceInsertStatements {
		sqlFile.WriteString(stmt + "\n")
	}

	var productToStoreRecords []map[string]string
	productToStoreMapping := GetProductToStoreMapping(productIdMapping)
	sqlFile.WriteString("TRUNCATE TABLE `" + productToStoreMapping.TableName + "`;\n")
	for _, product := range products {
		productToStoreRecords = append(productToStoreRecords, productToMap(product))
	}
	productToStoreInsertStatements := GenerateInsertStatement(productToStoreMapping.TableName, productToStoreMapping.ColumnOrder, productToStoreRecords, productToStoreMapping.Fields, "Model", productIdMapping)
	for _, stmt := range productToStoreInsertStatements {
		sqlFile.WriteString(stmt + "\n")
	}

	var productToCostRecords []map[string]string
	productToCostMapping := GetProductToCostMapping(productIdMapping)
	sqlFile.WriteString("TRUNCATE TABLE `" + productToCostMapping.TableName + "`;\n")
	for _, product := range products {
		productToCostRecords = append(productToCostRecords, productToMap(product))
	}
	productToCostInsertStatements := GenerateInsertStatement(productToCostMapping.TableName, productToCostMapping.ColumnOrder, productToCostRecords, productToCostMapping.Fields, "Model", productIdMapping)
	for _, stmt := range productToCostInsertStatements {
		sqlFile.WriteString(stmt + "\n")
	}
}
