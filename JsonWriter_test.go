package GoHttpService

import (
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type FakeSomethingJsonWriter struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func TestJsonWriter(t *testing.T) {
	Convey("The JSON Writer...", t, func() {

		Convey("will serialize a map to JSON", func() {
			object := make(map[string]string)
			object["id"] = "1"
			object["name"] = "Adam"

			expected := "{\"id\":\"1\",\"name\":\"Adam\"}"

			/*
			 * Mock the HTTP writer interface and call WriteJson
			 */
			writer := httptest.NewRecorder()
			WriteJson(writer, object, 200)

			actual := writer.Body.String()
			So(actual, ShouldEqual, expected)
		})

		Convey("will serialize a struct to JSON", func() {
			object := FakeSomethingJsonWriter{Id: 1, Name: "Adam"}
			expected := "{\"id\":1,\"name\":\"Adam\"}"

			/*
			 * Mock the HTTP writer interface and call WriteJson
			 */
			writer := httptest.NewRecorder()
			WriteJson(writer, object, 200)

			actual := writer.Body.String()
			So(actual, ShouldEqual, expected)
		})

		Convey("will write a status code of 200", func() {
			object := FakeSomethingJsonWriter{Id: 1, Name: "Adam"}
			expected := 200

			/*
			 * Mock the HTTP writer interface and call WriteJson
			 */
			writer := httptest.NewRecorder()
			WriteJson(writer, object, 200)

			actual := writer.Code
			So(actual, ShouldEqual, expected)
		})

	})
}