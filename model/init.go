package model

import (
	"github.com/elastic/go-elasticsearch/v8"
)

var esClient *elasticsearch.Client

func init() {
	var err error
	esClient, err = elasticsearch.NewDefaultClient()
	if err != nil {
		panic(err)
	}
}

type QueryBuilder struct {
	PageNum int32
	PageSize int32
	Must      map[string][]string
	MustNot   map[string][]string
	Should    map[string][]string
	ShouldOr  bool
	TimeRange struct {
		Field string
		GTE   string
		LTE   string
	}
	Sort map[string]string
}
