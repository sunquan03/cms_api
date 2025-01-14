package models

type ContentType struct {
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Required   bool   `json:"required"`
	Searchable bool   `json:"searchable"`
	Filterable bool   `json:"filterable"`
}
