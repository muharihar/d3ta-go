package schema

// ESIndexSearchResponse represent ESIndexSearchResponse
type ESIndexSearchResponse struct {
	Took     int64     `json:"took"`
	TimedOut bool      `json:"timed_out"`
	Shards   *ESShards `json:"_shards"`
	// Hits     *ESHits   `json:"hits"` // Depend On Custom Source
}

// ESHits represent ESHits
type ESHits struct {
	Total    *ESTotal    `json:"total"`
	MaxScore interface{} `json:"max_score"`
	// Hits     []*ESHit      `json:"hits"` // Depend On Custom Source
}

// ESHit represent ESHit
type ESHit struct {
	Index string      `json:"_index"`
	Type  string      `json:"_type"`
	ID    string      `json:"_id"`
	Score interface{} `json:"_score"`
	// Source Source      `json:"_source"` // Custom Source (define as we need)
	Sort []interface{} `json:"sort"`
}

// ESTotal represent ESTotal
type ESTotal struct {
	Value    int64  `json:"value"`
	Relation string `json:"relation"`
}

// ESShards represent ESShards
type ESShards struct {
	Total      int64 `json:"total"`
	Successful int64 `json:"successful"`
	Skipped    int64 `json:"skipped"`
	Failed     int64 `json:"failed"`
}
