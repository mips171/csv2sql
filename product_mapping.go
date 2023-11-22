package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
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
	Weight             string `csv:"Weight (shipping)"` // be sure to check sneaky neto who calls it Weigh (shipping) note the lowercase s
	TaxClassId         string `csv:"Tax Free Item"`
	DateAdded          string `csv:"Date Added"`
	DateModified       string `csv:"Date Modified"`
	Status             string `csv:"Approved"`
	UpSellProducts     string `csv:"Upsell Products"`
	CrossSellProducts  string `csv:"Cross-Sell Products"`
	Category           string `csv:"Category"` // accessed by category mapping
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
	case "Category":
		return p.Category
	// Add other fields as required
	default:
		return ""
	}
}

// Map out our actual SQL for Product
func GetProductMapping() TableMapping {
	return TableMapping{
		TableName: "oc_product",
		ColumnOrder: []string{"model", "sku", "location", "quantity",
			"stock_status_id", "image", "manufacturer_id", "shipping", "price", "tax_class_id", "date_available",
			"weight", "weight_class_id", "length", "width", "height", "length_class_id",
			"subtract", "minimum", "sort_order", "status",
			"date_added", "date_modified"},
		Fields: []FieldMapping{
			{"Model", "model", ToUpperCase("Model")},
			{"Model", "sku", ToUpperCase("Model")},
			{"", "location", EmptyString()},
			{"Quantity", "quantity", JustUse("Quantity")},

			{"", "stock_status_id", func(entity Entity) interface{} { return "7" }}, // Always 7 for "In Stock"
			{"Model", "image", MapImageFilePath},                                    // Using the Model to map the image file path
			{"", "manufacturer_id", func(entity Entity) interface{} { return "1" }}, // Always 1 for "Telco Antennas"
			{"", "shipping", func(entity Entity) interface{} { return "1" }},        // Always 1 for "Yes"
			{"Price", "price", GetRetailPrice},
			{"TaxClassId", "tax_class_id", GetTaxClassID},
			{"", "date_available", func(entity Entity) interface{} { return "2023-09-20" }}, // example date

			{"Weight", "weight", JustUse("Weight")},
			{"", "weight_class_id", func(entity Entity) interface{} { return "1" }}, // Always 1 for "Kilogram"
			{"Length", "length", JustUse("Length")},
			{"Width", "width", JustUse("Width")},
			{"Height", "height", JustUse("Height")},
			{"", "length_class_id", func(entity Entity) interface{} { return "4" }}, // Always 4 for "Meter"

			{"", "subtract", func(entity Entity) interface{} { return "1" }},   // Always 1 for "Yes"
			{"", "minimum", func(entity Entity) interface{} { return "1" }},    // Always 1 for "Yes"
			{"", "sort_order", func(entity Entity) interface{} { return "1" }}, // Always 1 for "Yes"
			{"Status", "status", MapProductStatus},

			{"DateAdded", "date_added", func(entity Entity) interface{} { return "2010-02-03 16:59:00" }},
			{"DateModified", "date_modified", func(entity Entity) interface{} { return "2010-02-03 16:59:00" }},
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
			{"Name", "name", JustUse("Name")},
			{"Description", "description", MapDescriptionURLs},
			{"", "tag", func(entity Entity) interface{} { return "" }},
			{"Name", "meta_title", JustUse("Name")},
			{"Name", "meta_description", JustUse("Name")},
			{"Name", "meta_keyword", JustUse("Name")},
		},
	}
}

// Update any URLs in the description to use the new paths
func MapDescriptionURLs(entity Entity) interface{} {
	description := entity.GetValue("Description").(string)

	// Replace any old URLs with the new ones
	description = strings.Replace(description, "assets/imported/site/sites/default/files", "image/catalog", -1)
	description = strings.Replace(description, "//www.telcoantennas.com.au", "", -1)

	return description
}

// Map out our actual SQL for ProductToRelated
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

// Map out our actual SQL for ProductSpecial
func GetProductDiscountMapping(productIdMapping map[string]int) TableMapping {
	return TableMapping{
		TableName:   "oc_product_discount",
		ColumnOrder: []string{"product_id", "customer_group_id", "quantity", "priority", "price", "date_start", "date_end"},
		Fields: []FieldMapping{
			{"Model", "product_id", func(entity Entity) interface{} {
				sku := entity.GetValue("Model").(string)
				return strconv.Itoa(productIdMapping[sku])
			}},
			{"", "customer_group_id", func(entity Entity) interface{} { return "2" }}, // Always 2 for Trade group
			{"", "quantity", func(entity Entity) interface{} { return "1" }},          // at least 1 or more
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
			{"Cost", "cost", JustUse("Cost")},
			{"Cost", "cost_amount", JustUse("Cost")},
			{"", "cost_percentage", func(entity Entity) interface{} { return "0.00" }}, // Default store value
			{"", "cost_additional", func(entity Entity) interface{} { return "0.00" }}, // Default store value
			{"", "costing_method", func(entity Entity) interface{} { return "0" }},     // Default store value
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

// INSERT INTO `oc_product_attribute` (`product_id`, `attribute_id`, `language_id`, `text`) VALUES
// (43, 2, 1, '1'),
// (47, 4, 1, '16GB'),
// (43, 4, 1, '8gb'),
// (42, 3, 1, '100mhz'),
// (47, 2, 1, '4');

// Create and maintain a map of SKU to product_id
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
	retailPrice, ok := entity.GetValue("Price").(string)
	if !ok || retailPrice == "" {
		return "0.00"
	}

	return removeGSTandFormat(retailPrice)
}

func removeGSTandFormat(price string) string {
	priceDec, err := decimal.NewFromString(price)
	if err != nil {
		return "0.00" // Return default value if conversion fails
	}

	// Reverse the 10% GST
	gstRate := decimal.NewFromFloat(1.10) // GST is 10%, so we divide by 1 + 0.10
	exGSTPrice := priceDec.DivRound(gstRate, 2)

	// Format to string with four decimal places
	return exGSTPrice.StringFixed(2)
}

func GetTradePrice(entity Entity) interface{} {
	tradePrice, okTrade := entity.GetValue("TradePrice").(string)
	retailPrice, okRetail := entity.GetValue("Price").(string)

	if okTrade && tradePrice != "" {
		formattedTradePrice := removeGSTandFormat(tradePrice)
		if formattedTradePrice != "0.00" {
			return formattedTradePrice
		}
	}

	if okRetail && retailPrice != "" {
		formattedRetailPrice := removeGSTandFormat(retailPrice)
		if formattedRetailPrice != "0.00" {
			return formattedRetailPrice
		}
	}

	return "0.0000"
}

func GetTaxClassID(entity Entity) interface{} {
	taxFreeItem, _ := entity.GetValue("TaxFreeItem").(string)

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

func GetCategoryId(productCategory string) int {
	return categoryToCategoryIdMap[productCategory]
}

func GenerateProductToCategorySQLStatements(product ProductRecord) []string {
	// Parse the product's category.
	categoryName := product.Category

	// Get the corresponding category ID.
	categoryID := GetCategoryId(categoryName)
	productID := MapSkuToProductId(product)

	var statements []string

	// Create the SQL statement linking the product to the category.
	statements = append(statements, fmt.Sprintf("INSERT IGNORE INTO `oc_product_to_category` (`product_id`, `category_id`) VALUES (%d, %d);", productID, categoryID))

	return statements
}
