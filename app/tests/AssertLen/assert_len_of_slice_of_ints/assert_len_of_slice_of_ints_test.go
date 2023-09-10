package assert_len_of_slice_of_ints

import (
	"testing"

	"github.com/EugeneGpil/tester"
)

func Test_assert_len_of_slice_of_ints(t *testing.T) {
	tester.SetTester(t)
	tester.AssertLen([]int{
		1,
		2,
	}, 2)
}
