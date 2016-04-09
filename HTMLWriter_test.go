package GoHttpService

import (
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHTMLWriter(t *testing.T) {
	Convey("The HTML Writer...", t, func() {

		Convey("will write HTML", func() {
			expected := "<p>This is a test</p>"

			/*
			 * Mock the HTTP writer interface and call WriteHTML
			 */
			writer := httptest.NewRecorder()
			WriteText(writer, "<p>This is a test</p>", 200)

			actual := writer.Body.String()
			So(actual, ShouldEqual, expected)

			actualCode := writer.Code
			So(actualCode, ShouldEqual, 200)
		})
	})
}
