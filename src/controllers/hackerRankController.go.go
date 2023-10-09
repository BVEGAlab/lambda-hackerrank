package controllers

import (
	"fmt"
	"hacker-rank-lambda/src/middleware"
	"hacker-rank-lambda/src/structures"
	"os"
	"time"
)

func HackerRankController() error {
    today := time.Now().Format("02-01-2006")
    folderName := fmt.Sprintf("./%s", today)
    err := os.Mkdir(folderName, 0755)
    if err != nil {
        fmt.Println("Error creating folder:", err)
        return fmt.Errorf("Error creating folder: %v", err)
    }
    fmt.Println("Folder created:", folderName)

	var testReponse []structures.Tests

	testReponse, errtest := middleware.GetTest()
	if errtest != nil {
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

	//err = middleware.InsertsShToDatabase()

	//fmt.Println("Inserts.sh executed successfully!")

    return nil
}