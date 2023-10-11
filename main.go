package main

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

func main() {

	products()
	categories()
	orders()

	fmt.Println("SQL file has been generated successfully.")
}

func categories() {

	// Open the file
	file, err := os.OpenFile("./data/product_cleaned.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Decode the CSV data
	var cats []CategoryRecord
	if err := gocsv.UnmarshalFile(file, &cats); err != nil {
		fmt.Println("Error:", err)
		return
	}

	sqlFile, err := os.Create("./data/import_categories.sql")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer sqlFile.Close()

	sqlFile.WriteString("TRUNCATE TABLE `oc_category`;\n")
	sqlFile.WriteString("TRUNCATE TABLE `oc_category_description`;\n")
	sqlFile.WriteString("TRUNCATE TABLE `oc_category_path`;\n")
	sqlFile.WriteString("TRUNCATE TABLE `oc_category_to_store`;\n")

	for _, category := range cats {
		for _, stmt := range GenerateCategorySQLStatements(category) {
			sqlFile.WriteString(stmt + "\n")
		}
	}
}

func orders() {
	orderMapping := GetOrderMapping()

	// Open the ordersFile
	ordersFile, err := os.OpenFile("./data/orders_short.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer ordersFile.Close()

	// Decode the CSV data
	var orders []OrderRecord
	if err := gocsv.UnmarshalFile(ordersFile, &orders); err != nil {
		fmt.Println("Error:", err)
		return
	}

	sqlFile, err := os.Create("./data/import_orders.sql")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer sqlFile.Close()

	orderIDMapping := make(map[string]int)
	for index, order := range orders {
		orderIDMapping[order.OrderID] = index + 1
	}

	entities := make([]Entity, len(orders))
	for i, v := range orders {
		entities[i] = v
	}

	// Open the product CSV file
	productFile, err := os.OpenFile("./data/products_export_full_20230815_210049_71306.csv", os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer productFile.Close()

	// Decode the product data from CSV
	var products []ProductRecord
	if err := gocsv.UnmarshalFile(productFile, &products); err != nil {
		fmt.Println("Error:", err)
		return
	}

	productIdMapping := make(map[string]int)
	for index, product := range products {
		productIdMapping[product.Model] = index + 1
	}


	// Use the helper function for each mapping
	processTable(orderMapping, entities, orderIDMapping, sqlFile)
	processTable(GetOrderProductMapping(orderIDMapping, productIdMapping), entities, orderIDMapping, sqlFile)

	processTable(GetOrderTotalMapping(orderIDMapping), entities, orderIDMapping, sqlFile)
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
	processTable(GetProductDiscountMapping(productIdMapping), entities, productIdMapping, sqlFile)
	processTable(GetProductToStoreMapping(productIdMapping), entities, productIdMapping, sqlFile)
	processTable(GetProductToCostMapping(productIdMapping), entities, productIdMapping, sqlFile)
	processTable(GetProductToCategoryMapping(productIdMapping), entities, productIdMapping, sqlFile)
}

func processTable(tableMapping TableMapping, entities []Entity, productIdMapping map[string]int, sqlFile *os.File) {
	sqlFile.WriteString("TRUNCATE TABLE `" + tableMapping.TableName + "`;\n")
	insertStatements := GenerateInsertStatement(tableMapping.TableName, tableMapping.ColumnOrder, entities, tableMapping.Fields, "Model", productIdMapping)
	for _, stmt := range insertStatements {
		sqlFile.WriteString(stmt + "\n")
	}
}
