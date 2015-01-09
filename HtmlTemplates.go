package GoHttpService

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func LoadHtml(pageName string) ([]byte, error) {
	fileName := pageName + ".html"
	log.Println("INFO - Loading file", fileName)
	return ioutil.ReadFile(fileName)
}

func RenderHtml(writer http.ResponseWriter, pageName string) error {
	body, err := LoadHtml(pageName)
	if err != nil {
		return err
	}

	writer.Header().Set("Content-Type", "text/html; charset=UTF-8")

	fmt.Fprintf(writer, string(body))
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
