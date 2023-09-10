package assert_len_of_slice_of_some_structs

import (
	"testing"

	"github.com/EugeneGpil/tester"
)

type someStruct struct {
	String string
}

func Test_assert_len_of_slice_of_some_structs(t *testing.T) {
	tester.SetTester(t)

	someStruct := []someStruct{
		{
			"Hello",
		},
		{
			"Hello",
		},
	}

	tester.AssertLen(someStruct, 2)
}
