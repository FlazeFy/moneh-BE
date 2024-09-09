package main

import (
	"fmt"
	"moneh/packages/tests"
	"testing"

	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFlows(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Moneh API Testing Suite - Flow")
}

var _ = Describe("Moneh API Testing - Flow", func() {
	const method = "get"
	const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjY3NzQ0MjMyNzIsImxldmVsIjoiYXBwbGljYXRpb24iLCJ1c2VybmFtZSI6ImZsYXplZnkifQ.fveMDGxf1oEVHSJ7H94NHp1-Fk9EGgAZtvQV_N0VYd8"
	const local_url = "http://127.0.0.1:1323"
	const ord = "desc"
	const mon = "7"
	const year = "2023"
	const typeVal = "spending"
	const view = "date"

	It(fmt.Sprintf("%s - All Flow", method), func() {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			Get(local_url + fmt.Sprintf("/api/v1/flows/%s", ord))

		tests.ValidateResponse(resp, err)
	})

	It(fmt.Sprintf("%s - All Flow (Export)", method), func() {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			Get(local_url + "/api/v2/flows")

		tests.ValidateResponse(resp, err)
	})

	It(fmt.Sprintf("%s - Monthly Flow Item", method), func() {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			Get(local_url + fmt.Sprintf("/api/v1/flows/month_item/%s/%s/%s", mon, year, typeVal))

		tests.ValidateResponse(resp, err)
	})

	It(fmt.Sprintf("%s - Monthly Flow Total", method), func() {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			Get(local_url + fmt.Sprintf("/api/v1/flows/month_total/%s/%s/%s", mon, year, typeVal))

		tests.ValidateResponse(resp, err)
	})

	It(fmt.Sprintf("%s - Summary By Type", method), func() {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			Get(local_url + fmt.Sprintf("/api/v1/flows/summary/%s", typeVal))

		tests.ValidateResponse(resp, err)
	})

	It(fmt.Sprintf("%s - Total Item Amount Per Date By Type", method), func() {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			Get(local_url + fmt.Sprintf("/api/v1/flows/dateammount/%s/%s", typeVal, view))

		tests.ValidateResponse(resp, err)
	})
})
