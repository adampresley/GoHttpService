Go HttpService
==============

HTTP methods I use a lot.

## Basic HTTP Responses
This package offers a simple structure for sending back basic HTTP responses. This structure has two keys: *Success* and *Message*. To create a message one can use the method **NewBasicJsonResponse()** like so.

```golang
import (
    "github.com/adampresley/GoHttpService"
)

func DoSomething() {
    response := GoHttpService.NewBasicJsonResponse(true, "Everything worked great")
}
```

There are simplier methods for the most common operations such as responding with an error or success. The following samples illustrate these.

```golang
import (
    "net/http"

    "github.com/adampresley/GoHttpService"
)

/*
{ "success": false, "message": "An error occurred" }
*/
func AnError(writer http.ResponseWriter, request *http.Request) {
    GoHttpService.Error(writer, "An error occurred")
}

/*
{ "success": false, "message": "Bad request sir!" }
*/
func BadRequest(writer http.ResponseWriter, request *http.Request) {
    GoHttpService.BadRequest(writer, "Bad request sir!")
}

/*
{ "success": false, "message": "Not found!" }
*/
func NotFound(writer http.ResponseWriter, request *http.Request) {
    GoHttpService.NotFound(writer, "Not found!")
}

/*
{ "success": true, "message": "Everything is awesome!" }
*/
func NotFound(writer http.ResponseWriter, request *http.Request) {
    GoHttpService.Success(writer, "Everything is awesome!")
}
```

## Writing Arbitrary Data
Very often you'll need to write some arbitrary structure out as a response. Here is an example of doing that.

```golang
type Something struct {
    Id int `json:"id"`
    Name string `json:"name"`
}

/*
{ "id": 1, "name": "Bob" }
*/
func Arbitrary(writer http.ResponseWriter, request *http.Request) {
    result := Something{Id: 1, Name: "Bob"}
    GoHttpService.WriteJson(writer, result, 200)
}
```

## Parsing Information
Below are examples of functions for parsing data.

```golang
type Something struct {
    Id int `json:"id"`
    Name string `json:"name"`
}

/*
With a body of:

{"id": 2, "name": "Adam"}

result will be filled in with values from the parsed body
*/
func ParseJsonInBodySample(writer http.ResponseWriter, request *http.Request) {
    result := Something{}
    GoHttpService.ParseJsonBody(request, &result)
}

```

## License
The MIT License (MIT)

Copyright (c) 2014 Adam Presley

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

