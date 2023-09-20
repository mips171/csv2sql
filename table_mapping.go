package main

type FieldMapping struct {
	CsvFieldName   string
	DbColumnName   string
	Transformation func(string, string) interface{}
}

type TableMapping struct {
	TableName   string
	ColumnOrder []string
	Fields      []FieldMapping
}

func TransformIdentity(value string) string {
	return value
}
