package model

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

type LogModel struct{}

func NewLogModel() *LogModel {
	return &LogModel{}
}

func (m *LogModel) Search(ctx context.Context, index string, queryBuilder *QueryBuilder) (int64, map[string]interface{}, error) {
	var buf bytes.Buffer
	searchBody := map[string]interface{}{}

	boolQuery := map[string]interface{}{}

	if len(queryBuilder.Must) > 0 {
		must := []map[string]interface{}{}
		for field, values := range queryBuilder.Must {
			must = append(must, map[string]interface{}{
				"terms": map[string]interface{}{
					field: values,
				},
			})
		}
		boolQuery["must"] = must
	}

	if len(queryBuilder.MustNot) > 0 {
		mustNot := []map[string]interface{}{}
		for field, values := range queryBuilder.MustNot {
			mustNot = append(mustNot, map[string]interface{}{
				"terms": map[string]interface{}{
					field: values,
				},
			})
		}
		boolQuery["must_not"] = mustNot
	}

	if queryBuilder.TimeRange.Field != "" {
		boolQuery["filter"] = []map[string]interface{}{
			{
				"range": map[string]interface{}{
					queryBuilder.TimeRange.Field: map[string]interface{}{
						"gte": queryBuilder.TimeRange.GTE,
						"lte": queryBuilder.TimeRange.LTE,
					},
				},
			},
		}
	}

	if len(queryBuilder.Should) > 0 {
		boolQuery["minimum_should_match"] = len(queryBuilder.Should)
		if queryBuilder.ShouldOr {
			boolQuery["minimum_should_match"] = 1
		}

		should := make([]map[string]interface{}, 0)
		for field, values := range queryBuilder.Should {
			should = append(should, map[string]interface{}{
				"terms": map[string]interface{}{
					field: values,
				},
			})
		}
		boolQuery["should"] = should
	}

	if len(queryBuilder.Sort) > 0 {
		sort := make([]map[string]interface{}, 0)
		for field, value := range queryBuilder.Sort {
			sort = append(sort, map[string]interface{}{
				field: map[string]interface{}{
					"order": value,
				},
			})
		}
	}

	searchBody["query"] = map[string]interface{}{"bool": boolQuery}

	if err := json.NewEncoder(&buf).Encode(searchBody); err != nil {
		return 0, nil, fmt.Errorf("error encoding search body: %w", err)
	}


	resp, err := esClient.Search(
		esClient.Search.WithContext(ctx),
		esClient.Search.WithBody(&buf),
		esClient.Search.WithIndex(index),
		esClient.Search.WithTrackTotalHits(true),
		esClient.Search.WithFrom(int((queryBuilder.PageNum - 1) * queryBuilder.PageSize)),
		esClient.Search.WithSize(int(queryBuilder.PageSize)),
	)

	if err != nil {
		return 0, nil, fmt.Errorf("error executing search request: %w", err)
	}
	defer resp.Body.Close()

	var esResp map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&esResp); err != nil {
		return 0, nil, fmt.Errorf("error decoding search response: %w", err)
	}
	totalHits := int64(esResp["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))
	return totalHits, esResp, nil
}
