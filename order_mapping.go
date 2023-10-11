package main

import (
	"fmt"
	"strconv"
)

// CSV Headers, need to match OrderRecord deserialization
// "Order ID","Order Status","Approved","Username","Email","Ship First Name","Ship Last Name","Ship Company","Ship Address Line 1","Ship Address Line 2","Ship City","Ship State","Ship Post Code","Ship Country","Ship Phone","Ship Fax","RUT950 Serial Number","RUT950 IMEI","Transaction Number","Misc 0","Misc Notes 1","Installation Service Notes","Misc Notes 2","Bill First Name","Bill Last Name","Bill Company","Bill Address Line 1","Bill Address Line 2","Bill City","Bill State","Bill Post Code","Bill Country","Bill Phone","Bill Fax","Payment Method","Shipping Method","Customer Instructions","Internal Notes","Amount Paid","Date Paid","Order Line SKU","Order Line Qty","Order Line Description","Order Line Unit Price","Order Line Unit Cost","Tax Free Shipping","Card Holder","Shipping Discount Amount","Order Line Serial Number","Order Payment Plan","Parent Order ID","Payment Due Date","User Group","Fraud Score","BPAY CRN","Order Line Bin Location","Sales Channel","Order Line Options","Order Line Dropship Supplier","Order Line Tax Free","Order Line Discount Amount","Order Line Shipping Cubic","Order Line Job","Tax Inclusive","Purchase Order ID","Order Type","Date Required","Order Line Shipping Weight","Order Line Notes","Sales Person","Currency Code","Date Invoiced","Shipping Cost","Order Line Delivery Date","Payment Terms","Date Placed","Document Template","Coupon Code"
// What we will actually import
// INSERT INTO `oc_order` (`order_id`, `invoice_no`, `invoice_prefix`, `store_id`, `store_name`, `store_url`, `customer_id`, `customer_group_id`, `firstname`, `lastname`, `email`, `telephone`, `fax`, `custom_field`, `payment_firstname`, `payment_lastname`, `payment_company`, `payment_address_1`, `payment_address_2`, `payment_city`, `payment_postcode`, `payment_country`, `payment_country_id`, `payment_zone`, `payment_zone_id`, `payment_address_format`, `payment_custom_field`, `payment_method`, `payment_code`, `shipping_firstname`, `shipping_lastname`, `shipping_company`, `shipping_address_1`, `shipping_address_2`, `shipping_city`, `shipping_postcode`, `shipping_country`, `shipping_country_id`, `shipping_zone`, `shipping_zone_id`, `shipping_address_format`, `shipping_custom_field`, `shipping_method`, `shipping_code`, `comment`, `total`, `order_status_id`, `affiliate_id`, `commission`, `marketing_id`, `tracking`, `language_id`, `currency_id`, `currency_code`, `currency_value`, `ip`, `forwarded_ip`, `user_agent`, `accept_language`, `date_added`, `date_modified`) VALUES ('2', '0', ”, '0', 'Telco Antennas', 'https://telcoshop.nbembedded.com/', '9', '0', 'Daryl', 'Sowinski', 'dsowinski@bigpond.com', '419653781', ”, ”, 'Daryl', 'Sowinski', ”, '37 Diamantina Street', ”, 'Hillcrest', '4118', 'Australia', '13', ”, '0', ”, ”, 'Direct Deposit (EFT)', 'cod', 'Daryl', 'Sowinski', ”, '37 Diamantina Street', ”, 'Hillcrest', '4118', 'Australia', '13', ”, '0', ”, ”, 'Default shipping.', ”, ”, '46.8900', '7', '0', '0.0000', '0', ”, '1', '0', 'AUD', '1.00000000', ”, ”, ”, ”, '2011-02-22 03:25:29', '2018-07-02 02:07:06');
// We want these two to match up, and ignore everything else.
type OrderRecord struct {
	OrderID              string `csv:"Order ID"`
	OrderStatus          string `csv:"Order Status"`
	Approved             string `csv:"Approved"`
	ShipFirstname        string `csv:"Ship First Name"`
	ShipLastname         string `csv:"Ship Last Name"`
	Email                string `csv:"Email"`
	Telephone            string `csv:"Ship Phone"`
	Fax                  string `csv:"Ship Fax"`
	PaymentFirstname     string `csv:"Bill First Name"`
	PaymentLastname      string `csv:"Bill Last Name"`
	PaymentAddress1      string `csv:"Bill Address Line 1"`
	PaymentAddress2      string `csv:"Bill Address Line 2"`
	PaymentCity          string `csv:"Bill City"`
	PaymentPostcode      string `csv:"Bill Post Code"`
	PaymentCountry       string `csv:"Bill Country"`
	ShippingFirstname    string `csv:"Ship First Name"`
	ShippingLastname     string `csv:"Ship Last Name"`
	ShippingCompany      string `csv:"Ship Company"`
	ShippingAddress1     string `csv:"Ship Address Line 1"`
	ShippingAddress2     string `csv:"Ship Address Line 2"`
	ShippingCity         string `csv:"Ship City"`
	ShippingPostcode     string `csv:"Ship Post Code"`
	ShippingCountry      string `csv:"Ship Country"`
	PaymentMethod        string `csv:"Payment Method"`
	ShippingMethod       string `csv:"Shipping Method"`
	ShippingCost         string `csv:"Shipping Cost"`
	Comment              string `csv:"Customer Instructions"`
	Total                string `csv:"Amount Paid"`
	PaymentCode          string // this will need to be mapped based on PaymentMethod, not present in CSV
	ShippingCode         string // this will need to be mapped based on ShippingMethod, not present in CSV
	CurrencyCode         string `csv:"Currency Code"`
	DateAdded            string `csv:"Date Placed"`
	DateModified         string `csv:"Date Invoiced"`
	OrderLineTaxFree     string `csv:"Order Line Tax Free"`
	OrderLineSKU         string `csv:"Order Line SKU"`
	OrderLineQty         string `csv:"Order Line Qty"`
	OrderLineDescription string `csv:"Order Line Description"`
	OrderLineUnitPrice   string `csv:"Order Line Unit Price"`
	AmountPaid           string `csv:"Amount Paid"`
	// Add any other relevant fields as needed.
}

