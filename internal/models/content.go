package models

const (
	UpdateELK = "update"
	CreateELK = "create"
	DeleteELK = "delete"
)

type ContentSync struct {
	ID          int64
	Operation   string
	ContentType string
	Payload     map[string]interface{}
}

func NewContentSync(id int64, operation, contentType string, payload map[string]interface{}) *ContentSync {
	return &ContentSync{
		ID:          id,
		Operation:   operation,
		ContentType: contentType,
		Payload:     payload,
	}
}
