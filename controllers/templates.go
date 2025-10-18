package controllers

import (
	"net/http"
	"text/template"
)

// Variable global compartida por todos los controladores
var tmpl = template.Must(template.ParseGlob("templates/*.html"))

// RenderTemplate función exportada para renderizar templates desde otros paquetes
func RenderTemplate(w http.ResponseWriter, templateName string, data interface{}) error {
	return tmpl.ExecuteTemplate(w, templateName, data)
}
