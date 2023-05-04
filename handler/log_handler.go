package handler

import (
	"context"

	"es_log/controller"
	"es_log/gen-go/es_log"
)

type LogHandler struct{}

func (h *LogHandler) Search(ctx context.Context, req *es_log.SearchRequest) (*es_log.SearchResponse, error) {
	logController := controller.NewLogController()
	return logController.Search(ctx, req)
}
