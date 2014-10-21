package GoHttpService

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func WriteJson(writer http.ResponseWriter, object interface{}, code int) {
	jsonBytes, _ := json.Marshal(object)
	content := strings.Replace(string(jsonBytes), "%", "%%", -1)

	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(code)
	fmt.Fprintf(writer, content)
}
