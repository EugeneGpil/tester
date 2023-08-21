package tests

import (
	"net/http"
	"testing"

	"github.com/EugeneGpil/tester"
)

func Test_should_return_implementation_of_net_http_responseWriter(t *testing.T) {
	getResponseWriter()
}

func Test_should_write_message(t *testing.T) {
	tester.SetTester(t)
	writer := tester.GetTestResponseWriter()

	writer.Write([]byte("Hello"))

	messages := writer.GetMessages()

	tester.AssertLen(messages, 1)

	tester.AssertSame(messages[0], []byte("Hello"))
}

func getResponseWriter() http.ResponseWriter {
	return tester.GetTestResponseWriter()
}
