package controllers

import (
	"text/template"
)

// Variable global compartida por todos los controladores
var tmpl = template.Must(template.ParseGlob("templates/*.html"))

