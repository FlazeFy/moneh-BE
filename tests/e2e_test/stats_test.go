package e2etest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"moneh/models"
	"moneh/tests"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ResponseGetMostContext struct {
	Data    []models.StatsContextTotal `json:"data"`
	Message string                     `json:"message"`
	Status  string                     `json:"status"`
}

type TestDataGetMostContext struct {
	TargetCol string
	Module    string
	Message   string
}

// API GET : Get Most Context Clothes
func TestSuccessGetMostContextWithValidData(t *testing.T) {
	var testData = []TestDataGetMostContext{
		// Test Case ID : TC-E2E-ST-001
		{TargetCol: "flow_type", Module: "flows", Message: "Flow fetched"},
		// Test Case ID : TC-E2E-ST-002
		{TargetCol: "pocket_type", Module: "pockets", Message: "Pocket fetched"},
	}

	for _, td := range testData {
		var res ResponseGetMostContext
		url := fmt.Sprintf("http://127.0.0.1:9000/api/v1/%s/most_context/%s", td.Module, td.TargetCol)
		token, _ := tests.TemplatePostBasicLogin(t, nil, nil, "user")

		// Exec
		req, err := http.NewRequest("GET", url, nil)
		assert.NoError(t, err)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		// Prepare Test
		body, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err)
		err = json.Unmarshal(body, &res)
		assert.NoError(t, err)

		// Get Template Test
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.NotEmpty(t, res.Status)
		assert.Equal(t, "success", res.Status)
		assert.NotEmpty(t, res.Message)
		assert.Equal(t, td.Message, res.Message)
		assert.NotNil(t, res.Data)

		for _, dt := range res.Data {
			// Check Object
			assert.NotEmpty(t, dt.Context)
			assert.NotEmpty(t, dt.Total)

			// Check Data Type
			assert.IsType(t, "", dt.Context)
			assert.IsType(t, 0, dt.Total)
		}
	}
}

func TestFailedGetMostContextWithInvalidTargetCol(t *testing.T) {
	var testData = []TestDataGetMostContext{
		// Test Case ID : TC-E2E-ST-003
		{TargetCol: "flow_invalid", Module: "flows"},
		// Test Case ID : TC-E2E-ST-004
		{TargetCol: "pocket_invalid", Module: "pockets"},
	}

	for _, td := range testData {
		var res tests.ResponseSimple
		url := fmt.Sprintf("http://127.0.0.1:9000/api/v1/%s/most_context/%s", td.Module, td.TargetCol)
		token, _ := tests.TemplatePostBasicLogin(t, nil, nil, "user")

		// Exec
		req, err := http.NewRequest("GET", url, nil)
		assert.NoError(t, err)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		// Prepare Test
		body, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err)
		err = json.Unmarshal(body, &res)
		assert.NoError(t, err)

		// Get Template Test
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.NotEmpty(t, res.Status)
		assert.Equal(t, "failed", res.Status)
		assert.NotEmpty(t, res.Message)
		assert.Equal(t, "Target col is not valid", res.Message)
	}
}
