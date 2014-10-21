package GoHttpService

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBasicJsonResponse(t *testing.T) {
	Convey("A BasicJsonResponse...", t, func() {

		Convey("is returned with success of TRUE when calling NewBasicJsonResponse() passing success true", func() {
			expected := BasicJsonResponse{
				Success: true,
				Message: "Message",
			}

			actual := NewBasicJsonResponse(true, "Message")
			So(actual, ShouldResemble, expected)
		})

		Convey("is returned with success of FALSE when calling NewBasicJsonResponse() passing success false", func() {
			expected := BasicJsonResponse{
				Success: false,
				Message: "Message",
			}

			actual := NewBasicJsonResponse(false, "Message")
			So(actual, ShouldResemble, expected)
		})

	})
}