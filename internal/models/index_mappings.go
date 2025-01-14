package models

type IndexMapping struct {
	Settings map[string]int           `json:"settings"`
	Mappings map[string]MappingFields `json:"mappings"`
}

type MappingFields struct {
	Properties map[string]FieldMapping `json:"properties"`
}

type FieldMapping struct {
	Type   string `json:"type"`
	Format string `json:"format,omitempty"` // Optional field for formats like date
}
