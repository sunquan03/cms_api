package service

import (
	"context"
	"github.com/sunquan03/cms_api/internal/models"
)

func (s *Service) CreateContent(contentType string, content map[string]interface{}) (int64, error) {
	id, err := s.postgresLayer.CreateContent(contentType, content)
	if err != nil {
		return 0, err
	}

	s.syncChan <- models.NewContentSync(id, models.CreateELK, contentType, content)

	return id, nil
}

func (s *Service) UpdateContent(contentType string, id int64, content map[string]interface{}) error {
	err := s.postgresLayer.UpdateContent(contentType, id, content)
	if err != nil {
		return err
	}

	s.syncChan <- models.NewContentSync(id, models.UpdateELK, contentType, content)

	return nil
}

func (s *Service) GetContentById(contentType string, id int64) (string, error) {
	return s.postgresLayer.GetContentById(contentType, id)
}

func (s *Service) DeleteContent(contentType string, id int64) error {
	err := s.postgresLayer.DeleteContent(contentType, id)
	if err != nil {
		return err
	}

	s.syncChan <- models.NewContentSync(id, models.DeleteELK, contentType, nil)
	return nil
}

func (s *Service) SearchContentByQuery(ctx context.Context, contentType string, searchQuery string) (map[string]interface{}, error) {
	fields, err := s.redisLayer.GetSearchableFieldsList(ctx, contentType)
	if err != nil {
		return nil, err
	}
	
	return s.elasticLayer.SearchContentByQuery(ctx, contentType, searchQuery, fields)
}
