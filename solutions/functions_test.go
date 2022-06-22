// package main

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

var _ = Describe("", func() {
	When("given a valid sequence", func() {
		It("returns the reversed sequence", func() {
			Expect(Reverse("ABCD")).To(Equal("DCBA"))
		})

		It("returns the correct reversed sequence", func() {
			Expect(Reverse("ZY EBS-73")).To(Equal("37-SBE YZ"))
		})
	})
})

var _ = Describe("Separate some sequences", func() {
	When("given a valid sequence", func() {
		It("returns the separated sequence", func() {
			Expect(Separate("ABCD")).To(Equal("A,B,C,D"))
		})

		It("returns the correct separated sequence", func() {
			Expect(Separate("wow")).To(Equal("w,o,w"))
		})
	})
})

// TODO: implement EncodeRot13 test
// TODO: implement DecodeRot13 test
// TODO: implement EncodeBase64 test
// TODO: implement DecodeBase64 test
// TODO: implement EncodeShiftLeft test
// TODO: implement DecodeShiftLeft test
// TODO: implement EncodeShiftRight test
// TODO: implement DecodeShiftRight test
