package main

func GetAddressMapping() TableMapping {
	return TableMapping{
		TableName:   "oc_address",
		ColumnOrder: []string{"customer_id", "firstname", "lastname", "company", "address_1", "address_2", "city", "postcode"},
		Fields: []FieldMapping{
			{"Bill First Name", "firstname", func(value string, _ string) string { return TransformIdentity(value) }},
			{"Bill Last Name", "lastname", func(value string, _ string) string { return TransformIdentity(value) }},
			{"Bill Company", "company", func(value string, _ string) string { return TransformIdentity(value) }},
			{"Bill Street Address Line 1", "address_1", func(value string, _ string) string { return TransformIdentity(value) }},
			{"Bill Street Address Line 2", "address_2", func(value string, _ string) string { return TransformIdentity(value) }},
			{"Bill City", "city", func(value string, _ string) string { return TransformIdentity(value) }},
			{"Bill Post Code", "postcode", func(value string, _ string) string { return TransformIdentity(value) }},
		},
	}
}
