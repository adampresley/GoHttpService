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
func ParseJsonBody(request *http.Request, receiver interface{}) (error) {
	var err error

	body, _ := ioutil.ReadAll(request.Body)
	err = json.Unmarshal(body, receiver)
	return err
}
