package controllers

import (
	"CRUD_GO/database"
	"CRUD_GO/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Mostrar lista de transacciones
func VerTransacciones(w http.ResponseWriter, r *http.Request) {
	var transacciones []models.Transaccion

	// Consultar transacciones con relaciones (Cliente y Persona del cliente)
	result := database.DB.Preload("Cliente.Persona").Find(&transacciones)
	if result.Error != nil {
		log.Println("Error al consultar transacciones:", result.Error)
		http.Error(w, "Error al cargar transacciones", http.StatusInternalServerError)
		return
	}

	// Renderizar plantilla en minúsculas
	err := tmpl.ExecuteTemplate(w, "transacciones.html", transacciones)
	if err != nil {
		log.Println("Error al renderizar:", err)
	}
}

// Crear transacción - Obtener lista de clientes para el formulario
func CrearTransaccion(w http.ResponseWriter, r *http.Request) {
	// Redirigir al listado (el modal está en la misma página)
	http.Redirect(w, r, "/transacciones", http.StatusSeeOther)
}

// ObtenerClientes - Endpoint para obtener lista de clientes con sus personas
func ObtenerClientes(w http.ResponseWriter, r *http.Request) {
	var clientes []models.Cliente

	// Consultar clientes con sus personas asociadas
	result := database.DB.Preload("Persona").Find(&clientes)
	if result.Error != nil {
		log.Println("Error al consultar clientes:", result.Error)
		http.Error(w, "Error al cargar clientes", http.StatusInternalServerError)
		return
	}

	// Construir lista de clientes válidos
	var clientesJSON []string

	for _, cliente := range clientes {
		// Verificar que la Persona exista
		if cliente.Persona.ID == 0 {
			continue
		}

		// Construir nombre completo
		nombreCompleto := cliente.Persona.PrimerNombre
		if cliente.Persona.SegundoNombre != "" {
			nombreCompleto += " " + cliente.Persona.SegundoNombre
		}
		nombreCompleto += " " + cliente.Persona.PrimerApellido
		if cliente.Persona.SegundoApellido != "" {
			nombreCompleto += " " + cliente.Persona.SegundoApellido
		}

		clienteJSON := fmt.Sprintf(`{"id":%d,"nombre":"%s"}`, cliente.ID, nombreCompleto)
		clientesJSON = append(clientesJSON, clienteJSON)
	}

	// Devolver JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("["))
	for i, clienteJSON := range clientesJSON {
		if i > 0 {
			w.Write([]byte(","))
		}
		w.Write([]byte(clienteJSON))
	}
	w.Write([]byte("]"))
}

// Guardar transacción
func GuardarTransaccion(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Obtener y convertir el ID del cliente
		idClienteStr := r.FormValue("idCliente")
		idCliente, err := strconv.ParseUint(idClienteStr, 10, 32)
		if err != nil {
			log.Println("Error al convertir ID cliente:", err)
			http.Error(w, "ID de cliente inválido", http.StatusBadRequest)
			return
		}

		// Obtener y convertir el monto
		montoStr := r.FormValue("monto")
		monto, err := strconv.ParseFloat(montoStr, 64)
		if err != nil {
			log.Println("Error al convertir monto:", err)
			http.Error(w, "Monto inválido", http.StatusBadRequest)
			return
		}

		transaccion := models.Transaccion{
			IDCliente:   uint(idCliente),
			Tipo:        r.FormValue("tipo"),
			Monto:       monto,
			Fecha:       time.Now(),
			Descripcion: r.FormValue("descripcion"),
		}

		// Guardar en la base de datos
		result := database.DB.Create(&transaccion)
		if result.Error != nil {
			log.Println("Error al guardar transacción:", result.Error)
			http.Error(w, "Error al guardar transacción", http.StatusInternalServerError)
			return
		}

		log.Println("Transacción registrada exitosamente:", transaccion.ID)
		http.Redirect(w, r, "/transacciones", http.StatusSeeOther)
	}
}

// Editar transacción
func EditarTransaccion(w http.ResponseWriter, r *http.Request) {
	// Redirigir al listado (no hay plantilla separada)
	http.Redirect(w, r, "/transacciones", http.StatusSeeOther)
}

// Eliminar transacción
func EliminarTransaccion(w http.ResponseWriter, r *http.Request) {
	log.Println("Transacción eliminada")
	http.Redirect(w, r, "/transacciones", http.StatusSeeOther)
}
