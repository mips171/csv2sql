package main

import (
	"strconv"
	"strings"
)

// This EntityRecord struct maps from CSV to a struct
type ProductRecord struct {
	Model              string `csv:"SKU*"`
	Name               string `csv:"Name"`
	Description        string `csv:"Description"`
	SEOPagetitle       string `csv:"SEO Page Title"`
	SEOMetaDescription string `csv:"SEO Meta Description"`
	SEOMetaKeywords    string `csv:"SEO Meta Keywords"`
	Price              string `csv:"Price (Retail)"`
	TradePrice         string `csv:"Price (Trade)"`
	Cost               string `csv:"Cost Price"`
	Quantity           string `csv:"Qty In Stock (Telco Antennas)"`
	Length             string `csv:"Length (Shipping)"`
	Width              string `csv:"Width (Shipping)"`
	Height             string `csv:"Height (Shipping)"`
	Weight             string `csv:"Weight (Shipping)"`
	TaxClassId         string `csv:"Tax Free Item"`
	DateAdded          string `csv:"Date Added"`
	DateModified       string `csv:"Date Modified"`
	Status             string `csv:"Approved"`
	UpSellProducts     string `csv:"Upsell Products"`
	CrossSellProducts  string `csv:"Cross-Sell Products"`
	// Add other fields as required
}

// Implement the Entity interface for ProductRecord
func (p ProductRecord) GetValue(fieldName string) interface{} {
	switch fieldName {
	case "Model":
		return p.Model
	case "Name":
		return p.Name
	case "Description":
		return p.Description
	case "SEOPagetitle":
		return p.SEOPagetitle
	case "SEOMetaDescription":
		return p.SEOMetaDescription
	case "SEOMetaKeywords":
		return p.SEOMetaKeywords
	case "Price":
		return p.Price
	case "TradePrice":
		return p.TradePrice
	case "Cost":
		return p.Cost
	case "Quantity":
		return p.Quantity
	case "Length":
		return p.Length
	case "Width":
		return p.Width
	case "Height":
		return p.Height
	case "Weight":
		return p.Weight
	case "TaxClassId":
		return p.TaxClassId
	case "DateAdded":
		return p.DateAdded
	case "DateModified":
		return p.DateModified
	case "Status":
		return p.Status
	case "UpsellProducts":
		return p.UpSellProducts
	case "CrossSellProducts":
		return p.CrossSellProducts
	// Add other fields as required
	default:
		return nil
	}
}

// This function maps from a ProductRecord to a map of strings
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
		"TradePrice":         product.TradePrice,
		"Cost":               product.Cost,
		"Quantity":           product.Quantity,
		"Length":             product.Length,
		"Width":              product.Width,
		"Height":             product.Height,
		"Weight":             product.Weight,
		"TaxClassId":         product.TaxClassId,
		"Status":             product.Status,
		"DateAdded":          product.DateAdded,
		"DateModified":       product.DateModified,
		"Image":              product.Model,
	}
}

// Map out our actual SQL for Product
func GetProductMapping() TableMapping {
	return TableMapping{
		TableName: "oc_product",
		ColumnOrder: []string{"model", "sku", "upc", "ean", "jan", "isbn", "mpn", "location", "quantity",
			"stock_status_id", "image", "manufacturer_id", "shipping", "price", "points", "tax_class_id", "date_available",
			"weight", "weight_class_id", "length", "width", "height", "length_class_id",
			"subtract", "minimum", "sort_order", "status", "viewed",
			"date_added", "date_modified"},
		Fields: []FieldMapping{
			{"Model", "model", DoNothing("Model")},
			{"Model", "sku", DoNothing("Model")},
			{"", "upc", ReturnEmptyString()},
			{"", "ean", ReturnEmptyString()},
			{"", "jan", ReturnEmptyString()},
			{"", "isbn", ReturnEmptyString()},
			{"", "mpn", ReturnEmptyString()},
			{"", "location", ReturnEmptyString()},
			{"Quantity", "quantity", DoNothing("Quantity")},
			{"", "stock_status_id", func(entity Entity) interface{} { return "7" }}, // Always 7 for "In Stock"
			{"Model", "image", MapImageFilePath},                                    // Using the Model to map the image file path
			{"", "manufacturer_id", func(entity Entity) interface{} { return "1" }}, // Always 1 for "Telco Antennas"
			{"", "shipping", func(entity Entity) interface{} { return "1" }},        // Always 1 for "Yes"
			{"Price", "price", GetRetailPrice},
			{"", "points", func(entity Entity) interface{} { return "0" }},
			{"TaxClassId", "tax_class_id", GetTaxClassID},
			{"", "date_available", func(entity Entity) interface{} { return "2023-09-20" }}, // example date
			{"Weight", "weight", DoNothing("Weight")},
			{"", "weight_class_id", func(entity Entity) interface{} { return "1" }}, // Always 1 for "Kilogram"
			{"Length", "length", DoNothing("Length")},
			{"Width", "width", DoNothing("Width")},
			{"Height", "height", DoNothing("Height")},
			{"", "length_class_id", func(entity Entity) interface{} { return "1" }}, // Always 1 for "Centimeter"
			{"", "subtract", func(entity Entity) interface{} { return "1" }},        // Always 1 for "Yes"
			{"", "minimum", func(entity Entity) interface{} { return "1" }},         // Always 1 for "Yes"
			{"", "sort_order", func(entity Entity) interface{} { return "1" }},      // Always 1 for "Yes"
			{"Status", "status", MapProductStatus},
			{"", "viewed", func(entity Entity) interface{} { return "0" }}, // Always 0 for "No"
			{"DateAdded", "date_added", func(entity Entity) interface{} { return GetDateAdded() }},
			{"DateModified", "date_modified", func(entity Entity) interface{} { return GetDateAdded() }},
		},
	}
}

