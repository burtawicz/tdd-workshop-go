package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Count some strings", func() {
	When("given a valid string", func() {
		It("returns the correct count for the string", func() {
			Expect(Count("ABCD")).To(Equal(4))
			Expect(Count("")).To(Equal(0))
		})
	})
})

// TODO: implement TestReverse
// TODO: implement TestSeparate
// TODO: implement TestEncodeRot13
// TODO: implement TestDecodeRot13
// TODO: implement TestEncodeBase64
// TODO: implement TestDecodeBase64
// TODO: implement TestEncodeShiftLeft
// TODO: implement TestDecodeShiftLeft
// TODO: implement TestEncodeShiftRight
// TODO: implement TestDecodeShiftRight
