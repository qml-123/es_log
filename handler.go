package main

import (
	"context"
	"log"

	"github.com/qml-123/es_log/controller"
	es_log "github.com/qml-123/es_log/kitex_gen/es_log"
)

// LogServiceImpl implements the last service interface defined in the IDL.
type LogServiceImpl struct{}

// Search implements the LogServiceImpl interface.
func (s *LogServiceImpl) Search(ctx context.Context, req *es_log.SearchRequest) (resp *es_log.SearchResponse, err error) {
	// TODO: Your code here...
	log.Printf("req params: %v", req)
	logController := controller.NewLogController()
	return logController.Search(ctx, req)
}
