package model

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	validUser = User{
		Username:   "testuser",
		Firstname:  "Test",
		Lastname:   "User",
		Email:      "testuser@example.com",
		Status:     "A",
		Department: "IT",
	}
)

func TestModel(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Model Suite")
}

var _ = Describe("User", func() {
	Describe("IsValid", func() {
		Context("with valid user", func() {
			It("should return true", func() {
				Expect(validUser.IsValid()).To(BeTrue())
			})
		})

		Context("with invalid username", func() {
			It("should return true when username is valid", func() {
				user := validUser
				user.Username = "testuser1"
				Expect(user.IsValid()).To(BeTrue())
			})

			It("should return false when empty", func() {
				user := validUser
				user.Username = ""
				Expect(user.IsValid()).To(BeFalse())
			})

			It("should return false when too short", func() {
				user := validUser
				user.Username = "ab"
				Expect(user.IsValid()).To(BeFalse())
			})
		})

		Context("with invalid email", func() {
			It("should return false when empty", func() {
				user := validUser
				user.Email = ""
				Expect(user.IsValid()).To(BeFalse())
			})

			It("should return false when invalid format", func() {
				user := validUser
				user.Email = "invalid-email"
				Expect(user.IsValid()).To(BeFalse())
			})
		})

		Context("with invalid status", func() {
			It("should return false when empty", func() {
				user := validUser
				user.Status = ""
				Expect(user.IsValid()).To(BeFalse())
			})

			It("should return false when invalid value", func() {
				user := validUser
				user.Status = "X"
				Expect(user.IsValid()).To(BeFalse())
			})
		})

		Context("with invalid names", func() {
			It("should return false when firstname is empty", func() {
				user := validUser
				user.Firstname = ""
				Expect(user.IsValid()).To(BeFalse())
			})

			It("should return false when lastname is empty", func() {
				user := validUser
				user.Lastname = ""
				Expect(user.IsValid()).To(BeFalse())
			})
		})
	})
})
