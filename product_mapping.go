package main

import (
	"strconv"
)

type ProductRecord struct {
	Model              string `csv:"SKU*"`
	Name               string `csv:"Name"`
	Description        string `csv:"Description"`
	SEOPagetitle       string `csv:"SEO Page Title"`
	SEOMetaDescription string `csv:"SEO Meta Description"`
	SEOMetaKeywords    string `csv:"SEO Meta Keywords"`
	Price              string `csv:"Price (Retail)"`
	Quantity           string `csv:"Qty In Stock (Telco Antennas)"`
	Length             string `csv:"Length (Shipping)"`
	Width              string `csv:"Width (Shipping)"`
	Height             string `csv:"Height (Shipping)"`
	Weight             string `csv:"Weight (Shipping)"`
	TaxClassId         string `csv:"Tax Free Item"`
	DateAdded          string `csv:"Date Added"`
	DateModified       string `csv:"Date Modified"`
	// Add other fields as required
}

func productToMap(product ProductRecord) map[string]string {
	return map[string]string{
		"Model":              product.Model,
		"SKU":                product.Model,
		"Name":               product.Name,
		"Description":        product.Description,
		"SEOPagetitle":       product.SEOPagetitle,
		"SEOMetaDescription": product.SEOMetaDescription,
		"SEOMetaKeywords":    product.SEOMetaKeywords,
		"Price":              product.Price,
		"Quantity":           product.Quantity,
		"Length":             product.Length,
		"Width":              product.Width,
		"Height":             product.Height,
		"Weight":             product.Weight,
		"TaxClassId":         product.TaxClassId,
		"DateAdded":          product.DateAdded,
		"DateModified":       product.DateModified,
		"Image":              product.Model,
	}
}

func GetProductMapping() TableMapping {
	return TableMapping{
		TableName:   "oc_product",
		ColumnOrder: []string{"model", "sku", "price", "quantity", "length", "width", "height", "weight", "tax_class_id", "date_added", "date_modified", "image"},
		Fields: []FieldMapping{
			{"Model", "model", DoNothing()},
			{"Model", "sku", DoNothing()},
			{"Price (Retail)", "price", func(value string, _ string) string { return GetRetailPrice(value, value) }},
			{"Qty In Stock (Telco Antennas)", "quantity", DoNothing()},
			{"Length (Shipping)", "length", DoNothing()},
			{"Width (Shipping)", "width", DoNothing()},
			{"Height (Shipping)", "height", DoNothing()},
			{"Weight (Shipping)", "weight", DoNothing()},
			{"Tax Free Item", "tax_class_id", func(value string, _ string) string { return GetTaxClassID(value) }},
			{"Date Added", "date_added", func(_ string, _ string) string { return GetDateAdded() }},
			{"Date Modified", "date_modified", func(_ string, _ string) string { return GetDateAdded() }},
			{"Model", "image", MapImageFilePath}, // Using the Model to map the image file path
		},
	}
}

func GetProductDescriptionMapping(productIdMapping map[string]int) TableMapping {
	return TableMapping{
		TableName:   "oc_product_description",
		ColumnOrder: []string{"product_id", "language_id", "name", "description", "meta_title", "meta_description", "meta_keyword"},
		Fields: []FieldMapping{
			{"Model", "product_id", func(sku string, _ string) string {
				return strconv.Itoa(productIdMapping[sku])
			}},
			{"", "language_id", func(_ string, _ string) string { return "1" }}, // Always 1 for English
			{"Name", "name", DoNothing()},
			{"Description", "description", DoNothing()},
			{"SEO Page Title", "meta_title", DoNothing()},
			{"SEO Meta Description", "meta_description", DoNothing()},
			{"SEO Meta Keywords", "meta_keyword", DoNothing()},
		},
	}
}

func GetProductRelatedMapping() TableMapping {
	return TableMapping{
		TableName:   "oc_product_related",
		ColumnOrder: []string{"product_id", "related_id"},
		Fields: []FieldMapping{
			{"Cross-Sell Products", "related_id", MapSkuToProductId},
			{"Upsell Products", "related_id", MapSkuToProductId},
		},
	}
}

func GetProductToCategoryMapping() TableMapping {
	return TableMapping{
		TableName:   "oc_product_to_category",
		ColumnOrder: []string{"product_id", "category_id"},
		Fields: []FieldMapping{
			{"Category", "category_id", MapCategoryToCategoryId},
		},
	}
}

func GetProductCostMapping() TableMapping {
	return TableMapping{
		TableName:   "oc_product_cost",
		ColumnOrder: []string{"product_id", "cost"},
		Fields: []FieldMapping{
			{"Cost Price", "cost", DoNothing()},
		},
	}
}

var categoryIDCounter = 1
var categoryToCategoryIdMap = make(map[string]string)

func MapCategoryToCategoryId(categoryName string, nothing string) string {
	if categoryId, ok := categoryToCategoryIdMap[categoryName]; ok {
		return categoryId
	}
	// Generate a new category_id and store the mapping
	categoryToCategoryIdMap[categoryName] = strconv.Itoa(categoryIDCounter)
	categoryIDCounter++
	return categoryToCategoryIdMap[categoryName]
}

var productIDCounter = 1
var skuToProductIdMap = make(map[string]int)

func MapSkuToProductId(sku string, nothing string) string {
	if productId, ok := skuToProductIdMap[sku]; ok {
		return strconv.Itoa(productId)
	}
	// Generate a new product_id and store the mapping
	skuToProductIdMap[sku] = productIDCounter
	productIDCounter++
	return strconv.Itoa(skuToProductIdMap[sku])
}

func GetRetailPrice(retailPrice, tradePrice string) string {
	if retailPrice != "" {
		return retailPrice
	}

	// If trade price exists, set both to their respective values
	if tradePrice != "" {
		return tradePrice
	}
	// If trade price does not exist, set the default price to the retail price
	return "0.00"
}

func GetTradePrice(retailPrice, tradePrice string) string {
	if tradePrice != "" {
		return tradePrice
	}

	// If trade price exists, set both to their respective values
	if retailPrice != "" {
		return retailPrice
	}
	// If trade price does not exist, set the default price to the retail price
	return "0.00"
}

func GetTaxClassID(taxFreeItem string) string {
	// Check if the item is a tax-free item
	if taxFreeItem == "y" {
		// Return the ID for tax-free items
		return "11"
	}

	// Return the ID representing GST for all other items
	return "9"
}

func GetProductSpecialMapping(productIdMapping map[string]int) TableMapping {
	return TableMapping{
		TableName:   "oc_product_special",
		ColumnOrder: []string{"product_id", "customer_group_id", "priority", "price", "date_start", "date_end"},
		Fields: []FieldMapping{
			{"Model", "product_id", func(sku string, _ string) string {
				return strconv.Itoa(productIdMapping[sku])
			}},
			{"Price (Trade)", "price", func(value string, _ string) string {
				if value != "" {
					return value
				}
				return ""
			}},
			{"", "customer_group_id", func(_ string, _ string) string { return "2" }},   // Always 2 for Trade group
			{"", "priority", func(_ string, _ string) string { return "0" }},            // Default priority value
			{"", "date_start", func(_ string, _ string) string { return "0000-00-00" }}, // Always active start date
			{"", "date_end", func(_ string, _ string) string { return "0000-00-00" }},   // Always active end date
		},
	}
}
