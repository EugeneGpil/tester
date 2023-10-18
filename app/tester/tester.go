package tester

import (
	"reflect"
	"testing"

	"github.com/EugeneGpil/tester/app/tester/methods"
)

var t *testing.T

func AssertLen(var1 interface{}, length int) {
	if len := reflect.ValueOf(var1).Len(); len != length {
		methods.Panic(`elements count = %q, want match for %q`, len, length)
	}
}

func AssertNil(var1 interface{}) {
	if var1 != nil {
		methods.Panic("variable is not nil")
	}
}

func AssertNotNil(var1 interface{}) {
	if var1 == nil {
		methods.Panic(`variable is nil`)
	}
}

func AssertSame(var1 interface{}, var2 interface{}) {
	methods.AssertSame(var1, var2, t)
}

func SetTester(tester *testing.T) {
	t = tester
}
