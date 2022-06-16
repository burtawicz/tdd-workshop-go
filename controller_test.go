package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Performing a MutationOperation", Label("Mutate"), func() {
	var router *gin.Engine
	var writer *httptest.ResponseRecorder
	var request *http.Request

	BeforeEach(func() {
		router = initializeRouter()
		writer = httptest.NewRecorder()
	})

	When("we try to reverse a sequence", func() {
		Context("and the request is well formed", func() {
			It("the result is reversed", func() {
				requestBody := []byte(`{
					"sequence": "ABCD",
					"operation": "reverse",
				}`)

				request, _ = http.NewRequest(
					http.MethodPost,
					"/mutate",
					bytes.NewReader(requestBody),
				)
				router.ServeHTTP(writer, request)

				Expect(writer).To(HaveHTTPStatus(http.StatusOK))
				Expect(writer).To(HaveHTTPBody(
					MatchJSON(`{"message": "We're currently working on our implementation. Stay tuned."}`),
				))
			})
		})
	})
})

var _ = Describe("Looking for a MutationResult", Label("Find"), func() {
	var router *gin.Engine
	var writer *httptest.ResponseRecorder
	var request *http.Request

	BeforeEach(func() {
		router = initializeRouter()
		writer = httptest.NewRecorder()
	})

	When("we submit a valid id", func() {
		Context("and the request is well formed", func() {
			It("the result is returned", func() {
				requestBody := []byte(`{
					"id": "fbe89b26-beb1-4595-a928-f2998a768752",
				}`)

				request, _ = http.NewRequest(
					http.MethodPost,
					"/find",
					bytes.NewReader(requestBody),
				)
				router.ServeHTTP(writer, request)

				Expect(writer).To(HaveHTTPStatus(http.StatusOK))
				Expect(writer).To(HaveHTTPBody(
					MatchJSON(`{"message": "We're currently working on our implementation. Stay tuned."}`),
				))
			})
		})
	})
})
