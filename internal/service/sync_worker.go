package service

import (
	"context"
	"github.com/sunquan03/cms_api/internal/models"
	"github.com/sunquan03/cms_api/internal/repository/elastic"
	"log"
	"sync"
)

type SyncWorker struct {
	ID       int
	syncChan chan *models.ContentSync
	elastic  *elastic.ElasticLayer
}

func NewSyncWorker(id int, syncChan chan *models.ContentSync, eslayer *elastic.ElasticLayer) *SyncWorker {
	return &SyncWorker{
		ID:       0,
		syncChan: syncChan,
		elastic:  eslayer,
	}
}

func (s *SyncWorker) Start(wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		for syncTask := range s.syncChan {
			switch syncTask.Operation {
			case models.CreateELK:
				err := s.elastic.CreateContent(context.TODO(), syncTask)
				if err != nil {
					log.Println(err)
				}
			case models.UpdateELK:
				err := s.elastic.UpdateContent(context.TODO(), syncTask)
				if err != nil {
					log.Println(err)
				}
			case models.DeleteELK:
				err := s.elastic.DeleteContent(context.TODO(), syncTask)
				if err != nil {
					log.Println(err)
				}

			}

		}
	}()
}
