package render

import (
	"bookings/pkg/models"
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")

	err := parsedTemplate.Execute(w, td)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}
