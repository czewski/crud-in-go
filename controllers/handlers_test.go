package controllers

import (
	"net/http"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPlayers(t *testing.T) {
	app := fiber.New()

	app.Get("/player/:id?", ReadPlayers)
	app.Post("/player", CreatePlayers)
	app.Put("/player/:id", UpdatePlayers)
	app.Delete("/player/:id", DeletePlayers)

	tests := []struct {
		description   string
		route         string
		method        string
		body          string
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "create player",
			route:         "/player",
			method:        "POST",
			body:          `{"name":"Gabriel Toledo de Alcântara Sguario","nick": "FalleN","team":"Imperial","game":"CS:GO","age":30}`,
			expectedError: false,
			expectedCode:  201,
		},
		{
			description:   "fail endpoint",
			route:         "/",
			method:        "POST",
			body:          ``,
			expectedError: false,
			expectedCode:  404,
		},
		// {
		// 	description:   "delete player",
		// 	route:         "/player/Gabriel",
		// 	method:        "DELETE",
		// 	body:          ``,
		// 	expectedError: false,
		// 	expectedCode:  404,
		// },
	}
	for _, test := range tests {
		//payload, _ := json.Marshal(test.body)
		req, _ := http.NewRequest(
			test.method,
			test.route,
			strings.NewReader(test.body),
		)
		req.Header.Add("Content-Type", "application/json")

		res, err := app.Test(req, -1)
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		//Verify if the status code is as expected
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		// Reading the response body should work everytime, such that
		// the err variable should be nil
		assert.Nilf(t, err, test.description)

		// Verify, that the reponse body equals the expected body
	}
}
