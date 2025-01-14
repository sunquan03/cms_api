package elastic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sunquan03/cms_api/internal/models"
)

func (l *ElsaticLayer) CreateContentTypeIndex(contentType *models.ContentType) error {
	idx_name := fmt.Sprintf("idx_%s", contentType.Name)
	idx_mapping, err := generateIndexMapping(contentType.Name, contentType.Fields)
	if err != nil {
		return err
	}

	res, err := l.client.Indices.Create(idx_name, l.client.Indices.Create.WithBody(bytes.NewReader(idx_mapping)))
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("create index %s", idx_name))
	}
	defer res.Body.Close()

	if res.IsError() {
		return errors.Wrap(err, fmt.Sprintf("create index %s", idx_name))
	}

	return nil
}

func generateIndexMapping(contentName string, fields []models.Field) ([]byte, error) {
	settings := map[string]int{
		"number_of_shards":   1,
		"number_of_replicas": 1,
	}
	properties := make(map[string]models.FieldMapping)
	for _, field := range fields {
		fieldMapping := models.FieldMapping{
			Type: field.Type,
		}

		if fieldMapping.Type == "date" {
			fieldMapping.Format = "yyyy-MM-dd"
		}

		properties[field.Name] = fieldMapping
	}

	mapping := models.IndexMapping{
		Settings: settings,
		Mappings: map[string]models.MappingFields{
			"properties": {
				Properties: properties,
			},
		},
	}

	jsonData, err := json.MarshalIndent(mapping, "", "  ")
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(fmt.Sprintf("marshal index mapping for type %s", contentName)))
	}

	return jsonData, nil
}
