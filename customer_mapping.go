package main

import "strconv"

type CustomerGroupRecord struct {
	Group    string `csv:"User Group"`
	Username string `csv:"Username*"`
}

// "Username*","Email Address","Newsletter Subscriber","Bill Phone","Bill Fax","Bill First Name","Bill Last Name","Bill Company","Bill Street Address Line 1","Bill Street Address Line 2","Bill City","Bill State","Bill Post Code","Bill Country","Cheque Account Number","Ship Post Code","Ship Address Line2","Ship Address Line 1","Parent Username","Ship State","Quote Approval Username","Active","Ship Company","Default Invoice Terms","Website URL","Ship Phone","Permission","ABN/ACN","Cheque Account Name","Ship Country","Default Discount Percentage","Ship Fax","Ship City","Gender","Referral Commission","Account Manager","Secondary Email Address","Ship Last Name","Ship First Name","Default Document Template","Referral Username","Default Order Type","Date Of Birth","VIP Customer","Sales Channel","Type","Survey's and Special Orders","Classification 2","Credit Limit","Classification 1","On Credit Hold","Skip Record","Internal Notes","Default Shipping Address"
type CustomerRecord struct {
	Username         string `csv:"Username*"`
	Email            string `csv:"Email Address"`
	Group            string `csv:"User Group"`
	Phone            string `csv:"Bill Phone"`
	Newsletter       string `csv:"Newsletter Subscriber"`
	Status           string `csv:"Active"`
	ABNACN           string `csv:"ABN/ACN"`
	BillFirstName    string `csv:"Bill First Name"`
	BillLastName     string `csv:"Bill Last Name"`
	BillCompany      string `csv:"Bill Company"`
	BillAddress1     string `csv:"Bill Street Address Line 1"`
	BillAddress2     string `csv:"Bill Street Address Line 2"`
	BillCity         string `csv:"Bill City"`
	BillState        string `csv:"Bill State"`
	BillPostCode     string `csv:"Bill Post Code"`
	BillCountry      string `csv:"Bill Country"`
	ShipFirstName    string `csv:"Ship First Name"`
	ShipLastName     string `csv:"Ship Last Name"`
	ShipCompany      string `csv:"Ship Company"`
	ShipAddressLine1 string `csv:"Ship Address Line 1"`
	ShipAddressLine2 string `csv:"Ship Address Line 2"`
	ShipCity         string `csv:"Ship City"`
	ShipState        string `csv:"Ship State"`
	ShipPostCode     string `csv:"Ship Post Code"`
	ShipCountry      string `csv:"Ship Country"`
	ShipPhone        string `csv:"Ship Phone"`
	ShipNotes        string `csv:"Internal Notes"`
	//... any other fields your CSV might have.
}

func (p CustomerRecord) GetValue(fieldName string) interface{} {
	switch fieldName {
	case "Email":
		return p.Email
	case "Group":
		return p.Group
	case "Phone":
		return p.Phone
	case "Newsletter":
		return p.Newsletter
	case "Status":
		return p.Status
	case "ABNACN":
		return p.ABNACN
	case "BillFirstName":
		return p.BillFirstName
	case "BillLastName":
		return p.BillLastName
	case "BillCompany":
		return p.BillCompany
	case "BillAddress1":
		return p.BillAddress1
	case "BillAddress2":
		return p.BillAddress2
	case "BillCity":
		return p.BillCity
	case "BillState":
		return p.BillState
	case "BillPostCode":
		return p.BillPostCode
	case "BillCountry":
		return p.BillCountry
	case "ShipFirstName":
		return p.ShipFirstName
	case "ShipLastName":
		return p.ShipLastName
	case "ShipCompany":
		return p.ShipCompany
	case "ShipAddressLine1":
		return p.ShipAddressLine1
	case "ShipAddressLine2":
		return p.ShipAddressLine2
	case "ShipCity":
		return p.ShipCity
	case "ShipState":
		return p.ShipState
	case "ShipPostCode":
		return p.ShipPostCode
	case "ShipCountry":
		return p.ShipCountry
	case "ShipPhone":
		return p.ShipPhone
	case "InternalNotes":
		return p.ShipNotes
	default:
		return ""
	}
}

// INSERT INTO `oc_customer` (`customer_id`, `customer_group_id`, `store_id`, `language_id`, `firstname`, `lastname`, `email`, `telephone`, `fax`, `password`, `salt`, `cart`, `wishlist`, `newsletter`, `address_id`, `custom_field`, `ip`, `status`, `safe`, `token`, `code`, `date_added`) VALUES ('1', '3', '0', '0', '', '', 'bradv@abmresources.com.au', '', '', '', '', 'a:0:{}', '', '1', '75441', '', '0', '1', '0', '', '', '2017-11-02 03:00:06');

func GetCustomerMapping() TableMapping {
	return TableMapping{
		TableName:   "oc_customer",
		ColumnOrder: []string{"customer_group_id", "firstname", "lastname", "email", "telephone", "newsletter", "status", "date_added", "store_id", "language_id"},
		Fields: []FieldMapping{
			{"Group", "customer_group_id", GetUserGroupID("Group")},  // note - need to pre-clean the data and add this field.
			{"BillFirstName", "firstname", JustUse("BillFirstName")}, // use firstname as the default, email as backup
			{"BillLastName", "lastname", JustUse("BillLastName")},    // use lastname as the default, email as backup
			{"Email", "email", JustUse("Email")},
			{"Phone", "telephone", JustUse("Phone")},
			{"Newsletter", "newsletter", GetNewsletterStatus("Newsletter")},
			{"Status", "status", GetStatus("Status")},
			{"", "date_added", GetDateAdded()},
			{"", "store_id", func(entity Entity) interface{} { return "0" }},
			{"", "language_id", func(entity Entity) interface{} { return "0" }},
		},
	}
}

func GetCustomerAddressMapping(customerIdMapping map[string]int) TableMapping {
	return TableMapping{
		TableName:   "oc_address",
		ColumnOrder: []string{"customer_id", "firstname", "lastname", "company", "address_1", "address_2", "city", "postcode", "country_id", "zone_id"},
		Fields: []FieldMapping{
			{"", "customer_id", GetCustomerIdTransformation(customerIdMapping)},
			{"BillFirstName", "firstname", JustUse("BillFirstName")}, // use firstname as the default, email as backup
			{"BillLastName", "lastname", JustUse("BillLastName")},    // use lastname as the default, email as backup
			{"BillCompany", "company", JustUse("BillCompany")},
			{"BillAddress", "address_1", JustUse("BillAddress")},
			{"BillAddress2", "address_2", JustUse("BillAddress2")},
			{"BillCity", "city", JustUse("BillCity")},
			{"BillPostCode", "postcode", JustUse("BillPostCode")},
			{"BillCountry", "country_id", MapCountryToCode("BillCountry")},
			{"", "zone_id", MapAustralianPostCodeToStateZoneID("BillPostCode")}, // derive state from postcode
		},
	}
}

func GetCustomerIdTransformation(productIdMapping map[string]int) func(entity Entity) interface{} {
	return func(entity Entity) interface{} {
		model := entity.GetValue("Email").(string)
		if id, exists := productIdMapping[model]; exists {
			return strconv.Itoa(id)
		}
		return nil
	}
}
