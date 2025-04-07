package pwd_test

import (
	"strings"

	"github.com/nlgolib/pwd"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Password hashing", func() {
	It("should hash a password", Label("pass"), func() {
		hash, err := pwd.Hash("password")
		Expect(err).Should(BeNil())
		Expect(len(hash)).Should(BeNumerically(">", 0))
	})

	It("should verify a password", Label("pass"), func() {
		hash, err := pwd.Hash("password")
		Expect(err).Should(BeNil())
		Expect(pwd.Verify("password", hash)).Should(BeTrue())
	})

	It("should not verify a wrong password", Label("pass"), func() {
		hash, err := pwd.Hash("password")
		Expect(err).Should(BeNil())
		Expect(pwd.Verify("wrong_password", hash)).Should(BeFalse())
	})

	It("should handle empty password", Label("pass"), func() {
		hash, err := pwd.Hash("")
		Expect(err).Should(BeNil())
		Expect(len(hash)).Should(BeNumerically(">", 0))
		Expect(pwd.Verify("", hash)).Should(BeTrue())
	})

	It("should handle special characters in password", Label("pass"), func() {
		password := "p@ssw0rd!#$%^&*()"
		hash, err := pwd.Hash(password)
		Expect(err).Should(BeNil())
		Expect(pwd.Verify(password, hash)).Should(BeTrue())
	})

	It("should handle unicode characters", Label("pass"), func() {
		password := "パスワード123アБВ"
		hash, err := pwd.Hash(password)
		Expect(err).Should(BeNil())
		Expect(pwd.Verify(password, hash)).Should(BeTrue())
	})

	It("should return false for empty hash", Label("pass"), func() {
		Expect(pwd.Verify("password", "")).Should(BeFalse())
	})

	It("should handle long passwords within bcrypt limits", Label("pass"), func() {
		// bcrypt has a maximum length of 72 bytes
		password := strings.Repeat("a", 72)
		hash, err := pwd.Hash(password)
		Expect(err).Should(BeNil())
		Expect(pwd.Verify(password, hash)).Should(BeTrue())
	})

	It("should fail for passwords exceeding bcrypt limits", Label("pass"), func() {
		// passwords longer than 72 bytes should return an error
		password := strings.Repeat("a", 73)
		hash, err := pwd.Hash(password)
		Expect(err).Should(Not(BeNil()))
		Expect(hash).Should(BeEmpty())
	})
})
