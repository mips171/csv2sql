package main

type Entity interface {
	GetValue(fieldName string) interface{}
}
type FieldMapping struct {
	CSVColumnName   string
	DBColumnName    string
	MappingFunction func(Entity) interface{}
}

type TableMapping struct {
	TableName   string
	ColumnOrder []string
	Fields      []FieldMapping
}

func TransformIdentity(value string) string {
	return value
}
