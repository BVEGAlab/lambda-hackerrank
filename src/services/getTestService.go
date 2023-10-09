package services

import (
	"encoding/json"
	"fmt"
	"hacker-rank-lambda/src/apis"
	"hacker-rank-lambda/src/structures"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func GetTest(test_id string) (structures.Tests, error) {

	client := &http.Client{}

	req, err := apis.GetTest(test_id)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	
	var testResponse structures.Tests

	err = json.Unmarshal(bodyBytes, &testResponse)
	if err != nil {
		panic(err)
	}

    today := time.Now().Format("02-01-2006")
    folderName := fmt.Sprintf("./%s/json_test", today)
    err = os.MkdirAll(folderName, 0755)
    if err != nil {
        panic(err)
    }

    fileName := fmt.Sprintf("%s/%s.json", folderName, test_id)
    file, err := os.Create(fileName)
    if err != nil {
        panic(err)
    }
    defer file.Close()


	_, err = file.Write(bodyBytes)
                if err != nil {
                    panic(err)
                }

	return testResponse, nil
}