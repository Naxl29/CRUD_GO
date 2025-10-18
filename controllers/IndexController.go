package controllers

import (
	"CRUD_GO/database"
	"CRUD_GO/models"
	"log"
	"net/http"
)

// VerIndex muestra la página principal con el listado de cuentas/clientes
func VerIndex(w http.ResponseWriter, r *http.Request) {
	var clientes []models.Cliente

	// Consultar clientes con su relación Persona (usando Preload)
	result := database.DB.Preload("Persona").Find(&clientes)

	if result.Error != nil {
		log.Println("Error al cargar clientes:", result.Error)
		// Mostrar página vacía en caso de error
		clientes = []models.Cliente{}
	}

	// Crear estructura de datos para el template
	data := struct {
		Clientes []models.Cliente
	}{
		Clientes: clientes,
	}

	err := tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Println("Error al renderizar index:", err)
		http.Error(w, "Error al cargar la página", http.StatusInternalServerError)
	}
}
