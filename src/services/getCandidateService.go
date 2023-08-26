package services

import (
	"encoding/json"
	"hacker-rank-lambda/src/apis"
	"hacker-rank-lambda/src/structures"
	"io/ioutil"
	"net/http"
	"time"
)

func GetCandidates(test_id string) (structures.CandidateResponse, error) {
    offset := 0


    client := &http.Client{}

    req, err := apis.GetCandidates(test_id, offset)
    if err != nil {
        panic(err)
    }


    var candidates []structures.CandidateData

    for {
		time.Sleep(2 * time.Second)
        resp, err := client.Do(req)
        if err != nil {
            panic(err)
        }
        defer resp.Body.Close()

        bodyBytes, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            panic(err)
        }

        var candidatesResponse structures.CandidateResponse

        err = json.Unmarshal(bodyBytes, &candidatesResponse)
        if err != nil {
            panic(err)
        }

        for _, candidate := range candidatesResponse.Data {
            if candidate.Status == 7 {

				attemtReq, err := apis.GetCandidateAttemptScores(test_id, candidate.Attemp_id)
				if err != nil {
					panic(err)
				}
				time.Sleep(1 * time.Second)
				attemptResp, err := client.Do(attemtReq)
				if err != nil {
					panic(err)
				}
				defer attemptResp.Body.Close()

				attemptBodyBytes, err := ioutil.ReadAll(attemptResp.Body)
				if err != nil {
					panic(err)
				}

				var attemptResponse structures.AttemptResponse

				err = json.Unmarshal(attemptBodyBytes, &attemptResponse)
				if err != nil {
					panic(err)
				}
			
                candidates = append(candidates, structures.CandidateData{
                    ID:                candidate.ID,
                    Test_id:           test_id,
                    Email:             candidate.Email,
                    Full_name:         candidate.Full_name,
                    Attemp_id:         candidate.Attemp_id,
                    Score:             candidate.Score,
                    Status:            candidate.Status,
                    Attempt_starttime: candidate.Attempt_starttime,
                    Attempt_endtime:   candidate.Attempt_endtime,
                    Invited_on:        candidate.Invited_on,
                    Pdf_url:           candidate.Pdf_url,
                    Score_tags_split:  attemptResponse.Model.ScoresTagsSplit,
                    Questions:         candidate.Questions,
                    Feedback:          candidate.Feedback,
                    Percentage_score:  candidate.Percentage_score,
                })
            }
        }

        if len(candidatesResponse.Data) == 0 {
            break
        }

        offset += 10
        req, err = apis.GetCandidates(test_id, offset)
        if err != nil {
            panic(err)
        }
    }

    responseData := structures.CandidateResponse{
        Data: candidates,
    }

    return responseData, nil
}
