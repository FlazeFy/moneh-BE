package e2etest

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"moneh/models"
	"moneh/tests"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type ResponseGetAllDictionary struct {
	Data    []models.Dictionary `json:"data"`
	Message string              `json:"message"`
	Status  string              `json:"status"`
}

func TestSuccessGetAllDictionaryWithValidData(t *testing.T) {
	var res ResponseGetAllDictionary
	url := "http://127.0.0.1:9000/api/v1/dictionaries"
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
	assert.Equal(t, "Dictionary fetched", res.Message)
	assert.NotNil(t, res.Data)

	for _, dt := range res.Data {
		// Check Object
		assert.NotEqual(t, uuid.Nil, dt.ID)
		assert.NotEmpty(t, dt.DictionaryType)
		assert.NotEmpty(t, dt.DictionaryName)
		assert.NotEmpty(t, dt.CreatedAt)

		// Check Data Type
		assert.IsType(t, "", dt.DictionaryName)
		assert.IsType(t, "", dt.DictionaryType)
		assert.IsType(t, time.Time{}, dt.CreatedAt)
	}
}

func TestFailedPostCreateDictionaryWithShortCharDictionaryName(t *testing.T) {
	var res tests.ResponseFailedValidation
	url := "http://127.0.0.1:9000/api/v1/dictionaries"

	// Test Data
	token, _ := tests.TemplatePostBasicLogin(t, nil, nil, "admin")
	reqBody := map[string]interface{}{
		"dictionary_type": "A-Wallet",
		"dictionary_name": "NT",
	}
	jsonValue, _ := json.Marshal(reqBody)

	// Exec
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
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

	// Check Validation Message
	assert.Equal(t, "DictionaryName must be at least 3 characters long", res.Message[0].Error)
	assert.Equal(t, "DictionaryName", res.Message[0].Field)
}

func TestFailedPostCreateDictionaryWithEmptyDictionaryName(t *testing.T) {
	var res tests.ResponseFailedValidation
	url := "http://127.0.0.1:9000/api/v1/dictionaries"

	// Test Data
	token, _ := tests.TemplatePostBasicLogin(t, nil, nil, "admin")
	reqBody := map[string]interface{}{
		"dictionary_type": "A-Wallet",
		"dictionary_name": "",
	}
	jsonValue, _ := json.Marshal(reqBody)

	// Exec
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
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

	// Check Validation Message
	assert.Equal(t, "DictionaryName is required", res.Message[0].Error)
	assert.Equal(t, "DictionaryName", res.Message[0].Field)
}
