package main

import (
	"fmt"
	"moneh/packages/tests"
	"testing"

	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestStats(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Moneh API Testing Suite - Stats")
}

var _ = Describe("Moneh API Testing - Stats", func() {
	const method = "get"
	const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjY3NjQ3OTg0NDIsImxldmVsIjoiYXBwbGljYXRpb24iLCJ1c2VybmFtZSI6ImZsYXplZnkifQ.-4KBslJIARdPoBtY_pyLvVGMhUs4zRz00keJrjMoiCM"
	const local_url = "http://127.0.0.1:1323"
	const ord = "desc"

	It(fmt.Sprintf("%s - Total Flow By Type", method), func() {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			Get(local_url + "/api/v1/stats/flowtype/" + ord)

		tests.ValidateResponse(resp, err)
	})

	It(fmt.Sprintf("%s - Total Flow By Category", method), func() {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			Get(local_url + "/api/v1/stats/flowcat/" + ord)

		tests.ValidateResponse(resp, err)
	})

	It(fmt.Sprintf("%s - Total Pocket By Type", method), func() {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			Get(local_url + "/api/v1/stats/pockettype/" + ord)

		tests.ValidateResponse(resp, err)
	})
})
