package controllers

import (
	"CRUD_GO/models"
	"log"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

// Ver listado de personas
func VerPersonas(w http.ResponseWriter, r *http.Request) {
	personas := []models.Persona{
		{ID: 1, PrimerNombre: "NICOLÁS", SegundoNombre: "ANTONIO", PrimerApellido: "ARRIETA", SegundoApellido: "LAGOS", Documento: "1099735735", Correo: "NICOLAS@GMAIL.COM", Direccion: "CALLE 29", Telefono: "3165041963"},
	}
	err := tmpl.ExecuteTemplate(w, "Personas.html", personas)
	if err != nil {
		log.Println("Error al cargar:", err)
	}
}

// Crear persona
func CrearPersona(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "crearPersona.html", nil)
}

// Guardar persona
func GuardarPersona(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		persona := models.Persona{
			PrimerNombre:  r.FormValue("primerNombre"),
			SegundoNombre: r.FormValue("segundoNombre"),
			PrimerApellido: r.FormValue("primerApellido"),
			SegundoApellido: r.FormValue("segundoApellido"),
			Documento:     r.FormValue("documento"),
			Correo:       r.FormValue("correo"),
			Direccion:    r.FormValue("direccion"),
			Telefono:    r.FormValue("telefono"),
		}
		log.Println("Persona registrada:", persona)
		http.Redirect(w, r, "/personas", http.StatusSeeOther)
	}
}

// Editar persona
func EditarPersona(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "editarPersona.html", nil)
}

// Eliminar persona
func EliminarPersona(w http.ResponseWriter, r *http.Request) {
	log.Println("Persona eliminada")
	http.Redirect(w, r, "/personas", http.StatusSeeOther)
}