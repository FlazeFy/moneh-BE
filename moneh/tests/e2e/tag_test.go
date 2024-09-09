package main

import (
	"fmt"
	"moneh/packages/tests"
	"testing"

	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTags(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Moneh API Testing Suite - Tag")
}

var _ = Describe("Moneh API Testing - Tag", func() {
	const method = "get"
	const local_url = "http://127.0.0.1:1323"
	const ord = "desc"
	const ord_target = "created_at"

	It(fmt.Sprintf("%s - All Tag", method), func() {
		client := resty.New()
		resp, err := client.R().
			Get(local_url + fmt.Sprintf("/api/v1/tag/%s", ord))

		tests.ValidateResponse(resp, err)
	})
	It(fmt.Sprintf("%s - All Tag (Firebase)", method), func() {
		client := resty.New()
		resp, err := client.R().
			Get(local_url + fmt.Sprintf("/api/v2/tag/%s", ord))

		tests.ValidateResponse(resp, err)
	})
})
