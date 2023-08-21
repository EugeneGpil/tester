package tests

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

func Test_assert_len_of_slice_of_ints(t *testing.T) {
	tester.SetTester(t)
	tester.AssertLen([]int{
		1,
		2,
	}, 2)
}

type someStruct struct {
	String string
}

func Test_assert_len_of_slice_of_some_structs(t *testing.T) {
	tester.SetTester(t)
	tester.AssertLen([]someStruct{
		{
			"Hello",
		},
		{
			"Hello",
		},
	}, 2)
}
