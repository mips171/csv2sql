package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
)

const (
	PRODUCTS_CSV        = "./data/products_export_full_20231101_114741_67572.csv"
	CUSTOMERS_CSV       = "./data/customer_export_full_20231101_114614_40007.csv"
	CUSTOMER_GROUPS_CSV = "./data/customer_groups.csv"
	ORDERS_CSV          = "./data/orders_export_full_20231101_114706_57298.csv"

	OUTPUT_CUSTOMERS_SQL  = "./data/import_customers.sql"
	OUTPUT_CATEGORIES_SQL = "./data/import_categories.sql"
	OUTPUT_ORDERS_SQL     = "./data/import_orders.sql"
	OUTPUT_PRODUCTS_SQL   = "./data/import_products.sql"
)

func main() {

	products()
	categories()
	orders()
	customers()

	fmt.Println("SQL file has been generated successfully.")
}

func customers() {

	customerMapping := GetCustomerMapping()

	// Open the file
	file, err := os.OpenFile(CUSTOMERS_CSV, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Decode the CSV data
	var customers []CustomerRecord
	if err := gocsv.UnmarshalFile(file, &customers); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Open the file
	groupsFile, err := os.OpenFile(CUSTOMER_GROUPS_CSV, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer groupsFile.Close()

	var customerGroups []CustomerGroupRecord
	if err := gocsv.UnmarshalFile(groupsFile, &customerGroups); err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i := range customers {
		for _, group := range customerGroups {
			if customers[i].Username == group.Username {
				customers[i].Group = group.Group
			}
		}
	}

	sqlFile, err := os.Create(OUTPUT_CUSTOMERS_SQL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer sqlFile.Close()

	customerIdMapping := make(map[string]int)
	for index, product := range customers {
		customerIdMapping[product.Email] = index + 1
	}

	entities := make([]Entity, len(customers))
	for i, v := range customers {
		entities[i] = v
	}

	processTable(customerMapping, entities, customerIdMapping, sqlFile)
	processTable(GetCustomerAddressMapping(customerIdMapping), entities, customerIdMapping, sqlFile)
}

func categories() {

	// Open the file
	file, err := os.OpenFile(PRODUCTS_CSV, os.O_RDWR|os.O_CREATE, os.ModePerm)
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

	sqlFile, err := os.Create(OUTPUT_CATEGORIES_SQL)
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

	// Open the ordersFile
	ordersFile, err := os.OpenFile(ORDERS_CSV, os.O_RDWR|os.O_CREATE, os.ModePerm)
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

	sqlFile, err := os.Create(OUTPUT_ORDERS_SQL)
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
	productFile, err := os.OpenFile(PRODUCTS_CSV, os.O_RDWR, os.ModePerm)
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

	// map customer email to ID
	// Open the product CSV file
	customersFile, err := os.OpenFile(CUSTOMERS_CSV, os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer customersFile.Close()

	// Decode the product data from CSV
	var customers []CustomerRecord
	if err := gocsv.UnmarshalFile(customersFile, &customers); err != nil {
		fmt.Println("Error:", err)
		return
	}

	customerIdMapping := make(map[string]int)
	for index, cust := range customers {
		customerIdMapping[cust.Email] = index + 1
	}

	orderLineItemsMap := make(map[string][]OrderRecord)
	for _, order := range orders {
		orderLineItemsMap[order.OrderID] = append(orderLineItemsMap[order.OrderID], order)
	}

	orderMapping := GetOrderMapping(customerIdMapping)

	// Use the helper function for each mapping
	processTable(orderMapping, entities, orderIDMapping, sqlFile)
	processTable(GetOrderProductMapping(productIdMapping), entities, orderIDMapping, sqlFile)

	sqlFile.WriteString("TRUNCATE TABLE `oc_order_total`;\n")

	// Process each group of line items per order
	for orderID, lineItems := range orderLineItemsMap {
		subTotalValue, shippingCost, taxValue, totalValue := CalculateOrderTotals(lineItems)
		sqlStatements := GenerateOrderTotalSQLStatements(orderID, subTotalValue, shippingCost, taxValue, totalValue)
		for _, stmt := range sqlStatements {
			sqlFile.WriteString(stmt + "\n")
		}
	}

}

func products() {
	productMapping := GetProductMapping()

	// Open the file
	file, err := os.OpenFile(PRODUCTS_CSV, os.O_RDWR|os.O_CREATE, os.ModePerm)
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

	sqlFile, err := os.Create(OUTPUT_PRODUCTS_SQL)
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
	insertStatements := GenerateInsertStatement(tableMapping.TableName, tableMapping.ColumnOrder, entities, tableMapping.Fields)
	for _, stmt := range insertStatements {
		sqlFile.WriteString(stmt + "\n")
	}
}

func normalizeOrderID(orderID string) string {
	// FIXME: Replace with proper normalization logic if needed
	return strings.ReplaceAll(orderID, "N", "")
}
