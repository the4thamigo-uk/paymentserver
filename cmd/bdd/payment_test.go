package tests

import (
	. "github.com/onsi/ginkgo"
	//	. "github.com/onsi/gomega"
)

var _ = Describe("Payment Suite", func() {
	When("no payments exist on the server", func() {
		It("should return an empty list of payments", func() {
		})

		It("should create payment X, and return a payment matching the one sent", func() {
		})

		When("sending a payment X that is invalid", func() {
			It("should fail to create payment X", func() {
			})

			It("should fail to replace payment X", func() {
			})

			It("should fail to update payment X", func() {
			})

			It("should fail to delete payment X", func() {
			})
		})
	})

	When("only payment X exists on the server", func() {
		When("using GET", func() {
			It("should return item X", func() {
			})
			It("should return one item in the list of payments", func() {
			})
		})

		When("using POST", func() {
			It("should create payment Y", func() {
			})
		})

		When("using DELETE", func() {
			It("should delete payment X using version 0", func() {
			})

			It("should delete payment X using version 1", func() {
			})

			It("should fail to delete payment X, using version 2", func() {
			})
		})

		When("using PUT", func() {
			It("should replace payment X at version 0", func() {
			})

			It("should replace payment X at version 1", func() {
			})

			It("should fail to update payment X, using version 2", func() {
			})
		})

		When("using PATCH", func() {
			It("should update payment X at version 0", func() {
			})

			It("should update payment X at version 1", func() {
			})

			It("should fail to update payment X, using version 2", func() {
			})
		})
	})
})
