package GoHttpService

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func LoadHtml(pageName string) ([]byte, error) {
	fileName := pageName + ".html"
	return ioutil.ReadFile(fileName)
}

func RenderHtml(writer http.ResponseWriter, pageName string) error {
	body, err := LoadHtml(pageName)
	if err != nil {
		return err
	}

	fmt.Fprint(writer, body)
	return nil
}

func RenderTemplate(writer http.ResponseWriter, pageName string, data interface{}) error {
	t, err := template.ParseFiles(pageName + ".html")
	if err != nil {
		return err
	}

	err = t.Execute(writer, data)
	return err
}
