package test

import (
	"final-lab/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestCreateBookSuccess(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("Positive Case", func(t *testing.T) {
		b := entity.Books{
			Title: "Avatar",
			Price: 450,
			Code:  "BK123456",
		}
		ok, err := govalidator.ValidateStruct(b)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}
