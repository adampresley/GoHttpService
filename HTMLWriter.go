package GoHttpService

import (
	"fmt"
	"net/http"
)

func WriteHTML(writer http.ResponseWriter, text string, code int) {
	writer.Header().Set("Content-Type", "text/html; charset=UTF-8")
	writer.WriteHeader(code)
	fmt.Fprintf(writer, text)
}
