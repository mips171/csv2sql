package main

type CustomerRecord struct {
	Email     string `csv:"Email"`
	FirstName string `csv:"First Name"`
	LastName  string `csv:"Last Name"`
	Address   string `csv:"Address"`
	City      string `csv:"City"`
	Phone     string `csv:"Phone"`
	//... any other fields your CSV might have.
}

// func GetCustomerMapping() TableMapping {
// 	return TableMapping{
// 		TableName:   "oc_customer",
// 		ColumnOrder: []string{"customer_id", "customer_group_id", "firstname", "lastname", "email", "telephone", "newsletter", "status", "date_added"},
// 		Fields: []FieldMapping{
// 			{"User Group", "customer_group_id", func(value string, _ string) interface{} { return GetUserGroupID(value) }},
// 			{"Bill First Name", "firstname", GetFirstName},
// 			{"Bill Last Name", "lastname", GetLastName},
// 			{"Email Address", "email", DoNothing("Email")},
// 			{"Bill Phone", "telephone", DoNothing("Telephone")},
// 			{"Newsletter Subscriber", "newsletter", func(value string, _ string) interface{} { return GetNewsletterStatus(value) }},
// 			{"Active", "status", func(value string, _ string) interface{} { return GetStatus(value) }},
// 			{"", "date_added", func(_ string, _ string) interface{} { return GetDateAdded() }},
// 		},
// 	}
// }
