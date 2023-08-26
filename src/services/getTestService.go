package services

import (
	"encoding/json"
	"hacker-rank-lambda/src/apis"
	"hacker-rank-lambda/src/structures"
	"io/ioutil"
	"net/http"
)

func GetTest(test_id string) (structures.Tests, error) {

	client := &http.Client{}

	req, err := apis.GetTest(test_id)
	if err != nil {
		panic(err)
	}

	var testResponse structures.Tests

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	
	err = json.Unmarshal(bodyBytes, &testResponse)
	if err != nil {
		panic(err)
	}

	return testResponse, nil
}