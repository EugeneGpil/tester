package assert_same_ints

import (
	"testing"

	"github.com/EugeneGpil/tester"
)

func Test_assert_same_ints(t *testing.T) {
	tester.SetTester(t)
	tester.AssertSame(1, 1)
}
