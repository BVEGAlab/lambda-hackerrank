package lambda_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"

	// "strings"
	"testing"

	//"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"hacker-rank-lambda/lambda"
	"hacker-rank-lambda/src/get_courses"
)

type MockGetCourses struct {
    mock.Mock
}

func (m *MockGetCourses) Get_courses() ([]string, string, error) {
    args := m.Called()
    return args.Get(0).([]string), args.String(1), args.Error(2)
}

func TestGetCourses(t *testing.T) {
    // mockCourses := []string{"1644325", "1644302"}
    mockResponses := []string{
        `{
            "data":[
                {
                    "duration":40,
                    "id":"1644325",
                    "locked":false,
                    "name":"Banco de Bogotá QA Engineer funcional",
                    "questions":[
                        "1467515"
                    ],
                    "state":"active",
                    "tags":[
                        "QA",
                        "Test"
                    ],
                    "unique_id":"5qqjbgkh5sn"
                },
                {
                    "duration":20,
                    "id":"1644302",
                    "locked":false,
                    "name":"Banco de Bogotá QA Engineer Automation",
                    "questions":[
                        "1577920"
                    ],
                    "state":"active",
                    "tags":[
                        
                    ],
                    "unique_id":"c08pkhb2oj4"
                }
            ]
        }`,
        `{
            "data":[
                
            ]
        }`,
    }

    var offset int
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(mockResponses[offset/10]))
        offset += 10
    }))
    defer ts.Close()

    get_courses.URL = ts.URL // Set the URL variable to the mock server URL
	ctx := context.Background()
	response2, _ := lambda.HandleRequest(ctx)
	fmt.Print(response2)
    //courses, response, err := get_courses.Get_courses()
    //assert.NoError(t, err)
	//fmt.Println(courses)
	//fmt.Println(response)
    //assert.Equal(t, strings.Join(mockResponses, ""), response)
}

// in development