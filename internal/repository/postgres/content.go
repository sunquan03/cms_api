package postgres

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

func (l *PostgresLayer) CreateContent(contentType string, content map[string]interface{}) (int64, error) {
	var builder strings.Builder
	var id int64

	columns := make([]string, 0, len(content))
	values := make([]interface{}, 0, len(content))
	placeholders := make([]string, 0, len(content))

	for key, val := range content {
		columns = append(columns, key)
		values = append(values, val)
		placeholders = append(placeholders, fmt.Sprintf("$%d", len(placeholders)+1))
	}

	builder.WriteString(fmt.Sprintf("INSERT INTO tb_%s (", contentType))
	builder.WriteString(strings.Join(columns, ", "))
	builder.WriteString(") VALUES (")
	builder.WriteString(strings.Join(placeholders, ", "))
	builder.WriteString(") RETURNING ID")

	query := builder.String()

	tx, err := l.db.Begin()
	if err != nil {
		return 0, errors.Wrap(err, "Begin transaction")
	}

	err = tx.QueryRow(query, values...).Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, errors.Wrap(err, "Exec statement")
	}

	err = tx.Commit()
	if err != nil {
		return 0, errors.Wrap(err, "Commit transaction")
	}

	return id, nil
}

func (l *PostgresLayer) UpdateContent(contentType string, id int64, content map[string]interface{}) error {
	var builder strings.Builder
	columnPlaceholders := make([]string, 0, len(content))
	values := make([]interface{}, 0, len(content))

	for key, val := range content {
		placeholder := fmt.Sprintf("%s=$%d", key, len(columnPlaceholders)+1)
		columnPlaceholders = append(columnPlaceholders, placeholder)
		values = append(values, val)
	}

	builder.WriteString(fmt.Sprintf("UPDATE tb_%s SET ", contentType))
	builder.WriteString(strings.Join(columnPlaceholders, ", "))
	builder.WriteString(fmt.Sprintf(" WHERE id = $%d", len(columnPlaceholders)+1))

	query := builder.String()
	tx, err := l.db.Begin()
	if err != nil {
		return errors.Wrap(err, "Begin transaction")
	}

	_, err = tx.Exec(query, values...)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "Exec statement")
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "Commit transaction")
	}

	return nil
}

func (l *PostgresLayer) DeleteContent(contentType string, id int64) error {
	query := fmt.Sprintf("UPDATE tb_%s SET is_deleted=1 WHERE id = $1", contentType, id)
	tx, err := l.db.Begin()
	if err != nil {
		return errors.Wrap(err, "Begin transaction")
	}

	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "Exec statement")
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "Commit transaction")
	}

	return nil
}

func (l *PostgresLayer) GetContentById(contentType string, id int64) (string, error) {
	var result string
	query := fmt.Sprintf(" SELECT row_to_json(t) FROM (SELECT * FROM tb_%s WHERE id = $1 and is_deleted=0) t", contentType)
	err := l.db.QueryRowx(query, id).Scan(&result)
	if err != nil {
		return "", errors.Wrap(err, "Query row json")
	}
	return result, nil
}
