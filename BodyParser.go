package GoHttpService

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseJsonBody(request *http.Request, receiver interface{}) (error) {
	var err error

	body, _ := ioutil.ReadAll(request.Body)
	err = json.Unmarshal(body, receiver)
	return err
}
