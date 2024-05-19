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

	It(fmt.Sprintf("%s - Total Ammount Flow By Type", method), func() {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			Get(local_url + "/api/v1/stats/ammountflowtype/" + ord)

		tests.ValidateResponse(resp, err)
	})

	It(fmt.Sprintf("%s - Total Dct By Type", method), func() {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			Get(local_url + "/api/v1/stats/dcttype/" + ord)

		tests.ValidateResponse(resp, err)
	})

	It(fmt.Sprintf("%s - Total Wishlist Type", method), func() {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			Get(local_url + "/api/v1/stats/wishlisttype/" + ord)

		tests.ValidateResponse(resp, err)
	})

	It(fmt.Sprintf("%s - Total Wishlist Priority", method), func() {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			Get(local_url + "/api/v1/stats/wishlistpriority/" + ord)

		tests.ValidateResponse(resp, err)
	})

	It(fmt.Sprintf("%s - Total Wishlist Is Achieved", method), func() {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			Get(local_url + "/api/v1/stats/wishlistisachieved/" + ord)

		tests.ValidateResponse(resp, err)
	})

	It(fmt.Sprintf("%s - Total Dictionary To Module", method), func() {
		table := "flows"
		col := "flows_category"

		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			Get(local_url + "/api/v1/stats/dctmod/" + table + "/" + col)

		tests.ValidateResponse(resp, err)
	})
})
