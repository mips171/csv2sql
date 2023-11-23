package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

// CSV Headers, need to match OrderRecord deserialization
// "Order ID","Order Status","Approved","Username","Email","Ship First Name","Ship Last Name","Ship Company","Ship Address Line 1","Ship Address Line 2","Ship City","Ship State","Ship Post Code","Ship Country","Ship Phone","Ship Fax","RUT950 Serial Number","RUT950 IMEI","Transaction Number","Misc 0","Misc Notes 1","Installation Service Notes","Misc Notes 2","Bill First Name","Bill Last Name","Bill Company","Bill Address Line 1","Bill Address Line 2","Bill City","Bill State","Bill Post Code","Bill Country","Bill Phone","Bill Fax","Payment Method","Shipping Method","Customer Instructions","Internal Notes","Amount Paid","Date Paid","Order Line SKU","Order Line Qty","Order Line Description","Order Line Unit Price","Order Line Unit Cost","Tax Free Shipping","Card Holder","Shipping Discount Amount","Order Line Serial Number","Order Payment Plan","Parent Order ID","Payment Due Date","User Group","Fraud Score","BPAY CRN","Order Line Bin Location","Sales Channel","Order Line Options","Order Line Dropship Supplier","Order Line Tax Free","Order Line Discount Amount","Order Line Shipping Cubic","Order Line Job","Tax Inclusive","Purchase Order ID","Order Type","Date Required","Order Line Shipping Weight","Order Line Notes","Sales Person","Currency Code","Date Invoiced","Shipping Cost","Order Line Delivery Date","Payment Terms","Date Placed","Document Template","Coupon Code"
// What we will actually import
// INSERT INTO `oc_order` (`order_id`, `invoice_no`, `invoice_prefix`, `store_id`, `store_name`, `store_url`, `customer_id`, `customer_group_id`, `firstname`, `lastname`, `email`, `telephone`, `fax`, `custom_field`, `payment_firstname`, `payment_lastname`, `payment_company`, `payment_address_1`, `payment_address_2`, `payment_city`, `payment_postcode`, `payment_country`, `payment_country_id`, `payment_zone`, `payment_zone_id`, `payment_address_format`, `payment_custom_field`, `payment_method`, `payment_code`, `shipping_firstname`, `shipping_lastname`, `shipping_company`, `shipping_address_1`, `shipping_address_2`, `shipping_city`, `shipping_postcode`, `shipping_country`, `shipping_country_id`, `shipping_zone`, `shipping_zone_id`, `shipping_address_format`, `shipping_custom_field`, `shipping_method`, `shipping_code`, `comment`, `total`, `order_status_id`, `affiliate_id`, `commission`, `marketing_id`, `tracking`, `language_id`, `currency_id`, `currency_code`, `currency_value`, `ip`, `forwarded_ip`, `user_agent`, `accept_language`, `date_added`, `date_modified`) VALUES ('2', '0', ”, '0', 'Telco Antennas', 'https://telcoshop.nbembedded.com/', '9', '0', 'Daryl', 'Sowinski', 'dsowinski@bigpond.com', '419653781', ”, ”, 'Daryl', 'Sowinski', ”, '37 Diamantina Street', ”, 'Hillcrest', '4118', 'Australia', '13', ”, '0', ”, ”, 'Direct Deposit (EFT)', 'cod', 'Daryl', 'Sowinski', ”, '37 Diamantina Street', ”, 'Hillcrest', '4118', 'Australia', '13', ”, '0', ”, ”, 'Default shipping.', ”, ”, '46.8900', '7', '0', '0.0000', '0', ”, '1', '0', 'AUD', '1.00000000', ”, ”, ”, ”, '2011-02-22 03:25:29', '2018-07-02 02:07:06');
// We want these two to match up, and ignore everything else.
type OrderRecord struct {
	OrderID     string `csv:"Order ID"`
	OrderStatus string `csv:"Order Status"`
	Approved    string `csv:"Approved"`

	Email     string `csv:"Email"`
	Telephone string `csv:"Ship Phone"`
	Fax       string `csv:"Ship Fax"`

	Firstname       string `csv:"Bill First Name"`
	Lastname        string `csv:"Bill Last Name"`
	PaymentCompany  string `csv:"Bill Company"`
	PaymentAddress1 string `csv:"Bill Address Line 1"`
	PaymentAddress2 string `csv:"Bill Address Line 2"`
	PaymentCity     string `csv:"Bill City"`
	PaymentState    string `csv:"Bill State"`
	PaymentPostcode string `csv:"Bill Post Code"`
	PaymentCountry  string `csv:"Bill Country"`

	ShippingFirstname string `csv:"Ship First Name"`
	ShippingLastname  string `csv:"Ship Last Name"`
	ShippingCompany   string `csv:"Ship Company"`
	ShippingAddress1  string `csv:"Ship Address Line 1"`
	ShippingAddress2  string `csv:"Ship Address Line 2"`
	ShippingCity      string `csv:"Ship City"`
	ShippingState     string `csv:"Ship State"`
	ShippingPostcode  string `csv:"Ship Post Code"`
	ShippingCountry   string `csv:"Ship Country"`

	PaymentMethod        string `csv:"Payment Method"`
	ShippingMethod       string `csv:"Shipping Method"`
	ShippingCost         string `csv:"Shipping Cost"`
	Total                string
	PaymentCode          string // this will need to be mapped based on PaymentMethod, not present in CSV
	ShippingCode         string // this will need to be mapped based on ShippingMethod, not present in CSV
	CurrencyCode         string `csv:"Currency Code"`
	DateAdded            string `csv:"Date Placed"`
	DateModified         string `csv:"Date Invoiced"`
	DatePaymentDue       string `csv:"Payment Due Date"`
	OrderLineTaxFree     string `csv:"Order Line Tax Free"`
	OrderLineSKU         string `csv:"Order Line SKU"`
	OrderLineQty         string `csv:"Order Line Qty"`
	OrderLineDescription string `csv:"Order Line Description"`
	OrderLineUnitPrice   string `csv:"Order Line Unit Price"`
	AmountPaid           string `csv:"Amount Paid"`
	// Add any other relevant fields as needed.

	// Comments
	RUT950SerialNumber   string `csv:"RUT950 Serial Number"`
	RUT950IMEI           string `csv:"RUT950 IMEI"`
	TransactionNumber    string `csv:"Transaction Number"`
	Misc0                string `csv:"Misc 0"`
	MiscNotes1           string `csv:"Misc Notes 1"`
	InstallationNotes    string `csv:"Installation Service Notes"`
	MiscNotes2           string `csv:"Misc Notes 2"`
	CustomerInstructions string `csv:"Customer Instructions"`
	InternalNotes        string `csv:"Internal Notes"`
}

