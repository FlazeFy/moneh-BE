package main

import (
	"fmt"
	"moneh/packages/tests"
	"testing"

	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDct(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Moneh API Testing Suite - Dictionary")
}

var _ = Describe("Moneh API Testing - Dictionary", func() {
	const method = "get"
	const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjY3NjQ3OTg0NDIsImxldmVsIjoiYXBwbGljYXRpb24iLCJ1c2VybmFtZSI6ImZsYXplZnkifQ.-4KBslJIARdPoBtY_pyLvVGMhUs4zRz00keJrjMoiCM"
	const local_url = "http://127.0.0.1:1323"
	const types = "inventory_unit"
	const page = "1"

	It(fmt.Sprintf("%s - Dictionary By Type", method), func() {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			Get(local_url + "/api/v1/dct/" + types + "?page=" + page)

		tests.ValidateResponse(resp, err)
	})
})
