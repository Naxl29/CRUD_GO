package controllers

import (
	"CRUD_GO/database"
	"CRUD_GO/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Ver listado de personas
func VerPersonas(w http.ResponseWriter, r *http.Request) {
	var personas []models.Persona

	// Consultar todas las personas de la base de datos
	result := database.DB.Find(&personas)
	if result.Error != nil {
		log.Println("Error al consultar personas:", result.Error)
		http.Error(w, "Error al cargar personas", http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla en minúsculas
	err := tmpl.ExecuteTemplate(w, "personas.html", personas)
	if err != nil {
		log.Println("Error al cargar:", err)
	}
}

// Crear persona
func CrearPersona(w http.ResponseWriter, r *http.Request) {
	// No hay plantilla separada para crear persona; redirigir al listado
	http.Redirect(w, r, "/personas", http.StatusSeeOther)
}

// Guardar persona
func GuardarPersona(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		persona := models.Persona{
			PrimerNombre:    r.FormValue("primerNombre"),
			SegundoNombre:   r.FormValue("segundoNombre"),
			PrimerApellido:  r.FormValue("primerApellido"),
			SegundoApellido: r.FormValue("segundoApellido"),
			Documento:       r.FormValue("documento"),
			Correo:          r.FormValue("correo"),
			Direccion:       r.FormValue("direccion"),
			Telefono:        r.FormValue("telefono"),
		}

		// Guardar en la base de datos
		result := database.DB.Create(&persona)
		if result.Error != nil {
			log.Println("Error al guardar persona:", result.Error)
			http.Error(w, "Error al guardar persona", http.StatusInternalServerError)
			return
		}

		log.Println("Persona registrada exitosamente:", persona.ID)
		http.Redirect(w, r, "/personas", http.StatusSeeOther)
	}
}

// Editar persona
func EditarPersona(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Obtener ID
		idStr := r.FormValue("id")
		var id uint64
		var err error
		if idStr != "" {
			id, err = strconv.ParseUint(idStr, 10, 32)
			if err != nil {
				log.Println("Error al convertir ID:", err)
				http.Error(w, "ID inválido", http.StatusBadRequest)
				return
			}
		}

		// Buscar la persona en la base de datos
		var persona models.Persona
		result := database.DB.First(&persona, id)
		if result.Error != nil {
			log.Println("Error al buscar persona:", result.Error)
			http.Error(w, "Persona no encontrada", http.StatusNotFound)
			return
		}

		// Actualizar los campos
		persona.PrimerNombre = r.FormValue("primerNombre")
		persona.SegundoNombre = r.FormValue("segundoNombre")
		persona.PrimerApellido = r.FormValue("primerApellido")
		persona.SegundoApellido = r.FormValue("segundoApellido")
		persona.Documento = r.FormValue("documento")
		persona.Correo = r.FormValue("correo")
		persona.Direccion = r.FormValue("direccion")
		persona.Telefono = r.FormValue("telefono")

		// Guardar cambios
		result = database.DB.Save(&persona)
		if result.Error != nil {
			log.Println("Error al actualizar persona:", result.Error)
			http.Error(w, "Error al actualizar persona", http.StatusInternalServerError)
			return
		}

		log.Println("Persona actualizada exitosamente:", persona.ID)
	}

	http.Redirect(w, r, "/personas", http.StatusSeeOther)
}

// Eliminar persona
func EliminarPersona(w http.ResponseWriter, r *http.Request) {
	// Obtener ID del query parameter
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		log.Println("Error al convertir ID:", err)
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Eliminar (soft delete con GORM)
	result := database.DB.Delete(&models.Persona{}, id)
	if result.Error != nil {
		log.Println("Error al eliminar persona:", result.Error)
		http.Error(w, "Error al eliminar persona", http.StatusInternalServerError)
		return
	}

	log.Println("Persona eliminada exitosamente:", id)
	http.Redirect(w, r, "/personas", http.StatusSeeOther)
}

// ObtenerPersonas - Endpoint para obtener lista de personas en JSON
func ObtenerPersonas(w http.ResponseWriter, r *http.Request) {
	var personas []models.Persona

	// Consultar todas las personas
	result := database.DB.Find(&personas)
	if result.Error != nil {
		log.Println("Error al consultar personas:", result.Error)
		http.Error(w, "Error al cargar personas", http.StatusInternalServerError)
		return
	}

	// Devolver JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("["))
	for i, persona := range personas {
		if i > 0 {
			w.Write([]byte(","))
		}
		nombreCompleto := persona.PrimerNombre + " " + persona.SegundoNombre + " " + persona.PrimerApellido + " " + persona.SegundoApellido
		w.Write([]byte(fmt.Sprintf(`{"id":%d,"nombre":"%s"}`, persona.ID, nombreCompleto)))
	}
	w.Write([]byte("]"))
}
