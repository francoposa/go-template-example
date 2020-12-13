package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

func NewTemplates(pattern, baseTemplatePath string) map[string]*template.Template {
	templates := make(map[string]*template.Template)

	templatePaths, err := filepath.Glob(pattern)
	if err != nil {
		log.Fatal(err)
	}

	for _, templatePath := range templatePaths {
		if templatePath == baseTemplatePath {
			continue
		}
		fileName := filepath.Base(templatePath)
		fileNameNoExt := strings.TrimSuffix(fileName, filepath.Ext(fileName))
		templates[fileNameNoExt] = template.Must(
			template.ParseFiles(templatePath, baseTemplatePath),
		)
	}

	return templates
}

type TemplateRenderer struct {
	templates        map[string]*template.Template
	baseTemplateName string
}

func NewTemplateRenderer(
	templates map[string]*template.Template,
	baseTemplateName string,
) *TemplateRenderer {
	return &TemplateRenderer{templates: templates, baseTemplateName: baseTemplateName}
}

func (tr *TemplateRenderer) RenderTemplate(w http.ResponseWriter, templateFileName string, data interface{}) error {
	templateToRender, ok := tr.templates[templateFileName]
	if !ok {
		return TemplateNotFoundError{Name: templateFileName}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return templateToRender.ExecuteTemplate(w, tr.baseTemplateName, data)
}

type TemplateNotFoundError struct {
	Name string
}

func (e TemplateNotFoundError) Error() string {
	return fmt.Sprintf("No Template found with Name %s", e.Name)
}
