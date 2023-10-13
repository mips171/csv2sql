package main

// "Username*","Email Address","Newsletter Subscriber","Bill Phone","Bill Fax","Bill First Name","Bill Last Name","Bill Company","Bill Street Address Line 1","Bill Street Address Line 2","Bill City","Bill State","Bill Post Code","Bill Country","Cheque Account Number","Ship Post Code","Ship Address Line2","Ship Address Line 1","Parent Username","Ship State","Quote Approval Username","Active","Ship Company","Default Invoice Terms","Website URL","Ship Phone","Permission","ABN/ACN","Cheque Account Name","Ship Country","Default Discount Percentage","Ship Fax","Ship City","Gender","Referral Commission","Account Manager","Secondary Email Address","Ship Last Name","Ship First Name","Default Document Template","Referral Username","Default Order Type","Date Of Birth","VIP Customer","Sales Channel","Type","Survey's and Special Orders","Classification 2","Credit Limit","Classification 1","On Credit Hold","Skip Record","Internal Notes","Default Shipping Address"
type CustomerRecord struct {
	Email      string `csv:"Email Address"`
	Group      string `csv:"User Group"`
	FirstName  string `csv:"First Name"`
	Surname   string `csv:"Last Name"`
	Address    string `csv:"Address"`
	City       string `csv:"City"`
	Phone      string `csv:"Phone"`
	Newsletter string `csv:"Newsletter Subscriber"`
	Status     string `csv:"Active"`
	ABNACN     string `csv:"ABN/ACN"`
	BillFirstName string `csv:"Bill First Name"`
	BillLastName string `csv:"Bill Last Name"`
	BillCompany string `csv:"Bill Company"`
	BillAddress string `csv:"Bill Street Address Line 1"`
	BillCity string `csv:"Bill City"`
	BillState string `csv:"Bill State"`
	BillPostCode string `csv:"Bill Post Code"`
	BillCountry string `csv:"Bill Country"`
	ShipFirstName string `csv:"Ship First Name"`
	ShipLastName string `csv:"Ship Last Name"`
	ShipCompany string `csv:"Ship Company"`
	ShipAddress string `csv:"Ship Address Line 1"`
	ShipCity string `csv:"Ship City"`
	ShipState string `csv:"Ship State"`
	ShipPostCode string `csv:"Ship Post Code"`
	ShipCountry string `csv:"Ship Country"`
	ShipPhone string `csv:"Ship Phone"`
	ShipAddress2 string `csv:"Ship Address Line2"`
	ShipEmail string `csv:"Secondary Email Address"`
	ShipNotes string `csv:"Internal Notes"`
	ShipDefault string `csv:"Default Shipping Address"`
	//... any other fields your CSV might have.
}

// INSERT INTO `oc_customer` (`customer_id`, `customer_group_id`, `store_id`, `language_id`, `firstname`, `lastname`, `email`, `telephone`, `fax`, `password`, `salt`, `cart`, `wishlist`, `newsletter`, `address_id`, `custom_field`, `ip`, `status`, `safe`, `token`, `code`, `date_added`) VALUES ('1', '3', '0', '0', '', '', 'bradv@abmresources.com.au', '', '', '', '', 'a:0:{}', '', '1', '75441', '', '0', '1', '0', '', '', '2017-11-02 03:00:06');

func GetCustomerMapping() TableMapping {
	return TableMapping{
		TableName:   "oc_customer",
		ColumnOrder: []string{"customer_id", "customer_group_id", "firstname", "lastname", "email", "telephone", "newsletter", "status", "date_added"},
		Fields: []FieldMapping{
			{"Group", "customer_group_id", GetUserGroupID("Group")}, // note - need to pre-clean the data and add this field.
			{"FirstName", "firstname", GetFirstName("FirstName", "Email")}, // use firstname as the default, email as backup
			{"LastName", "lastname", GetLastName("Surname", "Email")}, // use lastname as the default, email as backup
			{"EmailAddress", "email", JustUse("Email")},
			{"Phone", "telephone", JustUse("Phone")},
			{"Status", "status", GetStatus("Status")},
			{"Newsletter", "newsletter", GetNewsletterStatus("Newsletter")},
			{"", "date_added", GetDateAdded()},
			{"", "store_id", func(entity Entity) interface{} { return "0" }},
			{"", "language_id", func(entity Entity) interface{} { return "0" }},
		},
	}
}