func (o OrderRecord) GetValue(fieldName string) interface{} {
	switch fieldName {
	case "OrderID":
		return o.OrderID
	case "OrderStatus":
		return o.OrderStatus
	case "Approved":
		return o.Approved
	case "Email":
		return o.Email
	case "Telephone":
		return o.Telephone
	case "Fax":
		return o.Fax
	case "PaymentFirstname":
		return o.Firstname
	case "PaymentLastname":
		return o.Lastname
	case "PaymentCompany":
		return o.PaymentCompany
	case "PaymentAddress1":
		return o.PaymentAddress1
	case "PaymentAddress2":
		return o.PaymentAddress2
	case "PaymentCity":
		return o.PaymentCity
	case "PaymentState":
		return o.PaymentState
	case "PaymentPostcode":
		return o.PaymentPostcode
	case "PaymentCountry":
		return o.PaymentCountry
	case "ShippingFirstname":
		return o.ShippingFirstname
	case "ShippingLastname":
		return o.ShippingLastname
	case "ShippingCompany":
		return o.ShippingCompany
	case "ShippingAddress1":
		return o.ShippingAddress1
	case "ShippingAddress2":
		return o.ShippingAddress2
	case "ShippingCity":
		return o.ShippingCity
	case "ShippingState":
		return o.ShippingState
	case "ShippingPostcode":
		return o.ShippingPostcode
	case "ShippingCountry":
		return o.ShippingCountry
	case "PaymentMethod":
		return o.PaymentMethod
	case "ShippingMethod":
		return o.ShippingMethod
	case "ShippingCost":
		return o.ShippingCost
	case "Total":
		return o.Total
	case "PaymentCode":
		return o.PaymentCode
	case "ShippingCode":
		return o.ShippingCode
	case "CurrencyCode":
		return o.CurrencyCode
	case "DateAdded":
		return o.DateAdded
	case "DateModified":
		return o.DateModified
	case "DatePaymentDue":
		return o.DatePaymentDue
	case "OrderLineSKU":
		return o.OrderLineSKU
	case "OrderLineQty":
		return o.OrderLineQty
	case "OrderLineDescription":
		return o.OrderLineDescription
	case "OrderLineUnitPrice":
		return o.OrderLineUnitPrice
	// ... add other fields as required from your CSV
	case "AllComments":
		// conditionally build the string to avoid blanks
		var comments []string
		if o.RUT950SerialNumber != "" {
			comments = append(comments, fmt.Sprintf("RUT950 Serial Number: %s", o.RUT950SerialNumber))
		}
		if o.RUT950IMEI != "" {
			comments = append(comments, fmt.Sprintf("RUT950 IMEI: %s", o.RUT950IMEI))
		}
		if o.TransactionNumber != "" {
			comments = append(comments, fmt.Sprintf("Transaction Number: %s", o.TransactionNumber))
		}
		if o.Misc0 != "" {
			comments = append(comments, fmt.Sprintf("Misc 0: %s", o.Misc0))
		}
		if o.MiscNotes1 != "" {
			comments = append(comments, fmt.Sprintf("Misc Notes 1: %s", o.MiscNotes1))
		}
		if o.InstallationNotes != "" {
			comments = append(comments, fmt.Sprintf("Installation Service Notes: %s", o.InstallationNotes))
		}
		if o.MiscNotes2 != "" {
			comments = append(comments, fmt.Sprintf("Misc Notes 2: %s", o.MiscNotes2))
		}
		if o.CustomerInstructions != "" {
			comments = append(comments, fmt.Sprintf("Customer Instructions: %s", o.CustomerInstructions))
		}
		if o.InternalNotes != "" {
			comments = append(comments, fmt.Sprintf("Internal Notes: %s", o.InternalNotes))
		}
		return strings.Join(comments, "\n")
	default:
		return ""
	}
}

