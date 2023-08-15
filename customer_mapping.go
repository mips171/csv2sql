package main

func GetCustomerMapping() TableMapping {
	return TableMapping{
		TableName:   "oc_customer",
		ColumnOrder: []string{"customer_id", "customer_group_id", "firstname", "lastname", "email", "telephone", "newsletter", "status", "date_added"},
		Fields: []FieldMapping{
			{"User Group", "customer_group_id", func(value string, _ string) string { return GetUserGroupID(value) }},
			{"Bill First Name", "firstname", GetFirstName},
			{"Bill Last Name", "lastname", GetLastName},
			{"Email Address", "email", func(value string, _ string) string { return TransformIdentity(value) }},
			{"Bill Phone", "telephone", func(value string, _ string) string { return TransformIdentity(value) }},
			{"Newsletter Subscriber", "newsletter", func(value string, _ string) string { return GetNewsletterStatus(value) }},
			{"Active", "status", func(value string, _ string) string { return GetStatus(value) }},
			{"", "date_added", func(_ string, _ string) string { return GetDateAdded() }},
		},
	}
}
