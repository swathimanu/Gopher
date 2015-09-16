package reverse_test

import (
	. "github.com/Gopher/reverse"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Reverse", func() {


	Context("String is NULL", func() {
		It("should return NULL value", func() {
			Expect(Reverse("")).To(Equal(""))
		})

	})
})
