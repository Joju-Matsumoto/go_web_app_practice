package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestWebService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WebService Suite")
}
