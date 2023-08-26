package get_headers

import (
    "net/http"
    "hacker-rank-lambda/src"
)
// SetHeaders sets the required headers for making requests to the HackerRank API
func SetHeaders(req *http.Request) {
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Accept", "application/json")
    req.Header.Set("Connection", "keep-alive")
    req.Header.Set("Cookie", config.TOKEN)
}