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

func getSort(acs bool) string {
	if acs {
		return "acs"
	}
	return "desc"
}

func (c *LogController) Search(ctx context.Context, req *es_log.SearchRequest) (*es_log.SearchResponse, error) {
	logModel := model.NewLogModel()

	searchBody := &model.QueryBuilder{
		PageNum: req.GetPage(),
		PageSize: req.GetPageSize(),
		Must:     make(map[string][]string),
		MustNot:  make(map[string][]string),
		Should:   make(map[string][]string),
		ShouldOr: false,
		Sort: map[string]string{
			"timestamp": getSort(req.GetAcsSort()),
		},
	}

	// timerange
	if req.StartTime != nil && req.EndTime != nil {
		searchBody.TimeRange = struct {
			Field string
			GTE   string
			LTE   string
		}{Field: "timestamp", GTE: req.GetStartTime(), LTE: req.GetEndTime()}
	}

	// must
	if req.LogLevels != nil {
		searchBody.Must["log_level"] = req.GetLogLevels()
	}

	// must not
	if req.FilterWords != nil {
		searchBody.MustNot["log_message"] = req.GetFilterWords()
	}

	//should
	if req.KeyWords != nil {
		searchBody.Should["log_message"] = req.GetKeyWords()
		searchBody.ShouldOr = req.GetKeyWordsOr()
	}

	total, searchResp, err := logModel.Search(ctx, "log_index", searchBody)
	if err != nil {
		return nil, err
	}
	return &es_log.SearchResponse{
		Total: total,
		Hits: searchResp["hits"].([]map[string]string),
	}, nil
}
