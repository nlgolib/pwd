package pwd_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPwd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pwd Suite")
}
