package assert_not_nil_int

import (
	"testing"

	"github.com/EugeneGpil/tester"
)

func Test_assert_not_nil_int(t *testing.T) {
	tester.SetTester(t)
	tester.AssertNotNil(1)
}
