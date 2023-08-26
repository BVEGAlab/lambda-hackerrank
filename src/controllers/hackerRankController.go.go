package controllers

import (
	"fmt"
	"hacker-rank-lambda/src/middleware"
	"hacker-rank-lambda/src/structures"
)

func HackerRankController() error {
    fmt.Println("Controller called")

	var testReponse []structures.Tests

	testReponse, err := middleware.GetTest()
	if err != nil {
		fmt.Println("Error getting test:", err)
		return fmt.Errorf("Error getting test: %v", err)
	}

	fmt.Println("Total Test:", len(testReponse))
	
    candidatesResponse, err := middleware.GetCandidates(testReponse)
    if err != nil {
        fmt.Println("Error getting candidates:", err)
        return fmt.Errorf("Error getting candidates: %v", err)
    }
	
    err = middleware.GenerateSQL(testReponse, candidatesResponse)
    if err != nil {
        fmt.Println("Error generating SQL:", err)
        return fmt.Errorf("Error generating SQL: %v", err)
    }

    fmt.Println("SQL generated successfully!")

	err = middleware.InsertsShToDatabase()

	fmt.Println("Inserts.sh executed successfully!")

    return nil
}