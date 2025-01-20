package postgres

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sunquan03/cms_api/internal/config"
	"github.com/sunquan03/cms_api/internal/models"
	"strings"
)

func (l *PostgresLayer) GenerateContentTypeTable(contentType *models.ContentType) error {
	var builder strings.Builder
	builder.WriteString("CREATE TABLE IF NOT EXISTS ")
	builder.WriteString(fmt.Sprintf("tb_%s (\n", contentType.Name))

	last_idx := len(contentType.Fields) - 1
	for i, field := range contentType.Fields {

		builder.WriteString(fmt.Sprintf("\t%s %s", field.Name, config.AllowedDataTypes[field.Type].PostgresType))
		if i != last_idx {
			builder.WriteString(",\n")
		} else {
			builder.WriteString("is_deleted int default 0\n);")
		}
	}

	query := builder.String()
	_, err := l.db.Exec(query)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("create table `%s`", contentType.Name))
	}

	return nil
}
