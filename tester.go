package tester

import (
	"testing"

	internalTester "github.com/EugeneGpil/tester/app/tester"
	"github.com/EugeneGpil/tester/app/types"
)

func SetTester(tester *testing.T) {
	internalTester.SetTester(tester)
}

func AssertSame(var1 interface{}, var2 interface{}) {
	internalTester.AssertSame(var1, var2)
}

func AssertNotNil(var1 interface{}) {
	internalTester.AssertNotNil(var1)
}

func AssertLen(var1 interface{}, length int) {
	internalTester.AssertLen(var1, length)
}

func GetTestResponseWriter() types.ResponseWriter {
	return internalTester.GetTestResponseWriter()
}