// Map out our actual SQL for ProductDescription
func GetProductDescriptionMapping(productIdMapping map[string]int) TableMapping {
	return TableMapping{
		TableName:   "oc_product_description",
		ColumnOrder: []string{"product_id", "language_id", "name", "description", "tag", "meta_title", "meta_description", "meta_keyword"},
		Fields: []FieldMapping{
			{"Model", "product_id", GetProductIdTransformation(productIdMapping)},
			{"", "language_id", func(entity Entity) interface{} { return "1" }}, // Always 1 for English
			{"Name", "name", DoNothing("Name")},
			{"Description", "description", DoNothing("Description")},
			{"", "tag", func(entity Entity) interface{} { return "" }},
			{"Name", "meta_title", DoNothing("Name")},
			{"Name", "meta_description", DoNothing("Name")},
			{"Name", "meta_keyword", DoNothing("Name")},
		},
	}
}

// Map out our actual SQL for ProductToCategory
func GetProductRelatedMapping() TableMapping {
	return TableMapping{
		TableName:   "oc_product_related",
		ColumnOrder: []string{"product_id", "related_id"},
		Fields: []FieldMapping{
			{"UpsellProducts", "related_id", MapSkuToProductId},
			{"CrossSellProducts", "related_id", MapSkuToProductId},
		},
	}
}

// func GetProductToCategoryMapping() TableMapping {
// 	return TableMapping{
// 		TableName:   "oc_product_to_category",
// 		ColumnOrder: []string{"product_id", "category_id"},
// 		Fields: []FieldMapping{
// 			{"Category", "category_id", MapCategoryToCategoryId},
// 		},
// 	}
// }

// Map out our actual SQL for ProductSpecial
func GetProductSpecialMapping(productIdMapping map[string]int) TableMapping {
	return TableMapping{
		TableName:   "oc_product_special",
		ColumnOrder: []string{"product_id", "customer_group_id", "priority", "price", "date_start", "date_end"},
		Fields: []FieldMapping{
			{"Model", "product_id", func(entity Entity) interface{} {
				sku := entity.GetValue("Model").(string)
				return strconv.Itoa(productIdMapping[sku])
			}},
			{"", "customer_group_id", func(entity Entity) interface{} { return "2" }}, // Always 2 for Trade group
			{"", "priority", func(entity Entity) interface{} { return "0" }},
			{"TradePrice", "price", GetTradePrice},
			{"", "date_start", func(entity Entity) interface{} { return "0000-00-00" }}, // Always active start date
			{"", "date_end", func(entity Entity) interface{} { return "0000-00-00" }},   // Always active end date
		},
	}
}

// Map out our actual SQL for ProductCost
func GetProductToCostMapping(productIdMapping map[string]int) TableMapping {
	return TableMapping{
		TableName:   "oc_product_cost",
		ColumnOrder: []string{"product_id", "supplier_id", "cost", "cost_amount", "cost_percentage", "cost_additional", "costing_method"},
		Fields: []FieldMapping{
			{"Model", "product_id", GetProductIdTransformation(productIdMapping)},
			{"", "supplier_id", func(entity Entity) interface{} { return "0" }}, // Default store value
			{"Cost", "cost", DoNothing("Cost")},
			{"Cost", "cost_amount", DoNothing("Cost")},
			{"", "cost_percentage", func(entity Entity) interface{} { return "0.00" }},   // Default store value
			{"", "cost_additional", func(entity Entity) interface{} { return "0.0000" }}, // Default store value
			{"", "costing_method", func(entity Entity) interface{} { return "0" }},       // Default store value
		},
	}
}

