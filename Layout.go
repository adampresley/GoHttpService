package GoHttpService

import (
	"html/template"
	"net/http"
)

type Layout struct {
	BaseDirectory     string
	TemplateFileNames []string
	RenderedTemplates *template.Template
}

func NewLayout(baseDirectory string, templateFileNames []string) (*Layout, error) {
	result := &Layout{TemplateFileNames: templateFileNames}
	t := template.New("layout")

	for _, templateName := range templateFileNames {
		html, err := LoadHtml(baseDirectory + templateName)
		if err != nil {
			return &Layout{}, err
		}

		t, err = t.Parse(string(html))
		if err != nil {
			return &Layout{}, err
		}
	}

	result.BaseDirectory = baseDirectory
	result.RenderedTemplates = t
	return result, nil
}

func (this *Layout) RenderView(writer http.ResponseWriter, pageName string, data interface{}) error {
	writer.Header().Set("Content-Type", "text/html; charset=UTF-8")

	html, err := LoadHtml(this.BaseDirectory + pageName)
	if err != nil {
		return err
	}

	t, _ := this.RenderedTemplates.Parse(string(html))
	return t.Execute(writer, data)
}
