package e2etest

import (
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

type ResponseGetAllHistory struct {
	Data    []models.History `json:"data"`
	Message string           `json:"message"`
	Status  string           `json:"status"`
}

func TestSuccessGetAllHistoryWithValidData(t *testing.T) {
	var res ResponseGetAllHistory
	url := "http://127.0.0.1:9000/api/v1/histories/my"
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
	assert.Equal(t, "History fetched", res.Message)
	assert.NotNil(t, res.Data)

	for _, dt := range res.Data {
		// Check Object
		assert.NotEqual(t, uuid.Nil, dt.ID)
		assert.NotEmpty(t, dt.HistoryType)
		assert.NotEmpty(t, dt.HistoryContext)
		assert.NotEmpty(t, dt.CreatedAt)

		// Check Data Type
		assert.IsType(t, "", dt.HistoryContext)
		assert.IsType(t, "", dt.HistoryType)
		assert.IsType(t, time.Time{}, dt.CreatedAt)
	}
}

func TestFailedDeleteHistoryByIdWithInvalidId(t *testing.T) {
	var res tests.ResponseSimple
	id := "79ff75f8-4b24-4fd8-9811-11075d0ccc81"
	url := "http://127.0.0.1:9000/api/v1/histories/destroy/" + id

	// Test Data
	token, _ := tests.TemplatePostBasicLogin(t, nil, nil, "user")

	// Exec
	req, err := http.NewRequest("DELETE", url, nil)
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
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.NotEmpty(t, res.Status)
	assert.Equal(t, "failed", res.Status)

	// Check Validation Message
	assert.Equal(t, "History not found", res.Message)
}

func TestSuccessDeleteHistoryByIdWithValidId(t *testing.T) {
	var res tests.ResponseSimple
	id := "3f49d833-cfce-499e-a4aa-68ddb705a367"
	url := "http://127.0.0.1:9000/api/v1/histories/destroy/" + id

	// Test Data
	token, _ := tests.TemplatePostBasicLogin(t, nil, nil, "user")

	// Exec
	req, err := http.NewRequest("DELETE", url, nil)
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

	// Check Validation Message
	assert.Equal(t, "History permanentally deleted", res.Message)
}