func (o OrderRecord) GetValue(fieldName string) interface{} {
	switch fieldName {
	case "OrderID":
		return o.OrderID
	case "OrderStatus":
		return o.OrderStatus
	case "Approved":
		return o.Approved
	case "ShipFirstname":
		return o.ShipFirstname
	case "ShipLastname":
		return o.ShipLastname
	case "Email":
		return o.Email
	case "Telephone":
		return o.Telephone
	case "Fax":
		return o.Fax
	case "PaymentFirstname":
		return o.PaymentFirstname
	case "PaymentLastname":
		return o.PaymentLastname
	case "PaymentAddress1":
		return o.PaymentAddress1
	case "PaymentAddress2":
		return o.PaymentAddress2
	case "PaymentCity":
		return o.PaymentCity
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
	case "Comment":
		return o.Comment
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
	case "OrderLineSKU":
		return o.OrderLineSKU
	case "OrderLineQty":
		return o.OrderLineQty
	case "OrderLineDescription":
		return o.OrderLineDescription
	case "OrderLineUnitPrice":
		return o.OrderLineUnitPrice
	// ... add other fields as required from your CSV
	default:
		return nil
	}
}

// Need to map to this schema
// INSERT INTO `oc_order` (`order_id`, `invoice_no`, `invoice_prefix`, `store_id`, `store_name`, `store_url`, `customer_id`, `customer_group_id`, `firstname`, `lastname`, `email`, `telephone`, `fax`, `custom_field`, `payment_firstname`, `payment_lastname`, `payment_company`, `payment_address_1`, `payment_address_2`, `payment_city`, `payment_postcode`, `payment_country`, `payment_country_id`, `payment_zone`, `payment_zone_id`, `payment_address_format`, `payment_custom_field`, `payment_method`, `payment_code`, `shipping_firstname`, `shipping_lastname`, `shipping_company`, `shipping_address_1`, `shipping_address_2`, `shipping_city`, `shipping_postcode`, `shipping_country`, `shipping_country_id`, `shipping_zone`, `shipping_zone_id`, `shipping_address_format`, `shipping_custom_field`, `shipping_method`, `shipping_code`, `comment`, `total`, `order_status_id`, `affiliate_id`, `commission`, `marketing_id`, `tracking`, `language_id`, `currency_id`, `currency_code`, `currency_value`, `ip`, `forwarded_ip`, `user_agent`, `accept_language`, `date_added`, `date_modified`) VALUES ('2', '0', ”, '0', 'Telco Antennas', 'https://telcoshop.nbembedded.com/', '9', '0', 'Daryl', 'Sowinski', 'dsowinski@bigpond.com', '419653781', ”, ”, 'Daryl', 'Sowinski', ”, '37 Diamantina Street', ”, 'Hillcrest', '4118', 'Australia', '13', ”, '0', ”, ”, 'Direct Deposit (EFT)', 'cod', 'Daryl', 'Sowinski', ”, '37 Diamantina Street', ”, 'Hillcrest', '4118', 'Australia', '13', ”, '0', ”, ”, 'Default shipping.', ”, ”, '46.8900', '7', '0', '0.0000', '0', ”, '1', '0', 'AUD', '1.00000000', ”, ”, ”, ”, '2011-02-22 03:25:29', '2018-07-02 02:07:06');

