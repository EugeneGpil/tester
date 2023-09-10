package assert_same_bytes

import (
	"testing"

	"github.com/EugeneGpil/tester"
)

func Test_assert_same_bytes(t *testing.T) {
	tester.SetTester(t)
	tester.AssertSame([]byte("Hello"), []byte("Hello"))
}
