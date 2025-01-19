package cache

import (
	"context"
	"fmt"
)

func (r *RedisCache) SetSearchableField(ctx context.Context, contentType string, field string) error {
	key := fmt.Sprintf("%s:searchable", contentType)

	err := r.client.LPush(ctx, key, field).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisCache) GetSearchableFieldsList(ctx context.Context, contentType string) ([]string, error) {
	key := fmt.Sprintf("%s:searchable", contentType)
	res := r.client.LRange(ctx, key, 0, -1)
	if res.Err() != nil {
		return nil, res.Err()
	}

	fields, err := res.Result()
	if err != nil {
		return nil, err
	}
	return fields, nil
}
