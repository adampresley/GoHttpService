package GoHttpService

import (
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTextWriter(t *testing.T) {
	Convey("The Text Writer...", t, func() {

		Convey("will write text", func() {
			expected := "This is a test"

			/*
			 * Mock the HTTP writer interface and call WriteJson
			 */
			writer := httptest.NewRecorder()
			WriteText(writer, "This is a test", 200)

			actual := writer.Body.String()
			So(actual, ShouldEqual, expected)

			actualCode := writer.Code
			So(actualCode, ShouldEqual, 200)
		})
	})
}
