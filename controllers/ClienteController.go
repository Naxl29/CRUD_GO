package controllers

import (
	"CRUD_GO/database"
	"CRUD_GO/models"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Mostrar lista de clientes
func VerClientes(w http.ResponseWriter, r *http.Request) {
	var clientes []models.Cliente

	// Consultar clientes con sus personas asociadas
	result := database.DB.Preload("Persona").Find(&clientes)
	if result.Error != nil {
		log.Println("Error al consultar clientes:", result.Error)
		http.Error(w, "Error al cargar clientes", http.StatusInternalServerError)
		return
	}

	// Renderizar la plantilla de listado en minúsculas
	err := tmpl.ExecuteTemplate(w, "clientes.html", clientes)
	if err != nil {
		log.Println("Error al cargar clientes:", err)
	}
}

// Mostrar formulario de creación
func CrearCliente(w http.ResponseWriter, r *http.Request) {
	// No existe un archivo independiente para crear cliente; redirigir al listado
	http.Redirect(w, r, "/clientes", http.StatusSeeOther)
}

// Guardar cliente
func GuardarCliente(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Obtener y convertir IDPersona
		idPersonaStr := r.FormValue("idPersona")
		idPersona, err := strconv.ParseUint(idPersonaStr, 10, 32)
		if err != nil {
			log.Println("Error al convertir ID persona:", err)
			http.Error(w, "ID de persona inválido", http.StatusBadRequest)
			return
		}

		// Obtener y convertir Saldo
		saldoStr := r.FormValue("saldo")
		saldo, err := strconv.ParseFloat(saldoStr, 64)
		if err != nil {
			log.Println("Error al convertir saldo:", err)
			http.Error(w, "Saldo inválido", http.StatusBadRequest)
			return
		}

		// Crear el cliente
		cliente := models.Cliente{
			IDPersona:     uint(idPersona),
			TipoCuenta:    r.FormValue("tipoCuenta"),
			Saldo:         saldo,
			FechaRegistro: time.Now(),
		}

		// Guardar en la base de datos
		result := database.DB.Create(&cliente)
		if result.Error != nil {
			log.Println("Error al guardar cliente:", result.Error)
			http.Error(w, "Error al guardar cliente", http.StatusInternalServerError)
			return
		}

		log.Println("Cliente registrado exitosamente:", cliente.ID)
		http.Redirect(w, r, "/clientes", http.StatusSeeOther)
	}
}

// Editar cliente
func EditarCliente(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Obtener ID del cliente
		idStr := r.FormValue("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			log.Println("Error al convertir ID:", err)
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		// Obtener y convertir IDPersona
		idPersonaStr := r.FormValue("idPersona")
		idPersona, err := strconv.ParseUint(idPersonaStr, 10, 32)
		if err != nil {
			log.Println("Error al convertir ID persona:", err)
			http.Error(w, "ID de persona inválido", http.StatusBadRequest)
			return
		}

		// Obtener y convertir Saldo
		saldoStr := r.FormValue("saldo")
		saldo, err := strconv.ParseFloat(saldoStr, 64)
		if err != nil {
			log.Println("Error al convertir saldo:", err)
			http.Error(w, "Saldo inválido", http.StatusBadRequest)
			return
		}

		// Buscar el cliente en la base de datos
		var cliente models.Cliente
		result := database.DB.First(&cliente, id)
		if result.Error != nil {
			log.Println("Error al buscar cliente:", result.Error)
			http.Error(w, "Cliente no encontrado", http.StatusNotFound)
			return
		}

		// Actualizar los campos
		cliente.IDPersona = uint(idPersona)
		cliente.TipoCuenta = r.FormValue("tipoCuenta")
		cliente.Saldo = saldo

		// Guardar cambios
		result = database.DB.Save(&cliente)
		if result.Error != nil {
			log.Println("Error al actualizar cliente:", result.Error)
			http.Error(w, "Error al actualizar cliente", http.StatusInternalServerError)
			return
		}

		log.Println("Cliente actualizado exitosamente:", cliente.ID)
	}

	http.Redirect(w, r, "/clientes", http.StatusSeeOther)
}

// Eliminar cliente
func EliminarCliente(w http.ResponseWriter, r *http.Request) {
	// Obtener ID del query parameter
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		log.Println("Error al convertir ID:", err)
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Eliminar (soft delete con GORM)
	result := database.DB.Delete(&models.Cliente{}, id)
	if result.Error != nil {
		log.Println("Error al eliminar cliente:", result.Error)
		http.Error(w, "Error al eliminar cliente", http.StatusInternalServerError)
		return
	}

	log.Println("Cliente eliminado exitosamente:", id)
	http.Redirect(w, r, "/clientes", http.StatusSeeOther)
}
