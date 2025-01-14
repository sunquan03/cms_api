package service

import "github.com/sunquan03/cms_api/internal/models"

func (s *Service) CreateContentType(contentType *models.ContentType) error {
	err := s.postgresLayer.GenerateContentTypeTable(contentType)
	if err != nil {
		return err
	}

	err = s.elasticLayer.CreateContentTypeIndex(contentType)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteContentType(name string) error {

	return nil
}

func (s *Service) GetContentTypesList() ([]models.ContentType, error) {
	return nil, nil
}
