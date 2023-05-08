package model


type SearchResult struct {
	Hits Hits `json:"hits"`
}

type Hits struct {
	Total *Total       `json:"total"`
	Hits  []*HitDetails `json:"hits"`
}

type Total struct {
	Value int `json:"value"`
}

type HitDetails struct {
	Index  string                 `json:"_index"`
	Type   string                 `json:"_type"`
	ID     string                 `json:"_id"`
	Score  float64                `json:"_score"`
	Source map[string]interface{} `json:"_source"`
}