// Need to map to this schema
// INSERT INTO `oc_order` (`order_id`, `invoice_no`, `invoice_prefix`, `store_id`, `store_name`, `store_url`, `customer_id`, `customer_group_id`, `firstname`, `lastname`, `email`, `telephone`, `fax`, `custom_field`, `payment_firstname`, `payment_lastname`, `payment_company`, `payment_address_1`, `payment_address_2`, `payment_city`, `payment_postcode`, `payment_country`, `payment_country_id`, `payment_zone`, `payment_zone_id`, `payment_address_format`, `payment_custom_field`, `payment_method`, `payment_code`, `shipping_firstname`, `shipping_lastname`, `shipping_company`, `shipping_address_1`, `shipping_address_2`, `shipping_city`, `shipping_postcode`, `shipping_country`, `shipping_country_id`, `shipping_zone`, `shipping_zone_id`, `shipping_address_format`, `shipping_custom_field`, `shipping_method`, `shipping_code`, `comment`, `total`, `order_status_id`, `affiliate_id`, `commission`, `marketing_id`, `tracking`, `language_id`, `currency_id`, `currency_code`, `currency_value`, `ip`, `forwarded_ip`, `user_agent`, `accept_language`, `date_added`, `date_modified`);

func GetOrderMapping(customerEmailMapping map[string]int) TableMapping {
	return TableMapping{
		TableName:   "oc_order",
		ColumnOrder: []string{"order_id", "invoice_no", "store_id", "customer_id", "firstname", "lastname", "email", "telephone", "payment_firstname", "payment_lastname", "payment_company", "payment_address_1", "payment_address_2", "payment_city", "payment_postcode", "payment_zone", "payment_zone_id", "payment_country", "payment_country_id", "payment_method", "payment_code", "shipping_firstname", "shipping_lastname", "shipping_company", "shipping_address_1", "shipping_address_2", "shipping_city", "shipping_postcode", "shipping_zone", "shipping_zone_id", "shipping_country", "shipping_country_id", "shipping_method", "shipping_code", "comment", "total", "order_status_id", "date_added", "date_modified", "date_payment_due", "currency_id", "currency_code", "currency_value"},
		Fields: []FieldMapping{
			// No need for order_id since it's managed by the database.
			// As an example, here are a few more mappings:
			{"", "order_id", StripNPrefix},
			{"InvoiceNo", "invoice_no", JustUse("OrderID")},
			{"", "store_id", func(entity Entity) interface{} { return "0" }}, // always use 0
			{"CustomerID", "customer_id", MapCustomerEmailToID(customerEmailMapping)},
			{"Firstname", "firstname", JustUse("PaymentFirstname")},
			{"Lastname", "lastname", JustUse("PaymentLastname")},
			{"Email", "email", JustUse("Email")},
			{"Telephone", "telephone", JustUse("Telephone")},
			{"PaymentFirstname", "payment_firstname", JustUse("PaymentFirstname")},
			{"PaymentLastname", "payment_lastname", JustUse("PaymentLastname")},
			{"PaymentCompany", "payment_company", JustUse("PaymentCompany")},
			{"PaymentAddress1", "payment_address_1", JustUse("PaymentAddress1")},
			{"PaymentAddress2", "payment_address_2", JustUse("PaymentAddress2")},
			{"PaymentCity", "payment_city", JustUse("PaymentCity")},
			{"", "payment_zone", JustUse("PaymentState")},                                  // derive state from postcode
			{"", "payment_zone_id", MapAustralianPostCodeToStateZoneID("PaymentPostcode")}, // derive state from postcode
			{"PaymentPostcode", "payment_postcode", JustUse("PaymentPostcode")},
			{"PaymentCountry", "payment_country", JustUse("PaymentCountry")},
			{"", "payment_country_id", MapCountryToCode("PaymentCountry")},
			{"PaymentMethod", "payment_method", JustUse("PaymentMethod")},
			{"PaymentCode", "payment_code", func(entity Entity) interface{} { return "cod" }}, // TODO: change this to a mapping function
			{"ShippingFirstname", "shipping_firstname", JustUse("ShippingFirstname")},
			{"ShippingLastname", "shipping_lastname", JustUse("ShippingLastname")},
			{"ShippingCompany", "shipping_company", JustUse("ShippingCompany")},
			{"ShippingAddress1", "shipping_address_1", JustUse("ShippingAddress1")},
			{"ShippingAddress2", "shipping_address_2", JustUse("ShippingAddress2")},
			{"ShippingCity", "shipping_city", JustUse("ShippingCity")},
			{"", "shipping_zone", JustUse("ShippingState")},                                  // derive state from postcode
			{"", "shipping_zone_id", MapAustralianPostCodeToStateZoneID("ShippingPostcode")}, // derive state from postcode
			{"ShippingPostcode", "shipping_postcode", JustUse("ShippingPostcode")},
			{"ShippingCountry", "shipping_country", JustUse("ShippingCountry")},
			{"", "shipping_country_id", MapCountryToCode("ShippingCountry")},
			{"ShippingMethod", "shipping_method", JustUse("ShippingMethod")},
			{"ShippingCode", "shipping_code", func(Entity) interface{} { return "Default shipping." }}, // TODO: change this to a mapping function
			{"AllComments", "comment", JustUse("AllComments")},
			{"Total", "total", JustUse("Total")},
			{"OrderStatus", "order_status_id", MapOrderStatusID},
			{"DateAdded", "date_added", JustUse("DateAdded")},
			{"DateModified", "date_modified", GetDateModified("DateModified")},
			{"", "date_payment_due", GetDateDue("DatePaymentDue")},
			{"", "currency_id", func(entity Entity) interface{} { return "1" }},             // Default currency ID for AUD
			{"", "currency_code", func(entity Entity) interface{} { return "AUD" }},         // Default currency code for AUD
			{"", "currency_value", func(entity Entity) interface{} { return "1.00000000" }}, // Default currency value for AUD
		},
	}
}

