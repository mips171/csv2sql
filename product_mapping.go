package main

import (
	"strconv"
	"strings"
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

// INSERT INTO `oc_product` (`model`, `sku`, `upc`, `ean`, `jan`, `isbn`, `mpn`, `location`, `quantity`,
// `stock_status_id`, `image`, `manufacturer_id`, `shipping`, `price`, `points`, `tax_class_id`, `date_available`,
// `weight`, `weight_class_id`, `length`, `width`, `height`,  `length_class_id`,
// `subtract`, `minimum`, `sort_order`, `status`, `viewed`,
//`date_added`, `date_modified`) VALUES ('1', 'TESTPRODUCT', '', '', '', '', '', '', '', '0', '6', '', '0', '1', '0.0000', '0', '0', '2023-09-19', '0.00000000', '1', '0.00000000', '0.00000000', '0.00000000', '1', '1', '1', '1', '1', '0', '2023-09-19 14:05:23', '2023-09-19 14:05:23');

func GetProductMapping() TableMapping {
	return TableMapping{
		TableName: "oc_product",
		ColumnOrder: []string{"model", "sku", "upc", "ean", "jan", "isbn", "mpn", "location", "quantity",
			"stock_status_id", "image", "manufacturer_id", "shipping", "price", "points", "tax_class_id", "date_available",
			"weight", "weight_class_id", "length", "width", "height", "length_class_id",
			"subtract", "minimum", "sort_order", "status", "viewed",
			"date_added", "date_modified"},
		Fields: []FieldMapping{
			{"Model", "model", DoNothing()},
			{"Model", "sku", DoNothing()},
			{"", "upc", DoNothing()},
			{"", "ean", DoNothing()},
			{"", "jan", DoNothing()},
			{"", "isbn", DoNothing()},
			{"", "mpn", DoNothing()},
			{"", "location", DoNothing()},
			{"Qty In Stock (Telco Antennas)", "quantity", DoNothing()},
			{"", "stock_status_id", func(_ string, _ string) interface{} { return "7" }}, // Always 7 for "In Stock"
			{"Model", "image", MapImageFilePath},                                         // Using the Model to map the image file path
			{"", "manufacturer_id", func(_ string, _ string) interface{} { return "1" }}, // Always 1 for "Telco Antennas"
			{"", "shipping", func(_ string, _ string) interface{} { return "1" }},        // Always 1 for "Yes"
			{"Price", "price", func(value string, _ string) interface{} { return GetRetailPrice(value, value) }},
			{"", "points", DoNothing()},
			{"Tax Free Item", "tax_class_id", func(value string, _ string) interface{} { return GetTaxClassID(value) }},
			{"", "date_available", func(_ string, _ string) interface{} { return "" }},

			{"Weight (Shipping)", "weight", DoNothing()},
			{"", "weight_class_id", func(_ string, _ string) interface{} { return "1" }}, // Always 1 for "Kilogram"
			{"Length (Shipping)", "length", DoNothing()},
			{"Width (Shipping)", "width", DoNothing()},
			{"Height (Shipping)", "height", DoNothing()},

			{"", "length_class_id", func(_ string, _ string) interface{} { return "1" }}, // Always 1 for "Centimeter"
			{"", "subtract", func(_ string, _ string) interface{} { return "1" }},        // Always 1 for "Yes"
			{"", "minimum", func(_ string, _ string) interface{} { return "1" }},         // Always 1 for "Yes"
			{"", "sort_order", func(_ string, _ string) interface{} { return "1" }},      // Always 1 for "Yes"
			{"Approved", "status", func(value string, _ string) interface{} { return MapProductStatus(value, value) }},
			{"", "viewed", func(_ string, _ string) interface{} { return "0" }}, // Always 0 for "No"
			{"Date Added", "date_added", func(_ string, _ string) interface{} { return GetDateAdded() }},
			{"Date Modified", "date_modified", func(_ string, _ string) interface{} { return GetDateAdded() }},
		},
	}
}

func GetProductDescriptionMapping(productIdMapping map[string]int) TableMapping {
	return TableMapping{
		TableName:   "oc_product_description",
		ColumnOrder: []string{"product_id", "language_id", "name", "description", "tag", "meta_title", "meta_description", "meta_keyword"},
		Fields: []FieldMapping{
			{"Model", "product_id", func(sku string, _ string) interface{} {
				return strconv.Itoa(productIdMapping[sku])
			}},
			{"", "language_id", func(_ string, _ string) interface{} { return "1" }}, // Always 1 for English
			{"Name", "name", DoNothing()},
			{"Description", "description", DoNothing()},
			{"", "tag", func(_ string, _ string) interface{} { return "" }},
			{"Name", "meta_title", DoNothing()},
			{"Name", "meta_description", DoNothing()},
			{"Name", "meta_keyword", DoNothing()},
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

// need product cost
// INSERT INTO `oc_product_to_store` (`product_id`, `store_id`) VALUES ('5', '0');

func GetProductToStoreMapping(productIdMapping map[string]int) TableMapping {
	return TableMapping{
		TableName:   "oc_product_to_store",
		ColumnOrder: []string{"product_id", "store_id"},
		Fields: []FieldMapping{
			{"Model", "product_id", func(sku string, _ string) interface{} {
				return strconv.Itoa(productIdMapping[sku])
			}},
			{"", "store_id", func(_ string, _ string) interface{} { return "0" }}, // Default store value
		},
	}
}

// INSERT INTO `oc_product_cost` (`product_cost_id`, `product_id`, `supplier_id`, `cost`, `cost_amount`, `cost_percentage`, `cost_additional`, `costing_method`)
// VALUES ('4', '1', '0', '0.0000', '0.0000', '0.00', '0.0000', '0');
func GetProductToCostMapping(productIdMapping map[string]int) TableMapping {
	return TableMapping{
		TableName:   "oc_product_cost",
		ColumnOrder: []string{"product_id", "supplier_id", "cost", "cost_amount", "cost_percentage", "cost_additional", "costing_method"},
		Fields: []FieldMapping{
			{"Model", "product_id", func(sku string, _ string) interface{} {
				return strconv.Itoa(productIdMapping[sku])
			}},
			{"", "supplier_id", func(_ string, _ string) interface{} { return "0" }},                                // Default store value
			{"Cost Price", "cost", func(value string, _ string) interface{} { return GetTradePrice(value, value) }}, // reusing getretailprice
			{"", "cost_amount", func(_ string, _ string) interface{} { return "0.0000" }},                           // Default store value
			{"", "cost_percentage", func(_ string, _ string) interface{} { return "0.00" }},                         // Default store value
			{"", "cost_additional", func(_ string, _ string) interface{} { return "0.0000" }},                       // Default store value
			{"", "costing_method", func(_ string, _ string) interface{} { return "0" }},                             // Default store value
		},
	}
}

var categoryIDCounter = 1
var categoryToCategoryIdMap = make(map[string]string)

func MapCategoryToCategoryId(categoryName string, nothing string) interface{} {
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

func MapSkuToProductId(sku string, nothing string) interface{} {
	if productId, ok := skuToProductIdMap[sku]; ok {
		return strconv.Itoa(productId)
	}
	// Generate a new product_id and store the mapping
	skuToProductIdMap[sku] = productIDCounter
	productIDCounter++
	return strconv.Itoa(skuToProductIdMap[sku])
}

func GetRetailPrice(retailPrice, tradePrice string) interface{} {
	formattedRetailPrice := FormatPriceToFourDecimalPlaces(retailPrice)

	if formattedRetailPrice != "0.0000" {
		return formattedRetailPrice
	}

	formattedTradePrice := FormatPriceToFourDecimalPlaces(tradePrice)
	if formattedTradePrice != "0.0000" {
		return formattedTradePrice
	}

	return "0.0000"
}

func GetTradePrice(retailPrice, tradePrice string) interface{} {
	if tradePrice != "" {
		formattedTradePrice := FormatPriceToFourDecimalPlaces(tradePrice)
		if formattedTradePrice != "0.0000" {
			return formattedTradePrice
		}
	}

	// If trade price exists, set both to their respective values
	if retailPrice != "" {
		formattedRetailPrice := FormatPriceToFourDecimalPlaces(retailPrice)

		if formattedRetailPrice != "0.0000" {
			return formattedRetailPrice
		}
	}
	// If trade price does not exist, set the default price to the retail price
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

func GetTaxClassID(taxFreeItem string) interface{} {
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
			{"Model", "product_id", func(sku string, _ string) interface{} {
				return strconv.Itoa(productIdMapping[sku])
			}},
			{"", "customer_group_id", func(_ string, _ string) interface{} { return "2" }}, // Always 2 for Trade group
			{"", "priority", func(_ string, _ string) interface{} { return "0" }},
			{"Price (Trade)", "price", func(value string, _ string) interface{} { return GetTradePrice(value, value) }},
			{"", "date_start", func(_ string, _ string) interface{} { return "0000-00-00" }}, // Always active start date
			{"", "date_end", func(_ string, _ string) interface{} { return "0000-00-00" }},   // Always active end date
		},
	}
}
