package controller

import (
	"context"

	"es_log/gen-go/es_log"
	"es_log/model"
)

type LogController struct{}

func NewLogController() *LogController {
	return &LogController{}
}

func (c *LogController) Search(ctx context.Context, req *es_log.SearchRequest) (*es_log.SearchResponse, error) {
	logModel := model.NewLogModel()
	return logModel.Search(ctx, req)
}
