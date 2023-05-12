package internal_test

import (
	"context"
	. "from_scratch_wep_api/internal"
	"from_scratch_wep_api/internal/db"
	"from_scratch_wep_api/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("TestService", Ordered, func() {
	var svc *TestService

	BeforeEach(func() {
		db := db.NewLocalDB()
		svc = NewTestService(&db)
		Expect(svc).NotTo(BeNil())

	})

	Context("Testing func on Services", func() {
		JustBeforeEach(func() {
			tr := types.TestRequest{
				ID:       "1",
				Name:     "First",
				Quantity: 33,
			}
			_, _ = svc.CreateTest(context.Background(), tr)
		})

		When("Using local DB and simple", func() {
			It("can create a test from the service successfully", func(ctx SpecContext) {
				tr := types.TestRequest{
					ID:       "2",
					Name:     "Second",
					Quantity: 44,
				}
				t, err := svc.CreateTest(context.Background(), tr)
				Expect(err).NotTo(HaveOccurred())
				Expect(t).To(Equal(tr))
			})
			It("can fetch a test from the service successfully", func(ctx SpecContext) {
				t, err := svc.GetTests(ctx)
				Expect(err).NotTo(HaveOccurred())
				Expect(len(t)).ToNot(Equal(0))
				Expect(len(t)).To(Equal(1))
			})
		})
	})

})
