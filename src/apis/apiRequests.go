package apis

import (
	"fmt"
	"hacker-rank-lambda/src/headers"
	"net/http"
	"strings"
)

var candidatesURL = "https://www.hackerrank.com/x/api/v3/tests/{test_id}/candidates"
var testURL = "https://www.hackerrank.com/x/api/v3/tests/{test_id}"
var attemptURL = "https://www.hackerrank.com/x/api/v1/tests/{test_id}/attempts/{attempt_id}"

func GetCandidates(test_id string, offset int) (*http.Request, error) {
	const limit = 10
    candidatesNewUrl := strings.Replace(candidatesURL, "{test_id}", test_id, 1)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?limit=%d&offset=%d", candidatesNewUrl, limit, offset), nil)
    if err != nil {
        return nil, err
    }
    get_headers.SetHeaders(req)
    return req, nil
}

func GetTest(test_id string) (*http.Request, error) {
	testNewUrl := strings.Replace(testURL, "{test_id}", test_id, 1)
	req, err := http.NewRequest("GET", fmt.Sprintf(testNewUrl), nil)
	if err != nil {
		return nil, err
	}
	get_headers.SetHeaders(req)
	return req, nil
}

func GetCandidateAttemptScores(test_id string, attempt_id string) (*http.Request, error) {
    attemptNewURL := strings.Replace(attemptURL, "{test_id}", test_id, 1)
    attemptNewURL = strings.Replace(attemptNewURL, "{attempt_id}", attempt_id, 1)
    req, err := http.NewRequest("GET", attemptNewURL, nil)
    if err != nil {
        return nil, err
    }
    get_headers.SetHeaders(req)

    return req, nil
}