package service

func (s *Service) CreateContent(contentType string, content map[string]interface{}) (int64, error) {
	return s.postgresLayer.CreateContent(contentType, content)
}

func (s *Service) UpdateContent(contentType string, id int64, content map[string]interface{}) error {
	return s.postgresLayer.UpdateContent(contentType, id, content)
}

func (s *Service) GetContentById(contentType string, id int64) (string, error) {
	return s.postgresLayer.GetContentById(contentType, id)
}

// TODO: func(s *Service) DeleteContent(contentType string, id int64) error {}
