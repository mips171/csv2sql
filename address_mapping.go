package main

func GetAddressMapping() TableMapping {
	return TableMapping{
		TableName:   "oc_address",
		ColumnOrder: []string{"customer_id", "firstname", "lastname", "company", "address_1", "address_2", "city", "postcode"},
		Fields: []FieldMapping{
			{"Bill First Name", "firstname", DoNothing("Bill First Name")},
			{"Bill Last Name", "lastname", DoNothing("Bill Last Name")},
			{"Bill Company", "company", DoNothing("Bill Company")},
			{"Bill Street Address Line 1", "address_1", DoNothing("Bill Street Address Line 1")},
			{"Bill Street Address Line 2", "address_2", DoNothing("Bill Street Address Line 2")},
			{"Bill City", "city", DoNothing("Bill City")},
			{"Bill Post Code", "postcode", DoNothing("Bill Post Code")},
		},
	}
}
