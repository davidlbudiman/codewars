package katas_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestKatas(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Katas Suite")
}
