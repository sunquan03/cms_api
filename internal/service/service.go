package service

import (
	"github.com/sunquan03/cms_api/internal/models"
	"github.com/sunquan03/cms_api/internal/repository/elastic"
	"github.com/sunquan03/cms_api/internal/repository/postgres"
)

type Service struct {
	elasticLayer  *elastic.ElasticLayer
	postgresLayer *postgres.PostgresLayer
	syncChan      chan *models.ContentSync // channel to synchronize data between postgres and elastic
}

func NewService(elasticLayer *elastic.ElasticLayer, postgresLayer *postgres.PostgresLayer) *Service {
	return &Service{elasticLayer: elasticLayer, postgresLayer: postgresLayer}
}
