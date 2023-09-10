package assert_same_strings

import (
	"testing"

	"github.com/EugeneGpil/tester"
)

func Test_assert_same_strings(t *testing.T) {
	tester.SetTester(t)
	tester.AssertSame("test", "test")
}
