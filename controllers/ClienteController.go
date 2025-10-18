package controllers

import (
	"CRUD_GO/models"
	"log"
	"net/http"
	"time"
)

// Mostrar lista de clientes
func VerClientes(w http.ResponseWriter, r *http.Request) {
	clientes := []models.Cliente{
		{ID: 1, IDPersona: 1, TipoCuenta: "Ahorros", Saldo: 1000000.00, FechaRegistro: time.Now()},
	}
	err := tmpl.ExecuteTemplate(w, "Clientes.html", clientes)
	if err != nil {
		log.Println("Error al cargar clientes:", err)
	}
}

// Mostrar formulario de creación
func CrearCliente(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "crearCliente.html", nil)
}

// Guardar cliente 
func GuardarCliente(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		cliente := models.Cliente{
			IDPersona:     1, // Esto será reemplazado con un IDPersona real
			TipoCuenta:    r.FormValue("tipoCuenta"),
			Saldo:         0.0,
			FechaRegistro: time.Now(), 
		}

		log.Println("Cliente registrado:", cliente)
		http.Redirect(w, r, "/clientes", http.StatusSeeOther)
	}
}

// Editar cliente 
func EditarCliente(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "editarCliente.html", nil)
}

// Eliminar cliente
func EliminarCliente(w http.ResponseWriter, r *http.Request) {
	log.Println("Cliente eliminado")
	http.Redirect(w, r, "/clientes", http.StatusSeeOther)
}