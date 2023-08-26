package middleware

import (
	"fmt"
	config "hacker-rank-lambda/src"
	"hacker-rank-lambda/src/services"
	"hacker-rank-lambda/src/structures"
	"time"

    _ "github.com/lib/pq"
)

func GetCandidates(tests []structures.Tests) ([]structures.CandidateResponse, error) {
    var allCandidatesResponses []structures.CandidateResponse
    for _, test := range tests {
        candidatesResponse, err := services.GetCandidates(test.ID)
        if err != nil {
            fmt.Printf("Error getting candidates for test %s: %s\n", test.ID, err)
            continue
        }
        if candidatesResponse.Data == nil {
            fmt.Printf("Skipping test %s because candidates data is null\n", test.ID)
            continue
        }
        allCandidatesResponses = append(allCandidatesResponses, candidatesResponse)
    }

    return allCandidatesResponses, nil
}

func GetTest() ([]structures.Tests, error) {
    fmt.Println("Test to search: ", config.TESTS_TO_SEARCH)
    var testReponse []structures.Tests
    for _, test := range config.TESTS_TO_SEARCH {
        test, err := services.GetTest(test)
        if err != nil {
            fmt.Println("Error getting test:", err)
            return nil, fmt.Errorf("Error getting test: %v", err)
        }
        testReponse = append(testReponse, test)
        time.Sleep(3 * time.Second)
    }

    return testReponse, nil
}

func GenerateSQL(tests []structures.Tests, allCandidatesResponses []structures.CandidateResponse) error {

    err := services.GenerateTestSQL(tests)
    if err != nil {
        panic(err)
    }

    err = services.GenerateCandidateSQL(allCandidatesResponses)
    if err != nil {
        panic(err)
    }

    err = services.GenerateCandidateTestSQL(allCandidatesResponses)
    if err != nil {
        panic(err)
    }

    return nil
}


func InsertsShToDatabase() error {
    filepaths := []string{"Tests.sh", "Candidate.sh", "Candidate_test.sh",  "Tag.sh", "Cross_tag.sh", }

    db, err := services.ConnectToDatabase()
    if err != nil {
        return err
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        return err
    }

    for _, filepath := range filepaths {
        fmt.Println("Inserting data from", filepath)
        err = services.InsertData(db, filepath)

        if err != nil {
            fmt.Printf("Error inserting data from %s: %s\n", filepath, err)
            continue
        }
    }

    return nil

}