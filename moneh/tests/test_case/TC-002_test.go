package testcase

import (
	"encoding/json"
	"fmt"
	"moneh/packages/tests"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test_TC002(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Moneh API Testing Suite - TC-002 Logout")
}

var _ = Describe("Moneh API Testing Suite - TC-002 Logout", func() {
	const method = "get"
	var tokenLogin string
	const local_url = "http://127.0.0.1:1323"

	It(fmt.Sprintf("%s - TC-002 Logout", method), func() {
		body := map[string]string{
			"username": "flazefy",
			"password": "nopass123",
		}

		client := resty.New()
		login, err := client.R().
			SetFormData(body).
			Post(local_url + "/api/v1/login")

		tests.ValidateResponse(login, err)

		// Get token value
		var result map[string]interface{}
		err = json.Unmarshal(login.Body(), &result)
		if err != nil {
			Fail(fmt.Sprintf("Failed to parse response body: %v", err))
		}

		if data, ok := result["data"].(map[string]interface{}); ok {
			if token, ok := data["token"].(string); ok {
				tokenLogin = token
				fmt.Println("Token:", tokenLogin)

				resp, err := client.R().
					SetHeader("Authorization", fmt.Sprintf("Bearer %s", tokenLogin)).
					Post(local_url + "/api/v1/logout")

				tests.ValidateResponse(resp, err)

				var result map[string]interface{}
				err = json.Unmarshal(resp.Body(), &result)
				if err != nil {
					Fail(fmt.Sprintf("Failed to parse response body: %v", err))
				}

				// Audit
				var respond tests.Record
				strRes, _ := json.Marshal(result)

				respond.Context = "Integration Test"
				respond.Title = "TC-002 Logout Test"
				respond.Request = tokenLogin
				respond.Response = string(strRes)
				respond.CreatedAt = time.Now()

				tests.WriteAudit(respond)
			} else {
				Fail("Token not found in response data")
			}
		} else {
			Fail("Data field not found in response")
		}
	})
})
