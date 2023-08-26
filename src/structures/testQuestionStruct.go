package structures

import (
    "encoding/json"
    "strconv"
)

type Answer []int

func (a *Answer) UnmarshalJSON(data []byte) error {
    var value interface{}
    err := json.Unmarshal(data, &value)
    if err != nil {
        return err
    }

    switch v := value.(type) {
    case float64:
        *a = []int{int(v)}
    case []interface{}:
        ints := make([]int, len(v))
        for i, val := range v {
            if f, ok := val.(float64); ok {
                ints[i] = int(f)
            } else if s, ok := val.(string); ok {
                if i, err := strconv.Atoi(s); err == nil {
                    ints[i] = i
                }
            }
        }
        *a = ints
    default:
        return json.Unmarshal(data, &a)
    }

    return nil
}

type TestQuestions struct {
    ID               string    `json:"id"`
    ID_test          string    `json:"id_test"`
    TestCaseCount    int       `json:"test_case_count"`
    Languages        []string  `json:"languages"`
    UniqueID         string    `json:"unique_id"`
    ProblemStatement string    `json:"problem_statement"`
    Type             string    `json:"type"`
    Owner            string    `json:"owner"`
    CreatedAt        string    `json:"created_at"`
    Status           string    `json:"status"`
    InternalNotes    string    `json:"internal_notes"`
    Name             string    `json:"name"`
    Tags             []string  `json:"tags"`
    MaxScore         float64   `json:"max_score"`
    Answer           Answer    `json:"answer"`
    Options          []string  `json:"options"`
    ArchivedBy       string    `json:"archived_by"`
    RecommendedDuration int    `json:"recommended_duration"`
}