// INSERT INTO `oc_order_product` (`order_product_id`, `order_id`, `product_id`, `name`, `model`, `quantity`, `price`, `total`, `tax`, `reward`) VALUES ('4', '2', '0', 'Patch Lead for Bigpond Ultimate 312U USB', 'PL2001', '1', '19.9500', '17.9500', '2.0000', '0');
func GetOrderProductMapping(productIdMapping map[string]int) TableMapping {
	return TableMapping{
		TableName:   "oc_order_product",
		ColumnOrder: []string{"order_id", "product_id", "name", "model", "quantity", "price", "total", "tax", "reward"},
		Fields: []FieldMapping{
			// Assuming you'll handle order_product_id auto-increment outside this mapping.
			{"OrderID", "order_id", StripNPrefix},
			{"OrderLineSKU", "product_id", MapSKUToProductID(productIdMapping)}, // You'd want to change "ProductSKU" to the name from your CSV, for example "OrderLineSKU"
			{"OrderLineDescription", "name", JustUse("OrderLineDescription")},
			{"OrderLineSKU", "model", JustUse("OrderLineSKU")}, // Assuming SKU is also the model.
			{"OrderLineQty", "quantity", JustUse("OrderLineQty")},
			{"OrderLineUnitPrice", "price", CalculateOrderLineExGST},
			// For fields like "total", "tax", and "reward" you might need to calculate values or define new mapping functions.
			{"ProductTotal", "total", CalculateTotal},
			{"ProductTax", "tax", CalculateTax},
			{"ProductReward", "reward", CalculateReward},
		},
	}
}

