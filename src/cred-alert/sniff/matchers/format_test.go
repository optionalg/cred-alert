package matchers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"cred-alert/scanners"
	"cred-alert/sniff/matchers"
)

var _ = Describe("Format", func() {
	var matcher matchers.Matcher

	BeforeEach(func() {
		matcher = matchers.Format("AKIA[A-Z0-9]{16}")
	})

	It("returns true when the line matches case-sensitively", func() {
		line := &scanners.Line{Content: []byte("aws_access_key_id: AKIAIOSFOEXAMPLETPWI")}
		Expect(matcher.Match(line)).To(BeTrue())
	})

	It("returns false when the line does not match case-sensitively", func() {
		line := &scanners.Line{Content: []byte("aws_access_key_id: akiaiosfoexampletpwi")}
		Expect(matcher.Match(line)).To(BeFalse())
	})

	It("returns false when the line does not match", func() {
		line := &scanners.Line{Content: []byte("does not match")}
		Expect(matcher.Match(line)).To(BeFalse())
	})
})
