package controllers

import (
	"net/http"
	"CRUD_GO/models"
	"log"
	"time"
)

// Mostrar lista de transacciones
func VerTransacciones(w http.ResponseWriter, r *http.Request){
	transacciones := []models.Transaccion{
		{ID: 1, IDCliente: 1, Tipo: "Depósito", Monto: 500000.00, Fecha: time.Now(), Descripcion: "Depósito inicial"},
	}
	err := tmpl.ExecuteTemplate(w, "Transacciones.html", transacciones)
	if err != nil {
		log.Println(err)
	}
}

// Crear transacción
func CrearTransaccion(w http.ResponseWriter, r *http.Request){
	tmpl.ExecuteTemplate(w, "crearTransaccion.html", nil)
}

// Guardar transacción
func GuardarTransaccion(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		transaccion := models.Transaccion{
			IDCliente: 1,
			Tipo: 	r.FormValue("tipo"),
			Monto: 	0.0,
			Fecha: 	time.Now(),
			Descripcion: r.FormValue("descripcion"),
		}
		log.Println("Transacción registrada:", transaccion)
		http.Redirect(w, r, "/transacciones", http.StatusSeeOther)
		}
}

// Editar transacción
func EditarTransaccion(w http.ResponseWriter, r *http.Request){
	tmpl.ExecuteTemplate(w, "editarTransaccion.html", nil)
}

// Eliminar transacción
func EliminarTransaccion(w http.ResponseWriter, r *http.Request){
	log.Println("Transacción eliminada")
	http.Redirect(w, r, "/transacciones", http.StatusSeeOther)
}