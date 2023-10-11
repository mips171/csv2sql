package main

func GetAddressMapping() TableMapping {
	return TableMapping{
		TableName:   "oc_address",
		ColumnOrder: []string{"customer_id", "firstname", "lastname", "company", "address_1", "address_2", "city", "postcode"},
		Fields: []FieldMapping{
			{"Bill First Name", "firstname", JustUse("Bill First Name")},
			{"Bill Last Name", "lastname", JustUse("Bill Last Name")},
			{"Bill Company", "company", JustUse("Bill Company")},
			{"Bill Street Address Line 1", "address_1", JustUse("Bill Street Address Line 1")},
			{"Bill Street Address Line 2", "address_2", JustUse("Bill Street Address Line 2")},
			{"Bill City", "city", JustUse("Bill City")},
			{"Bill Post Code", "postcode", JustUse("Bill Post Code")},
		},
	}
}
