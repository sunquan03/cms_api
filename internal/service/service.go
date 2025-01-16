package service

import (
	"github.com/sunquan03/cms_api/internal/repository/elastic"
	"github.com/sunquan03/cms_api/internal/repository/postgres"
)

type Service struct {
	elasticLayer  *elastic.ElasticLayer
	postgresLayer *postgres.PostgresLayer
}

func NewService(elasticLayer *elastic.ElasticLayer, postgresLayer *postgres.PostgresLayer) *Service {
	return &Service{elasticLayer: elasticLayer, postgresLayer: postgresLayer}
}
