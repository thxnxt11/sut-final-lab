package test

import (
	"final-lab/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestCodeIsValid(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("Code must start with BK followed by 6 digits (0-9)", func(t *testing.T) {
		b := entity.Books{
			Title: "Avatar",
			Price: 450,
			Code:  "AK123456",
		}
		ok, err := govalidator.ValidateStruct(b)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Code must start with BK followed by 6 digits (0-9)"))
	})
	t.Run("Code must start with BK followed by 6 digits (0-9)", func(t *testing.T) {
		b := entity.Books{
			Title: "Avatar",
			Price: 500,
			Code:  "BK123456789",
		}
		ok, err := govalidator.ValidateStruct(b)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Code must start with BK followed by 6 digits (0-9)"))
	})
	t.Run("Code must start with BK followed by 6 digits (0-9)", func(t *testing.T) {
		b := entity.Books{
			Title: "Avatar",
			Price: 500,
			Code:  "123456BK",
		}
		ok, err := govalidator.ValidateStruct(b)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Code must start with BK followed by 6 digits (0-9)"))
	})
	t.Run("Code must start with BK followed by 6 digits (0-9)", func(t *testing.T) {
		b := entity.Books{
			Title: "Avatar",
			Price: 500,
			Code:  "BK",
		}
		ok, err := govalidator.ValidateStruct(b)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Code must start with BK followed by 6 digits (0-9)"))
	})
	t.Run("Code must start with BK followed by 6 digits (0-9)", func(t *testing.T) {
		b := entity.Books{
			Title: "Avatar",
			Price: 500,
			Code:  "BK_123456",
		}
		ok, err := govalidator.ValidateStruct(b)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Code must start with BK followed by 6 digits (0-9)"))
	})
}