// Map out our actual SQL for ProductToStore
func GetProductToStoreMapping(productIdMapping map[string]int) TableMapping {
	return TableMapping{
		TableName:   "oc_product_to_store",
		ColumnOrder: []string{"product_id", "store_id"},
		Fields: []FieldMapping{
			{"Model", "product_id", GetProductIdTransformation(productIdMapping)},
			{"", "store_id", func(entity Entity) interface{} { return "0" }}, // Default store value
		},
	}
}

var (
	categoryIDCounter       = 1
	categoryToCategoryIdMap = make(map[string]string)
)

func MapCategoryToCategoryId(categoryName string, nothing string) interface{} {
	if categoryId, ok := categoryToCategoryIdMap[categoryName]; ok {
		return categoryId
	}
	// Generate a new category_id and store the mapping
	categoryToCategoryIdMap[categoryName] = strconv.Itoa(categoryIDCounter)
	categoryIDCounter++
	return categoryToCategoryIdMap[categoryName]
}

var (
	skuToProductIdMap = make(map[string]int)
	productIDCounter  = 1
)

func MapSkuToProductId(entity Entity) interface{} {
	sku, _ := entity.GetValue("SKU").(string)
	if productId, ok := skuToProductIdMap[sku]; ok {
		return strconv.Itoa(productId)
	}
	// Generate a new product_id and store the mapping
	skuToProductIdMap[sku] = productIDCounter
	productIDCounter++
	return strconv.Itoa(skuToProductIdMap[sku])
}

func GetRetailPrice(entity Entity) interface{} {
	retailPrice := entity.GetValue("Price").(string)
	tradePrice := entity.GetValue("TradePrice").(string)

	// Default to retail price
	if retailPrice != "" {
		formattedRetailPrice := FormatPriceToFourDecimalPlaces(retailPrice)
		if formattedRetailPrice != "0.0000" {
			return formattedRetailPrice
		}
	}

	// Check and format trade price
	if tradePrice != "" {
		formattedTradePrice := FormatPriceToFourDecimalPlaces(tradePrice)
		if formattedTradePrice != "0.0000" {
			return formattedTradePrice
		}
	}

	// If neither trade nor retail prices are valid, return "0.0000"
	return "0.0000"
}

func GetTradePrice(entity Entity) interface{} {
	retailPrice := entity.GetValue("Price").(string)
	tradePrice := entity.GetValue("TradePrice").(string)

	// Check and format trade price
	if tradePrice != "" {
		formattedTradePrice := FormatPriceToFourDecimalPlaces(tradePrice)
		if formattedTradePrice != "0.0000" {
			return formattedTradePrice
		}
	}

	// Default to retail price
	if retailPrice != "" {
		formattedRetailPrice := FormatPriceToFourDecimalPlaces(retailPrice)
		if formattedRetailPrice != "0.0000" {
			return formattedRetailPrice
		}
	}

	// If neither trade nor retail prices are valid, return "0.0000"
	return "0.0000"
}

func FormatPriceToFourDecimalPlaces(price string) interface{} {
	// Split the price at the decimal
	parts := strings.Split(price, ".")

	// If there's no decimal
	if len(parts) == 1 {
		return price + ".0000"
	}

	// If there's a decimal, pad it out to four places
	for len(parts[1]) < 4 {
		parts[1] += "0"
	}

	return parts[0] + "." + parts[1][:4] // Return the formatted price
}

func GetTaxClassID(entity Entity) interface{} {
	taxFreeItem, _ := entity.GetValue("Tax Free Item").(string)

	// Check if the item is a tax-free item
	if taxFreeItem == "y" {
		// Return the ID for tax-free items
		return "11"
	}

	// Return the ID representing GST for all other items
	return "9"
}

func GetProductIdTransformation(productIdMapping map[string]int) func(entity Entity) interface{} {
	return func(entity Entity) interface{} {
		model := entity.GetValue("Model").(string)
		if id, exists := productIdMapping[model]; exists {
			return strconv.Itoa(id)
		}
		return nil
	}
}
