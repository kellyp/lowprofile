package lowprofile

import (
	. "github.com/kellyp/lowprofile/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/kellyp/lowprofile/Godeps/_workspace/src/github.com/onsi/gomega"
	"testing"
)

func TestLowprofile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lowprofile Suite")
}
