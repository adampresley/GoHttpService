package GoHttpService

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type FakeSomethingBodyParser struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func TestBodyParser(t *testing.T) {
	Convey("The body parser...", t, func() {

		Convey("will return a map with valid JSON in the request body", func() {
			actual := make(map[string]string)
			expected := make(map[string]string)

			expected["id"] = "1"
			expected["name"] = "Adam"

			/*
			 * Have our fake handler do the parse call and assign
			 * to the receiver "actual"
			 */
			handler := func(w http.ResponseWriter, r *http.Request) {
				ParseJsonBody(r, &actual)
			}

			/*
			 * Mock the HTTP writer interface, setup a POST body, and parse
			 */
			body := "{\"id\":\"1\",\"name\":\"Adam\"}"
			bodyReader := strings.NewReader(body)

			request, _ := http.NewRequest("POST", "http://localhost/testmap", bodyReader)
			request.Header.Set("Content-Type", "application/json")

			writer := httptest.NewRecorder()
			handler(writer, request)

			So(actual, ShouldResemble, expected)
		})

		Convey("will return a struct with valid JSON in the request body", func() {
			actual := FakeSomethingBodyParser{}
			expected := FakeSomethingBodyParser{Id: 2, Name: "Adam"}

			/*
			 * Have our fake handler do the parse call and assign
			 * to the receiver "actual"
			 */
			handler := func(w http.ResponseWriter, r *http.Request) {
				ParseJsonBody(r, &actual)
			}

			/*
			 * Mock the HTTP writer interface, setup a POST body, and parse
			 */
			body := "{\"id\":2,\"name\":\"Adam\"}"
			bodyReader := strings.NewReader(body)

			request, _ := http.NewRequest("POST", "http://localhost/teststruct", bodyReader)
			request.Header.Set("Content-Type", "application/json")

			writer := httptest.NewRecorder()
			handler(writer, request)

			So(actual, ShouldResemble, expected)
		})

		Convey("will return an error with invalid JSON in the request body", func() {
			var actual error
			receiver := make(map[string]string)
			expected := &json.SyntaxError{}

			/*
			 * Have our fake handler do the parse call and assign
			 * to the receiver "actual"
			 */
			handler := func(w http.ResponseWriter, r *http.Request) {
				actual = ParseJsonBody(r, &receiver)
			}

			/*
			 * Mock the HTTP writer interface, setup a POST body, and parse
			 */
			body := "bob"
			bodyReader := strings.NewReader(body)

			request, _ := http.NewRequest("POST", "http://localhost/testerror", bodyReader)
			request.Header.Set("Content-Type", "application/json")

			writer := httptest.NewRecorder()
			handler(writer, request)

			So(actual, ShouldHaveSameTypeAs, expected)
		})

	})
}