package controllers

import (
	"CRUD_GO/models"
	"log"
	"net/http"
	"time"
)

// Mostrar lista de transacciones
func VerTransacciones(w http.ResponseWriter, r *http.Request) {
	// Crear transacción de ejemplo sin ID (GORM lo asigna automáticamente)
	transaccion := models.Transaccion{
		IDCliente:   1,
		Tipo:        "Depósito",
		Monto:       500000.00,
		Fecha:       time.Now(),
		Descripcion: "Depósito inicial",
	}
	transaccion.ID = 1 // Asignar ID después de crear la estructura

	transacciones := []models.Transaccion{transaccion}

	err := tmpl.ExecuteTemplate(w, "Transacciones.html", transacciones)
	if err != nil {
		log.Println(err)
	}
}

// Crear transacción
func CrearTransaccion(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "crearTransaccion.html", nil)
}

// Guardar transacción
func GuardarTransaccion(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		transaccion := models.Transaccion{
			IDCliente:   1,
			Tipo:        r.FormValue("tipo"),
			Monto:       0.0,
			Fecha:       time.Now(),
			Descripcion: r.FormValue("descripcion"),
		}
		log.Println("Transacción registrada:", transaccion)
		http.Redirect(w, r, "/transacciones", http.StatusSeeOther)
	}
}

// Editar transacción
func EditarTransaccion(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "editarTransaccion.html", nil)
}

// Eliminar transacción
func EliminarTransaccion(w http.ResponseWriter, r *http.Request) {
	log.Println("Transacción eliminada")
	http.Redirect(w, r, "/transacciones", http.StatusSeeOther)
}
