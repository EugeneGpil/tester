package assert_len_of_slice_of_strings

import (
	"testing"

	"github.com/EugeneGpil/tester"
)

func Test_assert_len_of_slice_of_strings(t *testing.T) {
	tester.SetTester(t)
	tester.AssertLen([]string{
		"hello1",
		"hello2",
	}, 2)
}
