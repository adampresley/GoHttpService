package GoHttpService

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/*
This function is useful when a POST or PUT request has a
Content-Type of application/json and the body has JSON
data. Pass the address of a receiver variable and this
function will fill in the values from the request JSON
into any matching struct fields, or will hydrate them
into a map if possible.
*/
func ParseJsonBody(request *http.Request, receiver interface{}) error {
	var err error

	body, _ := ioutil.ReadAll(request.Body)
	err = json.Unmarshal(body, receiver)
	return err
}

/*
Returns true/false if the request contains a Content-Type
header, and if that header is set to "application/json".
*/
func RequestIsJsonContentType(request *http.Request) bool {
	headers := request.Header

	if headers.Get("Content-Type") == "application/json" {
		return true
	} else {
		return false
	}
}