func GetOrderMapping() TableMapping {
	return TableMapping{
		TableName:   "oc_order",
		ColumnOrder: []string{"invoice_no", "invoice_prefix", "store_id", "store_name", "store_url", "customer_id", "customer_group_id", "firstname", "lastname", "email", "telephone", "fax", "custom_field", "payment_firstname", "payment_lastname", "payment_company", "payment_address_1", "payment_address_2", "payment_city", "payment_postcode", "payment_country", "payment_country_id", "payment_zone", "payment_zone_id", "payment_address_format", "payment_custom_field", "payment_method", "payment_code", "shipping_firstname", "shipping_lastname", "shipping_company", "shipping_address_1", "shipping_address_2", "shipping_city", "shipping_postcode", "shipping_country", "shipping_country_id", "shipping_zone", "shipping_zone_id", "shipping_address_format", "shipping_custom_field", "shipping_method", "shipping_code", "comment", "total", "order_status_id", "affiliate_id", "commission", "marketing_id", "tracking", "language_id", "currency_id", "currency_code", "currency_value", "ip", "forwarded_ip", "user_agent", "accept_language", "date_added", "date_modified"},
		Fields: []FieldMapping{
			// No need for order_id since it's managed by the database.
			// As an example, here are a few more mappings:
			{"InvoiceNo", "invoice_no", DoNothing("InvoiceNo")},
			{"InvoicePrefix", "invoice_prefix", DoNothing("InvoicePrefix")},
			{"StoreID", "store_id", func(entity Entity) interface{} { return "1" }},
			{"StoreName", "store_name", func(entity Entity) interface{} { return "Telco Antennas" }},
			{"StoreURL", "store_url", func(entity Entity) interface{} { return "https://telcoantennas.com.au/" }},
			{"CustomerID", "customer_id", DoNothing("CustomerID")},
			{"CustomerGroupID", "customer_group_id", MapCustomerGroupID},
			{"PaymentFirstname", "firstname", DoNothing("PaymentFirstname")},
			{"PaymentLastname", "lastname", DoNothing("PaymentLastname")},
			{"Email", "email", DoNothing("Email")},
			{"Telephone", "telephone", DoNothing("Telephone")},
			{"Fax", "fax", DoNothing("Fax")},
			{"CustomField", "custom_field", DoNothing("CustomField")},
			{"PaymentFirstname", "payment_firstname", DoNothing("PaymentFirstname")},
			{"PaymentLastname", "payment_lastname", DoNothing("PaymentLastname")},
			{"PaymentCompany", "payment_company", DoNothing("PaymentCompany")},
			{"PaymentAddress1", "payment_address_1", DoNothing("PaymentAddress1")},
			{"PaymentAddress2", "payment_address_2", DoNothing("PaymentAddress2")},
			{"PaymentCity", "payment_city", DoNothing("PaymentCity")},
			{"PaymentPostcode", "payment_postcode", DoNothing("PaymentPostcode")},
			{"PaymentCountry", "payment_country", DoNothing("PaymentCountry")},
			{"PaymentCountryID", "payment_country_id", DoNothing("PaymentCountryID")},
			{"PaymentZone", "payment_zone", DoNothing("PaymentZone")},
			{"PaymentZoneID", "payment_zone_id", DoNothing("PaymentZoneID")},
			{"PaymentAddressFormat", "payment_address_format", DoNothing("PaymentAddressFormat")},
			{"PaymentCustomField", "payment_custom_field", DoNothing("PaymentCustomField")},
			{"PaymentMethod", "payment_method", DoNothing("PaymentMethod")},
			{"PaymentCode", "payment_code", func(entity Entity) interface{} { return "cod" }}, // TODO: change this to a mapping function
			{"ShippingFirstname", "shipping_firstname", DoNothing("ShippingFirstname")},
			{"ShippingLastname", "shipping_lastname", DoNothing("ShippingLastname")},
			{"ShippingCompany", "shipping_company", DoNothing("ShippingCompany")},
			{"ShippingAddress1", "shipping_address_1", DoNothing("ShippingAddress1")},
			{"ShippingAddress2", "shipping_address_2", DoNothing("ShippingAddress2")},
			{"ShippingCity", "shipping_city", DoNothing("ShippingCity")},
			{"ShippingPostcode", "shipping_postcode", DoNothing("ShippingPostcode")},
			{"ShippingCountry", "shipping_country", DoNothing("ShippingCountry")},
			{"ShippingCountryID", "shipping_country_id", DoNothing("ShippingCountryID")},
			{"ShippingZone", "shipping_zone", DoNothing("ShippingZone")},
			{"ShippingZoneID", "shipping_zone_id", DoNothing("ShippingZoneID")},
			{"ShippingAddressFormat", "shipping_address_format", DoNothing("ShippingAddressFormat")},
			{"ShippingCustomField", "shipping_custom_field", DoNothing("ShippingCustomField")},
			{"ShippingMethod", "shipping_method", DoNothing("ShippingMethod")},
			{"ShippingCode", "shipping_code", func(Entity) interface{} { return "Default shipping." }}, // TODO: change this to a mapping function
			{"Comment", "comment", DoNothing("Comment")},
			{"Total", "total", DoNothing("Total")},
			{"OrderStatusID", "order_status_id", MapOrderStatusID},
			{"AffiliateID", "affiliate_id", DoNothing("AffiliateID")},
			{"Commission", "commission", DoNothing("Commission")},
			{"MarketingID", "marketing_id", DoNothing("MarketingID")},
			{"Tracking", "tracking", DoNothing("Tracking")},
			{"LanguageID", "language_id", DoNothing("LanguageID")},
			{"CurrencyID", "currency_id", DoNothing("CurrencyID")},
			{"CurrencyCode", "currency_code", DoNothing("CurrencyCode")},
			{"CurrencyValue", "currency_value", DoNothing("CurrencyValue")},
			{"IP", "ip", DoNothing("IP")},
			{"ForwardedIP", "forwarded_ip", DoNothing("ForwardedIP")},
			{"UserAgent", "user_agent", DoNothing("UserAgent")},
			{"AcceptLanguage", "accept_language", DoNothing("AcceptLanguage")},
			{"DateAdded", "date_added", DoNothing("DateAdded")},
			{"DateModified", "date_modified", DoNothing("DateModified")},
		},
	}
}