func MapCustomerEmailToID(customerIdMapping map[string]int) func(entity Entity) interface{} {
	return func(entity Entity) interface{} {
		email, ok := entity.GetValue("Email").(string)
		if !ok {
			// fmt.Println("WARNING: Unable to convert Email to string for entity:", entity)
			return "0" // or some other default value or behavior
		}

		// fmt.Println("Processing email:", email) // Debugging line
		if id, exists := customerIdMapping[email]; exists {
			return strconv.Itoa(id)
		} else {
			// fmt.Printf("WARNING: Email %s not found in customer mapping. Assigning default ID.\n", email)
			return "0"
		}
	}
}

func StripNPrefix(value Entity) interface{} {
	orderId := value.GetValue("OrderID").(string)
	return strings.ReplaceAll(orderId, "N", "10")
}

func MapSKUToProductID(productIdMapping map[string]int) func(entity Entity) interface{} {
	return func(entity Entity) interface{} {
		sku, ok := entity.GetValue("OrderLineSKU").(string)
		if !ok {
			fmt.Println("WARNING: Unable to convert OrderLineSKU to string for entity:", entity)
			return "0" // or some other default value or behavior
		}

		// fmt.Println("Processing SKU:", sku) // Debugging line
		if id, exists := productIdMapping[sku]; exists {
			return strconv.Itoa(id)
		} else {
			fmt.Printf("WARNING: SKU %s not found in product mapping. Assigning default ID.\n", sku)
			return "0"
		}
	}
}

func CalculateOrderLineExGST(entity Entity) interface{} {
	priceStr := entity.GetValue("OrderLineUnitPrice").(string)

	// Convert the string values to decimals
	priceDec, err := decimal.NewFromString(priceStr)

	// Handle potential conversion errors
	if err != nil {
		fmt.Printf("Error converting price or quantity to decimal: %v\n", err)
		panic("Error converting price or quantity to decimal")
	}

	// Calculate total excluding GST
	taxDivisor := decimal.NewFromFloat(1.1)        // 10% GST
	exGSTPrice := priceDec.DivRound(taxDivisor, 2) // Divide price by 1.1 to get ex-GST price, rounded to 4 decimal places

	// Convert total to string with 4 decimal places
	return exGSTPrice.StringFixed(2)
}

