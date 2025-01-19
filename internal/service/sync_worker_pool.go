package service

import (
	"github.com/sunquan03/cms_api/internal/models"
	"github.com/sunquan03/cms_api/internal/repository/elastic"
	"sync"
)

type SWPool struct {
	workersCap int
	syncChan   chan *models.ContentSync
	wg         sync.WaitGroup
	elastic    *elastic.ElasticLayer
}

func NewSWPool(workersCap int, syncChan chan *models.ContentSync, eslayer *elastic.ElasticLayer) *SWPool {
	return &SWPool{
		workersCap: workersCap,
		syncChan:   syncChan,
		wg:         sync.WaitGroup{},
		elastic:    eslayer,
	}
}

func (s *SWPool) Run() {
	for i := 0; i < s.workersCap; i++ {
		worker := NewSyncWorker(i, s.syncChan, s.elastic)
		worker.Start(&s.wg)
	}
	s.wg.Wait()
}