// INSERT INTO `oc_order_product` (`order_product_id`, `order_id`, `product_id`, `name`, `model`, `quantity`, `price`, `total`, `tax`, `reward`) VALUES ('4', '2', '0', 'Patch Lead for Bigpond Ultimate 312U USB', 'PL2001', '1', '19.9500', '17.9500', '2.0000', '0');
func GetOrderProductMapping(orderIDMapping map[string]int, productIdMapping map[string]int) TableMapping {
	return TableMapping{
		TableName:   "oc_order_product",
		ColumnOrder: []string{"order_id", "product_id", "name", "model", "quantity", "price", "total", "tax", "reward"},
		Fields: []FieldMapping{
			// Assuming you'll handle order_product_id auto-increment outside this mapping.
			{"OrderID", "order_id", MapOrderID(orderIDMapping)},
			{"OrderLineSKU", "product_id", MapSKUToProductID(productIdMapping)}, // You'd want to change "ProductSKU" to the name from your CSV, for example "OrderLineSKU"
			{"OrderLineDescription", "name", DoNothing("OrderLineDescription")},
			{"OrderLineSKU", "model", DoNothing("OrderLineSKU")}, // Assuming SKU is also the model.
			{"OrderLineQty", "quantity", DoNothing("OrderLineQty")},
			{"OrderLineUnitPrice", "price", DoNothing("OrderLineUnitPrice")},
			// For fields like "total", "tax", and "reward" you might need to calculate values or define new mapping functions.
			{"ProductTotal", "total", CalculateTotal},
			{"ProductTax", "tax", CalculateTax},
			{"ProductReward", "reward", CalculateReward},
		},
	}
}

