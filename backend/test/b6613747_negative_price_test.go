package test

import (
	"final-lab/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPriceIsValid(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("Price must be between 50 and 5000", func(t *testing.T) {
		b := entity.Books{
			Title: "Avatar",
			Price: 49,
			Code:  "BK123456",
		}
		ok, err := govalidator.ValidateStruct(b)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Price must be between 50 and 5000"))
	})
	t.Run("Price must be between 50 and 5000", func(t *testing.T) {
		b := entity.Books{
			Title: "Avatar",
			Price: 5001,
			Code:  "BK123456",
		}
		ok, err := govalidator.ValidateStruct(b)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Price must be between 50 and 5000"))
	})
}
