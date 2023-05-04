package model

import (
	"context"

	"github.com/qml-123/es_log/gen-go/es_log"
)

type LogModel struct{}

func NewLogModel() *LogModel {
	return &LogModel{}
}

func (m *LogModel) Search(ctx context.Context, req *es_log.SearchRequest) (*es_log.SearchResponse, error) {
	// 实现您的查询数据库和ES的逻辑
	// ...
}
