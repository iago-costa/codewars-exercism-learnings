package reverse_words_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestReverseWords(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Reverse Words Suite")
}