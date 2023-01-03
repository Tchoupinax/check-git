package utils_test

import (
	String "ckg/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("String", func() {
	Describe("AddSpaceToEnd", func() {
		It("should add no space at the end of the given string", func() {
			Expect(String.AddSpaceToEnd("hello", 0)).To(Equal("hello"))
		})

		It("should add five spaces at the end of the string", func() {
			Expect(String.AddSpaceToEnd("hello", 10)).To(Equal("hello     "))
		})
	})

	Describe("ContainsOneOfThese", func() {
		It("should answer false whent the string does not contains the substring", func() {
			Expect(String.ContainsOneOfThese("", []string{"s"})).To(Equal(false))
		})

		It("should answer true whent the string contains the substring", func() {
			Expect(String.ContainsOneOfThese("hello", []string{"lo"})).To(Equal(true))
		})

		It("should answer true whent the string is equal to one of the given strings", func() {
			Expect(String.ContainsOneOfThese("aehfeihfziuhfe", []string{"lo", "aehfeihfziuhfe"})).To(Equal(true))
		})
	})
})
