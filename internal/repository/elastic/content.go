package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/pkg/errors"
	"github.com/sunquan03/cms_api/internal/models"
	"strconv"
)

func (l *ElasticLayer) CreateContent(ctx context.Context, contentSync *models.ContentSync) error {
	indexName := fmt.Sprintf("idx_%s", contentSync.ContentType)

	err := l.CheckContentIndexExists(ctx, indexName)
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(contentSync.Payload)
	if err != nil {
		return err
	}

	idxReq := esapi.IndexRequest{
		Index:      indexName,
		DocumentID: strconv.FormatInt(contentSync.ID, 10),
		Body:       bytes.NewReader(jsonData),
		Refresh:    "true",
	}

	idxResp, err := idxReq.Do(ctx, l.client)
	if err != nil {
		return err
	}
	defer idxResp.Body.Close()

	if idxResp.IsError() {
		return errors.New(idxResp.String())
	}

	return nil
}

func (l *ElasticLayer) UpdateContent(ctx context.Context, contentSync *models.ContentSync) error {
	indexName := fmt.Sprintf("idx_%s", contentSync.ContentType)

	err := l.CheckContentIndexExists(ctx, indexName)
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(contentSync.Payload)
	if err != nil {
		return err
	}

	updReq := esapi.UpdateRequest{
		Index:      indexName,
		DocumentID: strconv.FormatInt(contentSync.ID, 10),
		Body:       bytes.NewReader(jsonData),
		Refresh:    "true",
	}

	updResp, err := updReq.Do(ctx, l.client)
	if err != nil {
		return err
	}
	defer updResp.Body.Close()

	if updResp.IsError() {
		return errors.New(updResp.String())
	}

	return nil
}

func (l *ElasticLayer) DeleteContent(ctx context.Context, contentSync *models.ContentSync) error {
	indexName := fmt.Sprintf("idx_%s", contentSync.ContentType)

	err := l.CheckContentIndexExists(ctx, indexName)
	if err != nil {
		return err
	}

	delReq := esapi.DeleteRequest{
		Index:      indexName,
		DocumentID: strconv.FormatInt(contentSync.ID, 10),
		Refresh:    "true",
	}

	delResp, err := delReq.Do(ctx, l.client)
	if err != nil {
		return err
	}
	defer delResp.Body.Close()

	if delResp.IsError() {
		return errors.New(delResp.String())
	}

	return nil
}

func (l *ElasticLayer) CheckContentIndexExists(ctx context.Context, indexName string) error {

	existReq := esapi.IndicesExistsRequest{
		Index: []string{indexName},
	}

	existResp, err := existReq.Do(ctx, l.client)
	if err != nil {
		return err
	}
	defer existResp.Body.Close()

	if existResp.IsError() {
		return errors.New(existResp.String())
	}

	if existResp.StatusCode == 404 {
		return errors.New(fmt.Sprintf("index %s not exists", indexName))
	} else if existResp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("create index failed with status %d : %s", existResp.StatusCode, existResp.String()))
	}

	return nil
}
