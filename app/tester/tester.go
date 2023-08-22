package tester

import (
	"reflect"
	"testing"

	"github.com/EugeneGpil/tester/app/tester/methods"
	"github.com/EugeneGpil/tester/app/types"
)

var t *testing.T

func AssertLen(var1 interface{}, length int) {
	if len := reflect.ValueOf(var1).Len(); len != length {
		t.Fatalf(`routes count = %q, want match for %q`, len, length)
	}
}

func AssertNil(var1 interface{}) {
	if var1 != nil {
		t.Fatalf("variable is not nil")
	}
}

func AssertNotNil(var1 interface{}) {
	if var1 == nil {
		t.Fatalf(`variable is nil`)
	}
}

func AssertSame(var1 interface{}, var2 interface{}) {
	methods.AssertSame(var1, var2, t)
}

func GetTestResponseWriter() types.ResponseWriter {
	return types.NewResponseWriter()
}

func SetTester(tester *testing.T) {
	t = tester
}
