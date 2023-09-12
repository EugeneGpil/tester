package tester

import (
	"testing"

	internalTester "github.com/EugeneGpil/tester/app/tester"
)

func AssertLen(var1 interface{}, length int) {
	internalTester.AssertLen(var1, length)
}

func AssertNil(var1 interface{}) {
	internalTester.AssertNil(var1)
}

func AssertNotNil(var1 interface{}) {
	internalTester.AssertNotNil(var1)
}

func AssertSame(var1 interface{}, var2 interface{}) {
	internalTester.AssertSame(var1, var2)
}

func SetTester(tester *testing.T) {
	internalTester.SetTester(tester)
}
