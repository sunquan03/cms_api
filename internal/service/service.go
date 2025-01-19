package service

import (
	"github.com/sunquan03/cms_api/internal/cache"
	"github.com/sunquan03/cms_api/internal/models"
	"github.com/sunquan03/cms_api/internal/repository/elastic"
	"github.com/sunquan03/cms_api/internal/repository/postgres"
)

type Service struct {
	elasticLayer  *elastic.ElasticLayer
	postgresLayer *postgres.PostgresLayer
	redisLayer    *cache.RedisCache
	syncChan      chan *models.ContentSync // channel to synchronize data between postgres and elastic
}

func NewService(elasticLayer *elastic.ElasticLayer, postgresLayer *postgres.PostgresLayer, redisLayer *cache.RedisCache, syncChan chan *models.ContentSync) *Service {
	return &Service{
		elasticLayer:  elasticLayer,
		postgresLayer: postgresLayer,
		redisLayer:    redisLayer,
		syncChan:      syncChan,
	}
}
