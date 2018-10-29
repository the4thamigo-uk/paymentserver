package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/the4thamigo-uk/paymentserver/pkg/presentation"
	"net/http"
	"os/exec"
)

var _ = Describe("Payment Suite", func() {
	var (
		cmd         *exec.Cmd
		rootURL     string
		paymentsURL string
		err         error
		pp          *presentation.Payment
		method      string
	)

	BeforeEach(func() {
		cmd, rootURL, err = startServer()
		Expect(err).NotTo(HaveOccurred())
		paymentsURL = rootURL + "/payments"
	})

	BeforeEach(func() {
		pp = newDummyPayment()
	})

	AfterEach(func() {
		stopServer(cmd)
	})

	When("no payments exist on the server", func() {
		It("should return an empty list of payments", func() {
			_, body, err := httpDo("GET", paymentsURL, pp)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(body.Payments)).To(Equal(0))
		})

		It("should create payment X, and return a payment matching the one sent", func() {
			rsp, body, err := httpDo("POST", paymentsURL, pp)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.StatusCode).To(Equal(http.StatusOK))
			Expect(len(body.Payments)).To(Equal(1))
			xpp := body.Payments[0]

			By("creating a new ID at version 1 ")
			Expect(xpp.ID).ToNot(Equal(pp.ID))
			Expect(xpp.Version).To(Equal(1))

			By("returning a payment that matches the one sent")
			pp.Entity = xpp.Entity
			Expect(xpp).To(Equal(pp))
		})

		When("sending a payment X that is invalid", func() {
			JustBeforeEach(func() {
				pp.Attributes.ProcessingDate = ""
			})

			It("should fail to create payment X", func() {
				rsp, body, err := httpDo("POST", paymentsURL, pp)
				Expect(err).NotTo(HaveOccurred())
				Expect(rsp.StatusCode).To(Equal(http.StatusBadRequest))
				Expect(body.Error.Code).To(Equal(http.StatusBadRequest))
				Expect(body.Error.Message).ToNot(BeZero())
			})

			It("should fail to replace payment X", func() {
				rsp, body, err := httpDo("PUT", paymentsURL+"/"+pp.ID+"/0", pp)
				Expect(err).NotTo(HaveOccurred())
				Expect(rsp.StatusCode).To(Equal(http.StatusBadRequest))
				Expect(body.Error.Code).To(Equal(http.StatusBadRequest))
				Expect(body.Error.Message).ToNot(BeZero())
			})

			It("should fail to update payment X", func() {
				rsp, body, err := httpDo("PATCH", paymentsURL+"/"+pp.ID+"/0", pp)
				Expect(err).NotTo(HaveOccurred())
				Expect(rsp.StatusCode).To(Equal(http.StatusBadRequest))
				Expect(body.Error.Code).To(Equal(http.StatusBadRequest))
				Expect(body.Error.Message).ToNot(BeZero())
			})

			It("should fail to delete payment X", func() {
				rsp, body, err := httpDo("PATCH", paymentsURL+"/"+pp.ID+"/0", pp)
				Expect(err).NotTo(HaveOccurred())
				Expect(rsp.StatusCode).To(Equal(http.StatusBadRequest))
				Expect(body.Error.Code).To(Equal(http.StatusBadRequest))
				Expect(body.Error.Message).ToNot(BeZero())
			})
		})
	})

	When("only payment X exists on the server", func() {
		var (
			xpp *presentation.Payment
		)
		JustBeforeEach(func() {
			_, body, err := httpDo("POST", paymentsURL, pp)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(body.Payments)).To(Equal(1))
			xpp = body.Payments[0]
		})

		When("using GET", func() {
			It("should return item X", func() {
				_, body, err := httpDo("GET", paymentsURL+"/"+xpp.ID+"/0", pp)
				Expect(err).NotTo(HaveOccurred())
				Expect(len(body.Payments)).To(Equal(1))
				ypp := body.Payments[0]

				By("returning X in the list")
				Expect(ypp).To(Equal(xpp))
			})
			It("should return one item in the list of payments", func() {
				_, body, err := httpDo("GET", paymentsURL, pp)
				Expect(err).NotTo(HaveOccurred())
				Expect(len(body.Payments)).To(Equal(1))
				ypp := body.Payments[0]

				By("returning X in the list")
				Expect(ypp).To(Equal(xpp))
			})
		})

		When("using POST", func() {
			BeforeEach(func() {
				method = "POST"
			})
			It("should create payment Y", func() {
				rsp, body, err := httpDo(method, paymentsURL, pp)
				Expect(err).NotTo(HaveOccurred())
				Expect(rsp.StatusCode).To(Equal(http.StatusOK))
				Expect(len(body.Payments)).To(Equal(1))
				ypp := body.Payments[0]

				By("creating a new ID at version 1 ")
				Expect(ypp.ID).ToNot(Equal(xpp.ID))
				Expect(ypp.Version).To(Equal(1))

				By("returning a payment that matches the one sent")
				pp.Entity = ypp.Entity
				Expect(ypp).To(Equal(pp))
			})
		})

		When("using DELETE", func() {
			BeforeEach(func() {
				method = "DELETE"
			})
			It("should delete payment X using version 0", func() {
				rsp, body, err := httpDo(method, paymentsURL+"/"+xpp.ID+"/0", nil)
				Expect(err).NotTo(HaveOccurred())
				Expect(rsp.StatusCode).To(Equal(http.StatusOK))
				Expect(len(body.Payments)).To(Equal(1))
				ypp := body.Payments[0]

				By("returning a payment that matches X")
				Expect(ypp).To(Equal(xpp))
			})

			It("should delete payment X using version 1", func() {
				_, body, err := httpDo(method, paymentsURL+"/"+xpp.ID+"/1", nil)
				Expect(err).NotTo(HaveOccurred())
				Expect(len(body.Payments)).To(Equal(1))
				ypp := body.Payments[0]

				By("returning a payment that matches X")
				Expect(ypp).To(Equal(xpp))
			})

			It("should fail to delete payment X, using version 2", func() {
				rsp, _, err := httpDo(method, paymentsURL+"/"+xpp.ID+"/2", nil)
				Expect(err).NotTo(HaveOccurred())
				Expect(rsp.StatusCode).ToNot(Equal(http.StatusNotFound))
			})
		})

		When("using PUT", func() {
			BeforeEach(func() {
				method = "PUT"
			})
			It("should replace payment X at version 0", func() {
				pp.Attributes.Reference = "CHANGED"
				rsp, body, err := httpDo(method, paymentsURL+"/"+xpp.ID+"/0", pp)
				Expect(err).NotTo(HaveOccurred())
				Expect(rsp.StatusCode).To(Equal(http.StatusOK))
				Expect(len(body.Payments)).To(Equal(1))
				ypp := body.Payments[0]

				By("using the same ID as X and at version 2")
				Expect(ypp.ID).To(Equal(xpp.ID))
				Expect(ypp.Version).To(Equal(2))

				By("returning a payment that matches the one sent")
				pp.Entity = ypp.Entity
				Expect(ypp).To(Equal(pp))
			})

			It("should replace payment X at version 1", func() {
				pp.Attributes.Reference = "CHANGED"
				rsp, body, err := httpDo(method, paymentsURL+"/"+xpp.ID+"/1", pp)
				Expect(err).NotTo(HaveOccurred())
				Expect(rsp.StatusCode).To(Equal(http.StatusOK))
				Expect(len(body.Payments)).To(Equal(1))
				ypp := body.Payments[0]

				By("using the same ID as X and at version 2")
				Expect(ypp.ID).To(Equal(xpp.ID))
				Expect(ypp.Version).To(Equal(2))

				By("returning a payment that matches the one sent")
				pp.Entity = ypp.Entity
				Expect(ypp).To(Equal(pp))
			})

			It("should fail to update payment X, using version 2", func() {
				rsp, _, err := httpDo(method, paymentsURL+"/"+xpp.ID+"/2", nil)
				Expect(err).NotTo(HaveOccurred())
				Expect(rsp.StatusCode).ToNot(Equal(http.StatusNotFound))
			})
		})

		When("using PATCH", func() {
			BeforeEach(func() {
				method = "PATCH"
			})
			It("should update payment X at version 0", func() {
				pp.Attributes.Reference = "CHANGED"
				rsp, body, err := httpDo(method, paymentsURL+"/"+xpp.ID+"/0", pp)
				Expect(err).NotTo(HaveOccurred())
				Expect(rsp.StatusCode).To(Equal(http.StatusOK))
				Expect(len(body.Payments)).To(Equal(1))
				ypp := body.Payments[0]

				By("using the same ID as X and at version 2")
				Expect(ypp.ID).To(Equal(xpp.ID))
				Expect(ypp.Version).To(Equal(2))

				By("returning a payment that matches the one sent")
				pp.Entity = ypp.Entity
				Expect(ypp).To(Equal(pp))
			})

			It("should update payment X at version 1", func() {
				pp.Attributes.Reference = "CHANGED"
				rsp, body, err := httpDo(method, paymentsURL+"/"+xpp.ID+"/1", pp)
				Expect(err).NotTo(HaveOccurred())
				Expect(rsp.StatusCode).To(Equal(http.StatusOK))
				Expect(len(body.Payments)).To(Equal(1))
				ypp := body.Payments[0]

				By("using the same ID as X and at version 2")
				Expect(ypp.ID).To(Equal(xpp.ID))
				Expect(ypp.Version).To(Equal(2))

				By("returning a payment that matches the one sent")
				pp.Entity = ypp.Entity
				Expect(ypp).To(Equal(pp))
			})

			It("should fail to update payment X, using version 2", func() {
				rsp, _, err := httpDo(method, paymentsURL+"/"+xpp.ID+"/2", nil)
				Expect(err).NotTo(HaveOccurred())
				Expect(rsp.StatusCode).ToNot(Equal(http.StatusNotFound))
			})
		})
	})
})
