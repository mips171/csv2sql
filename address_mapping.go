package main

func GetAddressMapping() TableMapping {
	return TableMapping{
		TableName:   "oc_address",
		ColumnOrder: []string{"customer_id", "firstname", "lastname", "company", "address_1", "address_2", "city", "postcode"},
		Fields: []FieldMapping{
			{"BillFirstName", "firstname", JustUse("BillFirstName")},
			{"BillLastName", "lastname", JustUse("BillLastName")},
			{"BillCompany", "company", JustUse("Bill Company")},
			{"BillStreetAddressLine1", "address_1", JustUse("Bill Street Address Line 1")},
			{"BillStreetAddressLine2", "address_2", JustUse("Bill Street Address Line 2")},
			{"BillCity", "city", JustUse("Bill City")},
			{"BillPostCode", "postcode", JustUse("Bill Post Code")},
		},
	}
}
