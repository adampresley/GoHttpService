package GoHttpService

import (
	"fmt"
	"net/http"
)

func WriteText(writer http.ResponseWriter, text string, code int) {
	writer.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	writer.WriteHeader(code)
	fmt.Fprintf(writer, text)
}
