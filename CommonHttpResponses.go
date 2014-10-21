package GoHttpService

import (
	"net/http"
)

func BadRequest(writer http.ResponseWriter, message string) {
	WriteJson(writer, NewBasicJsonResponse(false, message), 400)
}

func Error(writer http.ResponseWriter, message string) {
	WriteJson(writer, NewBasicJsonResponse(false, message), 500)
}

func NotFound(writer http.ResponseWriter, message string) {
	WriteJson(writer, NewBasicJsonResponse(false, message), 404)
}

func Success(writer http.ResponseWriter, message string) {
	WriteJson(writer, NewBasicJsonResponse(true, message), 200)
}
