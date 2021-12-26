package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	myhttp "go-otoklix-blog/delivery/http"
	"go-otoklix-blog/infra/common"
	"go-otoklix-blog/infra/config"
)

var BASE_URL = "http://localhost:" + config.Get("HTTP_PORT") + common.API_CTX

func TestBlogPost(t *testing.T) {
	tests := []struct {
		description string

		// Test input
		method      string
		route       string
		requestBody map[string]string

		// Expected output
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "CreateOne",
			method:        "POST",
			route:         BASE_URL + "/BlogPost/",
			requestBody:   map[string]string{"title": "post-01", "content": "content post-01"},
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "",
		},
		{
			description:   "UpdateOne",
			method:        "PUT",
			route:         BASE_URL + "/BlogPost/1",
			requestBody:   map[string]string{"title": "post-01 OK", "content": "content post-01 OK"},
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "",
		},
		{
			description:   "FindAll",
			method:        "GET",
			route:         BASE_URL + "/BlogPost/",
			requestBody:   nil,
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "",
		},
		{
			description:   "FindById",
			method:        "GET",
			route:         BASE_URL + "/BlogPost/1",
			requestBody:   nil,
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "",
		},
		{
			description:   "DeleteOne",
			method:        "DELETE",
			route:         BASE_URL + "/BlogPost/1",
			requestBody:   nil,
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "",
		},
		{
			description:   "non existing Endpoint",
			method:        "GET",
			route:         "/i-dont-exist",
			requestBody:   nil,
			expectedError: false,
			expectedCode:  404,
			expectedBody:  "Cannot GET /i-dont-exist",
		},
	}

	db, err := sql.Open("sqlite3", config.Get("DB_NAME"))
	if err != nil {
		fmt.Println(err)
	}
	_, err = db.Exec("UPDATE sqlite_sequence SET seq='1' WHERE name='blog_post'")
	if err != nil {
		fmt.Println(err)
	}

	// Setup the app as it is done in the main function
	app := fiber.New()
	myhttp.Routes(app)

	// Iterate through test single test cases
	for _, test := range tests {
		jsonValue, _ := json.Marshal(test.requestBody)
		req := httptest.NewRequest(
			test.method,
			test.route,
			bytes.NewBuffer(jsonValue))

		req.Header.Set("Content-Type", "application/json")

		// Perform the request plain with the app.
		// The -1 disables request latency.
		res, err := app.Test(req, -1)

		// verify that no error occured, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		// As expected errors lead to broken responses, the next
		// test case needs to be processed
		if test.expectedError {
			continue
		}

		// Verify if the status code is as expected
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		// Reading the response body should work everytime, such that
		// the err variable should be nil
		assert.Nilf(t, err, test.description)
	}
}
