package katas_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestValidBraces(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ValidBraces Suite")
  }
