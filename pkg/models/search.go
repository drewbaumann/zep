package models

type SearchResult struct {
<<<<<<< Updated upstream
	Message *Message               `json:"message"`
	Summary *Summary               `json:"summary"`        // reserved for future use
	Meta    map[string]interface{} `json:"meta,omitempty"` // reserved for future use
	Dist    float64                `json:"dist"`
}

type SearchPayload struct {
	Text string                 `json:"text"`
	Meta map[string]interface{} `json:"meta,omitempty"` // reserved for future use
=======
	Message  *Message               `json:"message"`
	Summary  *Summary               `json:"summary"` // reserved for future use
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Dist     float64                `json:"dist"`
}

type SearchPayload struct {
	Text     string                 `json:"text"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
>>>>>>> Stashed changes
}
