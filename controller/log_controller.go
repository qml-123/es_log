package controller

import (
	"context"

	"github.com/qml-123/es_log/gen-go/es_log"
	"github.com/qml-123/es_log/model"
)

type LogController struct{}

func NewLogController() *LogController {
	return &LogController{}
}

func (c *LogController) Search(ctx context.Context, req *es_log.SearchRequest) (*es_log.SearchResponse, error) {
	logModel := model.NewLogModel()
	return logModel.Search(ctx, req)
}
