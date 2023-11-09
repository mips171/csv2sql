package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocarina/gocsv"
	"github.com/mips171/leo"
)

const (
	PRODUCTS_CSV        = "./data/products_export_full_20231101_114741_67572.csv"
	CUSTOMERS_CSV       = "./data/customer_export_full_20231101_114614_40007.csv"
	CUSTOMER_GROUPS_CSV = "./data/customer_groups.csv"
	ORDERS_CSV          = "./data/orders_export_full_20231101_114706_57298.csv"
	INFO_CSV            = "./data/content.csv"

	OUTPUT_CUSTOMERS_SQL  = "./data/import_customers.sql"
	OUTPUT_CATEGORIES_SQL = "./data/import_categories.sql"
	OUTPUT_ORDERS_SQL     = "./data/import_orders.sql"
	OUTPUT_PRODUCTS_SQL   = "./data/import_products.sql"
	OUTPUT_INFO_SQL       = "./data/import_info.sql"
)

func main() {

	// if input argument is -txt then just run the sku2txt function
	if len(os.Args) > 1 && os.Args[1] == "-txt" {
		sku2txt()
		return
	}

	// if input argument is -txt then just run the sku2txt function
	if len(os.Args) > 1 && os.Args[1] == "-img" {
		getImageURLs()
		return
	}

	graph := leo.TaskGraph()

	productTask := func() leo.TaskFunc {
		return func() error {
			products()
			return nil
		}
	}

	customerTask := func() leo.TaskFunc {
		return func() error {
			customers()
			return nil
		}
	}

	ordersTask := func() leo.TaskFunc {
		return func() error {
			orders()
			return nil
		}
	}

	categoriesTask := func() leo.TaskFunc {
		return func() error {
			categories()
			return nil
		}
	}

	infoTask := func() leo.TaskFunc {
		return func() error {
			information()
			return nil
		}
	}

	graph.Add("Products", productTask())
	graph.Add("Categories", categoriesTask())
	graph.Add("Orders", ordersTask())
	graph.Add("Customers", customerTask())
	graph.Add("Information", infoTask())

	graph.Succeed("Products", "Categories")

	executor := leo.NewExecutor(graph)

	if err := executor.Execute(); err != nil {
		fmt.Printf("Execution failed: %v\n", err)
	} else {
		fmt.Println("All tasks executed successfully.")
	}

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

	// Assume 'directory' is the path to the directory where images are stored
	directory := "catalog/products/"

	// Possible image file extensions
	extensions := []string{".jpg", ".png", ".jpeg"}

	for _, category := range cats {
		imgPath := "catalog/products/"
		for _, p := range products {
			if p.Category == category.Category {
				// Iterate over each file extension
				for _, ext := range extensions {
					imgPath = filepath.Join(directory, p.Model+ext)
					if _, err := os.Stat(imgPath); err == nil {
						// If the file exists, we've found our image
						break
					} else if os.IsNotExist(err) {
						// The file does not exist with this extension, try the next one
						continue
					} else {
						// Some other error occurred when checking the file
						// Handle this error accordingly
					}
				}
			}
		}

		for _, stmt := range GenerateCategorySQLStatements(category, imgPath) {
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
		orderID = normalizeOrderID(orderID)
		sqlStatements := GenerateOrderTotalSQLStatements(orderID, subTotalValue.StringFixed(2), shippingCost.StringFixed(2), taxValue.StringFixed(2), totalValue.StringFixed(2))
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

	sqlFile.WriteString("TRUNCATE TABLE `oc_product_image`;\n")
	// for every product, get its alt images
	for _, product := range products {
		altPaths := MapAltImageFilePaths(product)

		if len(altPaths) == 0 {
			fmt.Println("No alt images found for product:", product.Model)
		} else {
			fmt.Println("Found", len(altPaths), "alt images for product:", product.Model)
		}

		// for every alt image, generate a sql statement
		for _, path := range altPaths {
			sqlFile.WriteString(fmt.Sprintf("INSERT INTO `oc_product_image` (`product_id`, `image`, `sort_order`) VALUES ('%d', '%s', '0');\n", productIdMapping[product.Model], path))
		}
	}
}

func information() {
	productMapping := GetInformationMapping()

	// Open the file
	file, err := os.OpenFile(INFO_CSV, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Decode the CSV data
	var products []InformationRecord
	if err := gocsv.UnmarshalFile(file, &products); err != nil {
		fmt.Println("Error:", err)
		return
	}

	sqlFile, err := os.Create(OUTPUT_INFO_SQL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer sqlFile.Close()

	productIdMapping := make(map[string]int)
	for index, product := range products {
		productIdMapping[product.Name] = index + 1
	}

	entities := make([]Entity, len(products))
	for i, v := range products {
		entities[i] = v
	}

	// Use the helper function for each mapping
	processTable(productMapping, entities, productIdMapping, sqlFile)
	processTable(GetInformationDescriptionMapping(productIdMapping), entities, productIdMapping, sqlFile)
	processTable(GetInfoToStoreMapping(productIdMapping), entities, productIdMapping, sqlFile)
	processTable(GetInfoToLayoutMapping(productIdMapping), entities, productIdMapping, sqlFile)
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
	return strings.ReplaceAll(orderID, "N", "10")
}

func sku2txt() {

	PRODUCTS_TXT := "products.txt"
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

	// write all products to products.txt

	// Open the file
	file, err = os.OpenFile(PRODUCTS_TXT, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// write the SKU as plain text, one per line
	for _, product := range products {
		fmt.Fprintf(file, "%s\n", product.Model)
	}
}

func getImageURLs() {

	PRODUCTS_TXT := "emdedded_files.txt"
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

	// write all products to products.txt

	// Open the file
	file, err = os.OpenFile(PRODUCTS_TXT, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// write the imag path as plain text, one per line
	for _, product := range products {

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(product.Description))
		if err != nil {
			fmt.Printf("Error parsing the HTML: %v\n", err)
			return
		}

		// Find all links with href attribute
		doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
			href, exists := s.Attr("href")
			if exists && strings.HasSuffix(href, ".pdf") {
				fmt.Fprintf(file, "%s\n", href)
			}
		})

		// Find all image tags and extract the src attribute
		doc.Find("img[src]").Each(func(i int, s *goquery.Selection) {
			src, exists := s.Attr("src")
			if exists && (strings.HasSuffix(src, ".png") || strings.HasSuffix(src, ".jpg") || strings.HasSuffix(src, ".jpeg") || strings.HasSuffix(src, ".gif")) {
				fmt.Fprintf(file, "%s\n", src)
			}
		})

	}
}