// These are placeholder functions, you'd have to implement logic to calculate these.
func CalculateTotal(entity Entity) interface{} {
	priceStr := entity.GetValue("OrderLineUnitPrice").(string)
	qtyStr := entity.GetValue("OrderLineQty").(string)

	// Convert the string values to decimals
	priceDec, err1 := decimal.NewFromString(priceStr)
	qtyDec, err2 := decimal.NewFromString(qtyStr)

	// Handle potential conversion errors
	if err1 != nil || err2 != nil {
		fmt.Printf("Error converting price or quantity to decimal: %v, %v\n", err1, err2)
		panic("Error converting price or quantity to decimal")
	}

	// Calculate total excluding GST
	taxDivisor := decimal.NewFromFloat(1.1)        // 10% GST
	exGSTPrice := priceDec.DivRound(taxDivisor, 2) // Divide price by 1.1 to get ex-GST price, rounded to 4 decimal places
	total := exGSTPrice.Mul(qtyDec)

	// Convert total to string with 4 decimal places
	return total.StringFixed(2)
}

func CalculateTax(entity Entity) interface{} {
	priceStr := entity.GetValue("OrderLineUnitPrice").(string)

	// Convert the string values to decimals
	priceDec, err := decimal.NewFromString(priceStr)

	// Handle potential conversion errors
	if err != nil {
		fmt.Printf("Error converting price or quantity to decimal: %v\n", err)
		panic("Error converting price or quantity to decimal")
	}

	// Calculate total excluding GST
	taxDivisor := decimal.NewFromFloat(1.1)        // 10% GST
	exGSTPrice := priceDec.DivRound(taxDivisor, 2) // Divide price by 1.1 to get ex-GST price, rounded to 4 decimal places

	return priceDec.Sub(exGSTPrice).StringFixed(2)
}

func CalculateReward(entity Entity) interface{} {
	// Placeholder logic
	return "0.0000" // Assuming no reward
}

func GenerateOrderTotalSQLStatements(orderID, subTotalValue, shippingCost, taxValue, totalValue string) []string {
	statements := []string{
		fmt.Sprintf("INSERT IGNORE INTO `oc_order_total` (`order_id`, `code`, `title`, `value`, `sort_order`) VALUES ('%s', 'sub_total', 'Sub-Total', '%s', 1);", orderID, subTotalValue),
		fmt.Sprintf("INSERT IGNORE INTO `oc_order_total` (`order_id`, `code`, `title`, `value`, `sort_order`) VALUES ('%s', 'shipping', 'Shipping', '%s', 3);", orderID, shippingCost),
		fmt.Sprintf("INSERT IGNORE INTO `oc_order_total` (`order_id`, `code`, `title`, `value`, `sort_order`) VALUES ('%s', 'total', 'Total', '%s', 6);", orderID, totalValue),
		fmt.Sprintf("INSERT IGNORE INTO `oc_order_total` (`order_id`, `code`, `title`, `value`, `sort_order`) VALUES ('%s', 'tax', 'GST', '%s', 5);", orderID, taxValue),
		fmt.Sprintf("UPDATE `oc_order` set total = '%s' WHERE order_id = '%s';", totalValue, orderID),
	}

	return statements
}

func MapTotalTitle(entity Entity) interface{} {
	return "Total"
}

func safeToFloat(value interface{}) float64 {
	switch v := value.(type) {
	case float64:
		return v
	case string:
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return f
		}
		fmt.Printf("Warning: Unable to convert string %s to float. Using 0.0 as default.\n", v)
	default:
		fmt.Printf("Warning: Unexpected type %T (%v). Using 0.0 as default.\n", v, v)
	}
	return 0.0
}

