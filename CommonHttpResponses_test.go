package GoHttpService

import (
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCommonHttpResponses(t *testing.T) {
	Convey("Testing common HTTP responses...", t, func() {

		Convey("BadRequest() will have a success of false and an HTTP status of 400", func() {
			expectedBody := "{\"success\":false,\"message\":\"Bad request\"}"
			expectedStatusCode := 400

			writer := httptest.NewRecorder()
			BadRequest(writer, "Bad request")

			actualBody := writer.Body.String()
			actualStatusCode := writer.Code

			So(actualBody, ShouldEqual, expectedBody)
			So(actualStatusCode, ShouldEqual, expectedStatusCode)
		})

		Convey("NotFound() will have a success of false and an HTTP status of 404", func() {
			expectedBody := "{\"success\":false,\"message\":\"Not found\"}"
			expectedStatusCode := 404

			writer := httptest.NewRecorder()
			NotFound(writer, "Not found")

			actualBody := writer.Body.String()
			actualStatusCode := writer.Code

			So(actualBody, ShouldEqual, expectedBody)
			So(actualStatusCode, ShouldEqual, expectedStatusCode)
		})

		Convey("Error() will have a success of false and an HTTP status of 500", func() {
			expectedBody := "{\"success\":false,\"message\":\"Error!\"}"
			expectedStatusCode := 500

			writer := httptest.NewRecorder()
			Error(writer, "Error!")

			actualBody := writer.Body.String()
			actualStatusCode := writer.Code

			So(actualBody, ShouldEqual, expectedBody)
			So(actualStatusCode, ShouldEqual, expectedStatusCode)
		})

		Convey("Success() will have a success of true and an HTTP status of 200", func() {
			expectedBody := "{\"success\":true,\"message\":\"All good!\"}"
			expectedStatusCode := 200

			writer := httptest.NewRecorder()
			Success(writer, "All good!")

			actualBody := writer.Body.String()
			actualStatusCode := writer.Code

			So(actualBody, ShouldEqual, expectedBody)
			So(actualStatusCode, ShouldEqual, expectedStatusCode)
		})

	})
}