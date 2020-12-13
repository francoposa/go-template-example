package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

//func NewTemplates(pattern, baseTemplatePath string) map[string]*template.Template {
//	templates := make(map[string]*template.Template)
//
//	templatePaths, err := filepath.Glob(pattern)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, templatePath := range templatePaths {
//		if templatePath == baseTemplatePath {
//			continue
//		}
//		fileName := filepath.Base(templatePath)
//		templates[fileName] = template.Must(
//			template.ParseFiles(templatePath, baseTemplatePath),
//		)
//	}
//
//	return templates
//}

func NewTemplates(pattern, baseTemplatePath string) *template.Template {
	templates := template.New("http_templates")

	templatePaths, err := filepath.Glob(pattern)
	if err != nil {
		log.Fatal(err)
	}

	templates = template.Must(
		template.ParseFiles(templatePaths...),
	)

	//for _, templatePath := range templatePaths {
	//	//if templatePath == baseTemplatePath {
	//	//	continue
	//	//}
	//	//fileName := filepath.Base(templatePath)
	//	//fileNameNoExt := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	//	templates = template.Must(
	//		template.ParseFiles(templatePath),
	//	)
	//}

	return templates
}

type TemplateRenderer struct {
	templates        *template.Template
	baseTemplateName string
}

func NewTemplateRenderer(
	templates *template.Template,
	baseTemplateName string,
) *TemplateRenderer {
	return &TemplateRenderer{templates: templates, baseTemplateName: baseTemplateName}
}

func (tr *TemplateRenderer) RenderTemplate(w http.ResponseWriter, templateFileName string, data interface{}) error {
	//templateToRender, ok := tr.templates[templateFileName]
	//if !ok {
	//	return TemplateNotFoundError{Name: templateFileName}
	//}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return tr.templates.ExecuteTemplate(w, templateFileName, data)
}

type TemplateNotFoundError struct {
	Name string
}

func (e TemplateNotFoundError) Error() string {
	return fmt.Sprintf("No Template found with Name %s", e.Name)
}
