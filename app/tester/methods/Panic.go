package methods

import "fmt"

func Panic(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	panic(message)
}
