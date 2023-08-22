package tests

import (
	"testing"

	"github.com/EugeneGpil/tester"
)

func Test_assert_nil(t *testing.T) {
	tester.SetTester(t)
	tester.AssertNil(nil)
}
