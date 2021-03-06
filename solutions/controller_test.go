// package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"io/ioutil"
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

	When("we submit a malformed request", func() {
		Context("the request body is empty", func() {
			It("returns a 400 response", func() {
				request, _ = http.NewRequest(
					http.MethodPost,
					"/mutate",
					nil,
				)
				router.ServeHTTP(writer, request)

				Expect(writer).To(HaveHTTPStatus(http.StatusBadRequest))
				Expect(writer).To(HaveHTTPBody(
					MatchJSON(`{
						"message": "Invalid data contained in request body."
					}`),
				))
			})
		})

		Context("the request body is missing a required property", func() {
			It("returns a 400 response", func() {
				requestBody := []byte(`{"sequence": "wow,such,a,sequence"}`)

				request, _ = http.NewRequest(
					http.MethodPost,
					"/mutate",
					bytes.NewReader(requestBody),
				)
				router.ServeHTTP(writer, request)

				Expect(writer).To(HaveHTTPStatus(http.StatusBadRequest))
				Expect(writer).To(HaveHTTPBody(
					MatchJSON(`{
						"message": "Invalid data contained in request body."
					}`),
				))
			})
		})

		Context("the request body contains an invalid type", func() {
			It("returns a 400 response", func() {
				requestBody := []byte(`{
					"sequence": "wow,such,a,sequence",
					"operation": 5
				}`)

				request, _ = http.NewRequest(
					http.MethodPost,
					"/mutate",
					bytes.NewReader(requestBody),
				)
				router.ServeHTTP(writer, request)

				Expect(writer).To(HaveHTTPStatus(http.StatusBadRequest))
				Expect(writer).To(HaveHTTPBody(
					MatchJSON(`{
						"message": "Invalid data contained in request body."
					}`),
				))
			})
		})
	})

	When("we submit a well formed request", func() {
		Context("we try to count the number of characters in a sequence", func() {
			It("the proper count is returned", func() {
				requestBody := []byte(`{
					"sequence": "ABCD",
					"operation": "count"
				}`)

				request, _ = http.NewRequest(
					http.MethodPost,
					"/mutate",
					bytes.NewReader(requestBody),
				)
				router.ServeHTTP(writer, request)

				Expect(writer).To(HaveHTTPStatus(http.StatusOK))
				responseBody, readErr := ioutil.ReadAll(writer.Body)
				Expect(readErr).To(BeNil())
				var response MutationResult
				err := json.Unmarshal(responseBody, &response)
				Expect(err).To(BeNil())
				Expect(response.Result).To(Equal("4"))
			})
		})

		Context("we try to reverse a sequence", func() {
			It("returns a result with the reversed sequence", func() {
				requestBody := []byte(`{
					"sequence": "ABCD",
					"operation": "reverse"
				}`)

				request, _ = http.NewRequest(
					http.MethodPost,
					"/mutate",
					bytes.NewReader(requestBody),
				)
				router.ServeHTTP(writer, request)

				Expect(writer).To(HaveHTTPStatus(http.StatusInternalServerError))
				responseBody, readErr := ioutil.ReadAll(writer.Body)
				Expect(readErr).To(BeNil())
				var response MutationResult
				err := json.Unmarshal(responseBody, &response)
				Expect(err).To(BeNil())
				Expect(response.Result).To(Equal(""))
			})
		})

		Context("we try to separate a sequence", func() {
			It("the result is separated", func() {
				requestBody := []byte(`{
					"sequence": "ABCD",
					"operation": "separate"
				}`)

				request, _ = http.NewRequest(
					http.MethodPost,
					"/mutate",
					bytes.NewReader(requestBody),
				)
				router.ServeHTTP(writer, request)

				Expect(writer).To(HaveHTTPStatus(http.StatusInternalServerError))
				responseBody, readErr := ioutil.ReadAll(writer.Body)
				Expect(readErr).To(BeNil())
				var response MutationResult
				err := json.Unmarshal(responseBody, &response)
				Expect(err).To(BeNil())
				Expect(response.Result).To(Equal("A,B,C,D"))
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

	When("we submit a malformed request", func() {
		Context("the request body is empty", func() {
			It("returns a 400 response", func() {
				request, _ = http.NewRequest(
					http.MethodPost,
					"/find",
					nil,
				)
				router.ServeHTTP(writer, request)

				Expect(writer).To(HaveHTTPStatus(http.StatusBadRequest))
				Expect(writer).To(HaveHTTPBody(
					MatchJSON(`{
						"message": "Invalid data contained in request body."
					}`),
				))
			})
		})

		Context("the request body is missing the id parameter", func() {
			It("returns a 400 response", func() {
				requestBody := []byte(`{"another_field": "that isn't id"}`)

				request, _ = http.NewRequest(
					http.MethodPost,
					"/find",
					bytes.NewReader(requestBody),
				)
				router.ServeHTTP(writer, request)

				Expect(writer).To(HaveHTTPStatus(http.StatusBadRequest))
				Expect(writer).To(HaveHTTPBody(
					MatchJSON(`{
						"message": "Invalid data contained in request body."
					}`),
				))
			})
		})
	})

	When("we submit a well formed request", func() {
		Context("the result is not in the cache", func() {
			It("returns a 404 response", func() {
				requestBody := []byte(`{"id": "fbe89b26-beb1-4595-a928-f2998a768752"}`)

				request, _ = http.NewRequest(
					http.MethodPost,
					"/find",
					bytes.NewReader(requestBody),
				)
				router.ServeHTTP(writer, request)

				Expect(writer).To(HaveHTTPStatus(http.StatusNotFound))
				Expect(writer).To(HaveHTTPBody(
					MatchJSON(`{
						"message": "The requested MutationResult could not be found."
					}`),
				))
			})
		})

		Context("the result is in the cache", func() {
			It("the result is returned", func() {
				resultId := uuid.MustParse("fbe89b26-beb1-4595-a928-f2998a768752")
				// populate the cache
				addToCache(MutationResult{
					Id:     resultId,
					Result: "A,B,C,D",
				})

				requestBody := []byte(`{"id": "fbe89b26-beb1-4595-a928-f2998a768752"}`)

				request, _ = http.NewRequest(
					http.MethodPost,
					"/find",
					bytes.NewReader(requestBody),
				)
				router.ServeHTTP(writer, request)

				Expect(writer).To(HaveHTTPStatus(http.StatusOK))
				Expect(writer).To(HaveHTTPBody(
					MatchJSON(`{
						"id": "fbe89b26-beb1-4595-a928-f2998a768752",
						"result": "A,B,C,D"
					}`),
				))
			})
		})
	})
})