func MapSKUToProductID(productIdMapping map[string]int) func(entity Entity) interface{} {
	return func(entity Entity) interface{} {
		sku, ok := entity.GetValue("OrderLineSKU").(string)
		if !ok {
			fmt.Println("WARNING: Unable to convert OrderLineSKU to string for entity:", entity)
			return "0" // or some other default value or behavior
		}

		fmt.Println("Processing SKU:", sku) // Debugging line
		if id, exists := productIdMapping[sku]; exists {
			return strconv.Itoa(id)
		} else {
			fmt.Printf("WARNING: SKU %s not found in product mapping. Assigning default ID.\n", sku)
			return "0"
		}
	}
}

// These are placeholder functions, you'd have to implement logic to calculate these.
func CalculateTotal(entity Entity) interface{} {
	priceStr := entity.GetValue("OrderLineUnitPrice").(string)
	qtyStr := entity.GetValue("OrderLineQty").(string)

	// Convert the string values to float64 and int respectively
	price, err1 := strconv.ParseFloat(priceStr, 64)
	qty, err2 := strconv.Atoi(qtyStr)

	// Handle potential conversion errors
	if err1 != nil || err2 != nil {
		fmt.Printf("Error converting price or quantity to number: %v, %v\n", err1, err2)
		panic("Error converting price or quantity to number")
	}

	total := price * float64(qty)

	// Convert total to string with 4 decimal places
	return fmt.Sprintf("%.4f", total)
}

func CalculateTax(entity Entity) interface{} {
	// Retrieve the price as a string
	priceStr := entity.GetValue("OrderLineUnitPrice").(string)

	// Convert the string to float64
	price, err := strconv.ParseFloat(priceStr, 64)

	// Handle potential conversion errors
	if err != nil {
		fmt.Printf("Error converting price to number: %v\n", err)
		return "ERROR" // or handle this more gracefully, depending on your needs
	}

	// Calculate the tax (10% in this example)
	tax := price * 0.1

	// Convert tax to string with 4 decimal places
	return fmt.Sprintf("%.4f", tax)
}

func CalculateReward(entity Entity) interface{} {
	// Placeholder logic
	return "0.0000" // Assuming no reward
}

func GetOrderTotalMapping(orderIDMapping map[string]int) TableMapping {
	return TableMapping{
		TableName:   "oc_order_total",
		ColumnOrder: []string{"order_id", "code", "title", "value", "sort_order"},
		Fields: []FieldMapping{
			{"OrderID", "order_id", MapOrderID(orderIDMapping)},
			{"AmountPaid", "code", MapTotalCode},
			{"AmountPaid", "title", MapTotalTitle},
			{"AmountPaid", "value", CalculateTotalValue},
			{"AmountPaid", "sort_order", MapSortOrder},
			//... Additional fields as required from your CSV
		},
	}
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

func CalculateTotalValue(entity Entity) interface{} {
	taxRate := 0.1
	if entity.GetValue("OrderLineTaxFree") == "y" {
		taxRate = 0.0
	}

	// Calculate sub-total for the line item
	price := safeToFloat(entity.GetValue("OrderLineUnitPrice"))
	quantity := safeToFloat(entity.GetValue("OrderLineQty"))
	subTotalValue := price * quantity

	// Add ShippingCost to the sub-total
	shippingCost := safeToFloat(entity.GetValue("ShippingCost"))
	subTotalValue += shippingCost

	// Calculate tax on the sub-total
	taxValue := subTotalValue * taxRate

	totalValue := subTotalValue + taxValue

	// Debugging output
	fmt.Printf("OrderID: %v, ShippingCost: %f, SubTotal: %f, TaxValue: %f, TotalValue: %f\n",
		entity.GetValue("OrderID"), shippingCost, subTotalValue, taxValue, totalValue)

	return fmt.Sprintf("%.4f", totalValue)
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
	status, _ := entity.GetValue("OrderStatusID").(string)
	statusMap := map[string]string{
		"Dispatched":     "3",
		"Cancelled":      "7",
		"Pending Pickup": "1",
		"Quote":          "17",
		"Pick":           "2",
		"New":            "1",
		"Pack":           "2",
		"On Hold":        "8",
	}

	return statusMap[status]
}
