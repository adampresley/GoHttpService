package GoHttpService

import (
	"net/http"
)

/*
Convenience method sending out a BasicJsonResponse in JSON
format with a success of FALSE. This will set the HTTP status
code to 400. This is useful for telling a caller that the
request is malformed.
*/
func BadRequest(writer http.ResponseWriter, message string) {
	WriteJson(writer, NewBasicJsonResponse(false, message), 400)
}

/*
Convenience method sending out a BasicJsonResponse in JSON
format with a success of FALSE. This will set the HTTP status
code to 500. This is useful for indicating some type of
error ocurred during processing.
*/
func Error(writer http.ResponseWriter, message string) {
	WriteJson(writer, NewBasicJsonResponse(false, message), 500)
}

/*
Convenience method sending out a BasicJsonResponse in JSON
format with a success of FALSE. This will set the HTTP status
code to 404. This is useful in telling a caller that
the resources they are trying to get is not found.
*/
func NotFound(writer http.ResponseWriter, message string) {
	WriteJson(writer, NewBasicJsonResponse(false, message), 404)
}

/*
Convenience method sending out a BasicJsonResponse in JSON
format with a success of TRUE. This will set the HTTP status
code to 200.
*/
func Success(writer http.ResponseWriter, message string) {
	WriteJson(writer, NewBasicJsonResponse(true, message), 200)
}
