package config

type DataType struct {
	Description       string
	PostgresType      string
	ElasticsearchType string
}

var AllowedDataTypes = map[string]DataType{
	"text": {
		Description:       "Full-text fields for search",
		PostgresType:      "TEXT",
		ElasticsearchType: "text",
	},
	"keyword": {
		Description:       "Exact matches and aggregations",
		PostgresType:      "VARCHAR(500)",
		ElasticsearchType: "keyword",
	},
	"integer": {
		Description:       "Integer numbers",
		PostgresType:      "INTEGER",
		ElasticsearchType: "integer",
	},
	"float": {
		Description:       "Decimal numbers",
		PostgresType:      "REAL",
		ElasticsearchType: "float",
	},
	"boolean": {
		Description:       "True/false values",
		PostgresType:      "BOOLEAN",
		ElasticsearchType: "boolean",
	},
	"date": {
		Description:       "Date and time values",
		PostgresType:      "DATE",
		ElasticsearchType: "date",
	},
	"geo_point": {
		Description:       "Geographical locations",
		PostgresType:      "",
		ElasticsearchType: "geo_point",
	},
	"nested": {
		Description:       "Nested JSON objects",
		PostgresType:      "JSONB",
		ElasticsearchType: "nested",
	},
	"array": {
		Description:       "Array of values",
		PostgresType:      "JSONB",
		ElasticsearchType: "array",
	},
}
