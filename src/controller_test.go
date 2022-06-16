package main

import (
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Performing a MutationOperation", Label("Mutate"), func() {
	var router *gin.Engine
	BeforeEach(func() {
		router := initializeRouter()
	})

	When("we try to reverse a sequence", func() {
		BeforeEach(func() {
			Expect(router.Run()).To(Succeed())
		})

		Context("and the request is well formed", func() {
			It("the result is reversed", func() {

			})
		})

	})
})
