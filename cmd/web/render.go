package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type templateData struct {
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, t string, tData *templateData) {

	var tmpl *template.Template
	if app.config.useCache {
		if templateFromMap, ok := app.templateMap[t]; ok {
			tmpl = templateFromMap
		}
	}
	if tmpl == nil {
		newTemplate, err := app.buildTemplateFromDisk(t)
		if err != nil {
			log.Println("Error building template", err)
			return
		}
		log.Println("Building template from disk")
		tmpl = newTemplate
	}

	if tData == nil {
		tData = &templateData{}
	}

	if err := tmpl.ExecuteTemplate(w, t, tData); err != nil {
		log.Println("Error executing template", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (app *application) buildTemplateFromDisk(t string) (*template.Template, error) {

	templateSlice := []string{
		"./templates/base.layout.gohtml",
		"./templates/partials/footer.partials.gohtml",
		"./templates/partials/header.partial.gohtml",
		fmt.Sprintf("./templates/%s", t),
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		return nil, err
	}
	app.templateMap[t] = tmpl
	return tmpl, nil
}
