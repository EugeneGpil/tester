package tests

import (
	"testing"

	"github.com/EugeneGpil/tester"
)

func Test_assert_not_nil_int(t *testing.T) {
	tester.SetTester(t)
	tester.AssertNotNil(1)
}

func Test_assert_not_nil_string(t *testing.T) {
	tester.SetTester(t)
	tester.AssertNotNil("nil")
}