func CalculateOrderTotals(lineItems []OrderRecord) (subTotalValue, shippingCost, taxValue, totalValue decimal.Decimal) {

	taxDivisor := decimal.NewFromFloat(1.1) // Divisor to remove GST

	for _, item := range lineItems {
		// Parse price and quantity to decimal
		price, err := decimal.NewFromString(item.OrderLineUnitPrice)
		if err != nil {
			price = decimal.Zero // Handle the error as appropriate
		}
		quantity, err := decimal.NewFromString(item.OrderLineQty)
		if err != nil {
			quantity = decimal.Zero // Handle the error as appropriate
		}

		// Remove GST from the price by dividing by taxDivisor (1 + taxRate)
		exGSTPrice := price.Div(taxDivisor)

		// Compute the sub-total for the current item
		subTotalItem := exGSTPrice.Mul(quantity)
		subTotalValue = subTotalValue.Add(subTotalItem)

		// Determine if the item is tax-free
		itemTaxValue := decimal.Zero
		if !strings.Contains(item.OrderLineTaxFree, "y") {
			// Compute the tax for the current item (tax was already included in the price)
			itemTaxValue = price.Sub(exGSTPrice).Mul(quantity)
			taxValue = taxValue.Add(itemTaxValue)
		}
	}

	// Assuming shipping cost is same for all line items, take from first item
	if len(lineItems) > 0 {
		shippingCostParsed, err := decimal.NewFromString(lineItems[0].ShippingCost)
		if err != nil {
			shippingCost = decimal.Zero // Handle the error as appropriate
		} else {
			shippingCost = shippingCostParsed
		}
	}

	// Compute the total value
	totalValue = subTotalValue.Add(shippingCost).Add(taxValue)
	return
}

func MapTitle(entity Entity) interface{} {
	code := entity.GetValue("AmountPaid").(string) // Assumes AmountPaid is a string
	if code == "total" {
		return "Total"
	} else if code == "sub_total" {
		return "Sub-Total"
	} else if code == "shipping" {
		return "Shipping"
	} // add more cases as necessary
	return "Unknown"
}

func MapOrderID(orderIDMapping map[string]int) MappingFunction {
	return func(entity Entity) interface{} {
		orderID, _ := entity.GetValue("OrderID").(string)
		return orderIDMapping[orderID]
	}
}

func MapTotalCode(entity Entity) interface{} {
	title, _ := entity.GetValue("TotalTitle").(string)
	switch title {
	case "Sub-Total":
		return "sub_total"
	case "Shipping":
		return "shipping"
	case "Total":
		return "total"
	case "VAT":
		return "tax"
	default:
		return "" // return an empty string or a default value for unknown titles
	}
}

func MapSortOrder(entity Entity) interface{} {
	title, _ := entity.GetValue("TotalTitle").(string)
	switch title {
	case "Sub-Total":
		return 1
	case "Shipping":
		return 3
	case "Total":
		return 6
	case "VAT":
		return 5
	default:
		return 99 // default high sort order for unknown titles
	}
}

type MappingFunction func(entity Entity) interface{}

func MapCustomerGroupID(entity Entity) interface{} {
	groupName, ok := entity.GetValue("CustomerGroupID").(string)
	if !ok {
		// Handle the case where the conversion failed or the field does not exist.
		return nil
	}

	switch groupName {
	case "Trade":
		return "2"
	case "Retail":
		return "3"
	default:
		return "1"
	}
}

func MapOrderStatusID(entity Entity) interface{} {
	if entity.GetValue("AmountPaid") != entity.GetValue("Total") {
		return "15"
	}

	status, _ := entity.GetValue("OrderStatus").(string)
	statusMap := map[string]string{
		"Dispatched":      "4",
		"Cancelled":       "5",
		"Pending Pickup":  "13",
		"Quote":           "10",
		"Pick":            "2",
		"New":             "1",
		"Pack":            "3",
		"On Hold":         "14",
		"Sent On Account": "15",
	}

	return statusMap[status]
}
