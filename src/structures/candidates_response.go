package structures

type CandidateResponse struct {
	Data        []CandidateData `json:"data"`
	PageTotal   int    `json:"page_total"`
	Offset      int    `json:"offset"`
	Previous    string `json:"previous"`
	Next        string `json:"next"`
	First       string `json:"first"`
	Last        string `json:"last"`
	Total       int    `json:"total"`
}