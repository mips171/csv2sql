package main

func GetAddressMapping() TableMapping {
	return TableMapping{
		TableName:   "oc_address",
		ColumnOrder: []string{"customer_id", "firstname", "lastname", "company", "address_1", "address_2", "city", "postcode"},
		Fields: []FieldMapping{
			{"Bill First Name", "firstname", DoNothing()},
			{"Bill Last Name", "lastname", DoNothing()},
			{"Bill Company", "company", DoNothing()},
			{"Bill Street Address Line 1", "address_1", DoNothing()},
			{"Bill Street Address Line 2", "address_2", DoNothing()},
			{"Bill City", "city", DoNothing()},
			{"Bill Post Code", "postcode", DoNothing()},
		},
	}
}
