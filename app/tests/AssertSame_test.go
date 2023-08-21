package tests

import (
	"testing"

	"github.com/EugeneGpil/tester"
)

func Test_assert_same_strings(t *testing.T) {
	tester.SetTester(t)
	tester.AssertSame("test", "test")
}

func Test_assert_same_ints(t *testing.T) {
	tester.SetTester(t)
	tester.AssertSame(1, 1)
}

func Test_assert_same_bytes(t *testing.T) {
	tester.SetTester(t)
	tester.AssertSame([]byte("Hello"), []byte("Hello"))
}
