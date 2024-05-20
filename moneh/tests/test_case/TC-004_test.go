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

func Test_TC004(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Moneh API Testing Suite - TC-004 Show all pocket")
}

var _ = Describe("Moneh API Testing Suite - TC-004 Show all pocket", func() {
	const method = "get"
	var tokenLogin string
	const local_url = "http://127.0.0.1:1323"
	page := "1"
	ord := "desc"

	It(fmt.Sprintf("%s - TC-004 Show all pocket", method), func() {
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
					Get(local_url + "/api/v1/pockets/headers/" + ord + "?page=" + page)

				tests.ValidateResponse(resp, err)

				// validate item column
				var resultItem map[string]interface{}
				err = json.Unmarshal(resp.Body(), &resultItem)
				if err != nil {
					Fail(fmt.Sprintf("Failed to parse response body: %v", err))
				}

				if data, ok := resultItem["data"].(map[string]interface{}); ok {
					if dataArray, ok := data["data"].([]interface{}); ok {
						for _, item := range dataArray {
							if flow, ok := item.(map[string]interface{}); ok {
								if _, ok := flow["id"].(string); !ok {
									Fail("id field is missing or not a string")
								}
								if _, ok := flow["pockets_name"].(string); !ok {
									Fail("pockets_name field is missing or not a string")
								}
								if _, ok := flow["pockets_type"].(string); !ok {
									Fail("pockets_type field is missing or not a string")
								}
								if _, ok := flow["pockets_desc"].(string); !ok {
									Fail("pockets_desc field is missing or not a string")
								}
								if pocketsLimit, ok := flow["pockets_limit"].(float64); ok {
									if pocketsLimit != float64(int64(pocketsLimit)) {
										Fail("pockets_limit field is not a valid integer")
									}
								} else {
									Fail("pockets_limit field is missing or not an integer")
								}
							} else {
								Fail("Item in data array is not a valid object")
							}
						}
					} else {
						Fail("data field is not an array")
					}
				} else {
					Fail("data field not found in response")
				}

				var result map[string]interface{}
				err = json.Unmarshal(resp.Body(), &result)
				if err != nil {
					Fail(fmt.Sprintf("Failed to parse response body: %v", err))
				}

				// Audit
				var respond tests.Record
				strRes, _ := json.Marshal(result["data"])

				respond.Context = "Integration Test"
				respond.Title = "TC-004 Show all pocket Test"
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
