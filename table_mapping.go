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
	ColumnOrder []string
	Fields      []FieldMapping
	TableName   string
}

func TransformIdentity(value string) string {
	return value
}
