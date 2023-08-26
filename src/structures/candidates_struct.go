package structures

type CandidateData struct {
    ID                string   `json:"id"`
    Test_id           string   `json:"test_id"`
    Email             string   `json:"email"`
    Full_name         string   `json:"full_name"`
    Attemp_id         string   `json:"attempt_id"`
    Score             float64  `json:"score"`
    Status            int      `json:"status"`
    Attempt_starttime string   `json:"attempt_starttime"`
    Attempt_endtime   string   `json:"attempt_endtime"`
    Invited_on        string   `json:"invited_on"`
    Pdf_url           string   `json:"pdf_url"`
    Score_tags_split  map[string]string `json:"score_tags_split|ScoresTagsSplit"`
    Questions         map[string]float64  `json:"questions"`
    Feedback          string   `json:"feedback"`
    Percentage_score  float64  `json:"percentage_score"`
}