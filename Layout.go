package GoHttpService

import (
	"html/template"
	"net/http"
)

type Layout struct {
	BaseDirectory     string
	TemplateFileNames []string
	TemplateBodies    map[string]string
}

func NewLayout(baseDirectory string, templateFileNames []string) (Layout, error) {
	result := Layout{TemplateFileNames: templateFileNames}
	result.BaseDirectory = baseDirectory
	result.TemplateBodies = make(map[string]string)

	for _, templateName := range templateFileNames {
		html, err := LoadHtml(baseDirectory + templateName)
		if err != nil {
			return result, err
		}

		result.TemplateBodies[templateName] = string(html)
	}

	return result, nil
}

func (this Layout) parseTemplates(body string, data interface{}) (*template.Template, error) {
	t := template.New("layout")
	var err error

	for _, html := range this.TemplateBodies {
		t, err = t.Parse(html)
		if err != nil {
			return t, err
		}
	}

	t, err = t.Parse(body)
	return t, nil
}

func (this Layout) RenderView(writer http.ResponseWriter, pageName string, data interface{}) error {
	writer.Header().Set("Content-Type", "text/html; charset=UTF-8")

	html, err := LoadHtml(this.BaseDirectory + pageName)
	if err != nil {
		return err
	}

	t, err := this.parseTemplates(string(html), data)
	if err != nil {
		return err
	}

	return t.Execute(writer, data)
}
