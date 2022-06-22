package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Count some sequences", func() {
	When("given a valid sequence", func() {
		It("returns the correct count for the sequence", func() {
			Expect(Count("ABCD")).To(Equal(4))
		})

		It("returns the correct count for an empty sequence", func() {
			Expect(Count("")).To(Equal(0))
		})
	})
})

// TODO: implement Reverse test
// TODO: implement Separate test
// TODO: implement EncodeRot13 test
// TODO: implement DecodeRot13 test
// TODO: implement EncodeBase64 test
// TODO: implement DecodeBase64 test
// TODO: implement EncodeShiftLeft test
// TODO: implement DecodeShiftLeft test
// TODO: implement EncodeShiftRight test
// TODO: implement DecodeShiftRight test
