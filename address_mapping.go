package main

func GetAddressMapping() TableMapping {
	return TableMapping{
		TableName:   "oc_address",
		ColumnOrder: []string{"customer_id", "firstname", "lastname", "company", "address_1", "address_2", "city", "postcode", "country_id", "zone_id"},
		Fields: []FieldMapping{
			{"", "firstname", JustUse("BillFirstName")}, // use firstname as the default, email as backup
			{"", "lastname", JustUse("BillLastName")},   // use lastname as the default, email as backup
			{"BillCompany", "company", JustUse("BillCompany")},
			{"BillAddress", "address_1", JustUse("BillAddress")},
			{"", "address_2", JustUse("BillAddress2")},
			{"BillCity", "city", JustUse("BillCity")},
			{"BillPostCode", "postcode", JustUse("BillPostCode")},
			{"BillCountry", "country_id", MapCountryToCode("BillCountry")},
			{"", "zone_id", MapAustralianPostCodeToStateZoneID("BillPostCode")},
		},
	}
}
