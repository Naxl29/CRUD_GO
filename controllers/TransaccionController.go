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
		tipo := r.FormValue("tipo")

		// Si es una transferencia, manejarla de forma especial
		if tipo == "Transferencia" {
			guardarTransferencia(w, r)
			return
		}

		// Para Depósito y Retiro, procesar normalmente
		idClienteStr := r.FormValue("idCliente")
		idCliente, err := strconv.ParseUint(idClienteStr, 10, 32)
		if err != nil {
			log.Println("Error al convertir ID cliente:", err)
			http.Error(w, "ID de cliente inválido", http.StatusBadRequest)
			return
		}

		montoStr := r.FormValue("monto")
		monto, err := strconv.ParseFloat(montoStr, 64)
		if err != nil {
			log.Println("Error al convertir monto:", err)
			http.Error(w, "Monto inválido", http.StatusBadRequest)
			return
		}

		// Validar que el monto sea positivo
		if monto <= 0 {
			http.Error(w, "El monto debe ser mayor a 0", http.StatusBadRequest)
			return
		}

		// Iniciar transacción de base de datos
		tx := database.DB.Begin()

		// Obtener el cliente
		var cliente models.Cliente
		if err := tx.First(&cliente, idCliente).Error; err != nil {
			tx.Rollback()
			log.Println("Error al obtener cliente:", err)
			http.Error(w, "Cliente no encontrado", http.StatusNotFound)
			return
		}

		// Validar saldo para retiros
		if tipo == "Retiro" && cliente.Saldo < monto {
			tx.Rollback()
			http.Error(w, "Saldo insuficiente para realizar el retiro", http.StatusBadRequest)
			return
		}

		// Actualizar saldo según el tipo
		if tipo == "Depósito" {
			cliente.Saldo += monto
		} else if tipo == "Retiro" {
			cliente.Saldo -= monto
		}

		// Guardar cambios en el cliente
		if err := tx.Save(&cliente).Error; err != nil {
			tx.Rollback()
			log.Println("Error al actualizar saldo:", err)
			http.Error(w, "Error al actualizar saldo", http.StatusInternalServerError)
			return
		}

		// Crear la transacción
		transaccion := models.Transaccion{
			IDCliente:   uint(idCliente),
			Tipo:        tipo,
			Monto:       monto,
			Fecha:       time.Now(),
			Descripcion: r.FormValue("descripcion"),
		}

		if err := tx.Create(&transaccion).Error; err != nil {
			tx.Rollback()
			log.Println("Error al guardar transacción:", err)
			http.Error(w, "Error al guardar transacción", http.StatusInternalServerError)
			return
		}

		// Confirmar transacción
		tx.Commit()
		log.Println("Transacción registrada exitosamente:", transaccion.ID)
		http.Redirect(w, r, "/transacciones", http.StatusSeeOther)
	}
}

// Función auxiliar para guardar transferencias
func guardarTransferencia(w http.ResponseWriter, r *http.Request) {
	// Obtener IDs de cliente origen y destino
	idOrigenStr := r.FormValue("idCliente")
	idDestinoStr := r.FormValue("idClienteDestino")

	idOrigen, err := strconv.ParseUint(idOrigenStr, 10, 32)
	if err != nil {
		http.Error(w, "ID de cliente origen inválido", http.StatusBadRequest)
		return
	}

	idDestino, err := strconv.ParseUint(idDestinoStr, 10, 32)
	if err != nil {
		http.Error(w, "ID de cliente destino inválido", http.StatusBadRequest)
		return
	}

	// Validar que no sean el mismo cliente
	if idOrigen == idDestino {
		http.Error(w, "No puede transferir a la misma cuenta", http.StatusBadRequest)
		return
	}

	// Obtener monto
	montoStr := r.FormValue("monto")
	monto, err := strconv.ParseFloat(montoStr, 64)
	if err != nil {
		http.Error(w, "Monto inválido", http.StatusBadRequest)
		return
	}

	if monto <= 0 {
		http.Error(w, "El monto debe ser mayor a 0", http.StatusBadRequest)
		return
	}

	// Iniciar transacción de base de datos
	tx := database.DB.Begin()

	// Obtener cliente origen
	var clienteOrigen models.Cliente
	if err := tx.Preload("Persona").First(&clienteOrigen, idOrigen).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Cliente origen no encontrado", http.StatusNotFound)
		return
	}

	// Obtener cliente destino
	var clienteDestino models.Cliente
	if err := tx.Preload("Persona").First(&clienteDestino, idDestino).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Cliente destino no encontrado", http.StatusNotFound)
		return
	}

	// Validar saldo suficiente
	if clienteOrigen.Saldo < monto {
		tx.Rollback()
		http.Error(w, "Saldo insuficiente para realizar la transferencia", http.StatusBadRequest)
		return
	}

	// Actualizar saldos
	clienteOrigen.Saldo -= monto
	clienteDestino.Saldo += monto

	// Guardar cambios
	if err := tx.Save(&clienteOrigen).Error; err != nil {
		tx.Rollback()
		log.Println("Error al actualizar saldo origen:", err)
		http.Error(w, "Error al procesar transferencia", http.StatusInternalServerError)
		return
	}

	if err := tx.Save(&clienteDestino).Error; err != nil {
		tx.Rollback()
		log.Println("Error al actualizar saldo destino:", err)
		http.Error(w, "Error al procesar transferencia", http.StatusInternalServerError)
		return
	}

	// Crear descripción automática si no se proporcionó
	descripcion := r.FormValue("descripcion")
	if descripcion == "" {
		nombreDestino := clienteDestino.Persona.PrimerNombre + " " + clienteDestino.Persona.PrimerApellido
		descripcion = fmt.Sprintf("Transferencia a %s", nombreDestino)
	}

	// Crear transacción de salida (retiro del cliente origen)
	transaccionOrigen := models.Transaccion{
		IDCliente:   uint(idOrigen),
		Tipo:        "Transferencia",
		Monto:       monto,
		Fecha:       time.Now(),
		Descripcion: descripcion,
	}

	if err := tx.Create(&transaccionOrigen).Error; err != nil {
		tx.Rollback()
		log.Println("Error al crear transacción origen:", err)
		http.Error(w, "Error al registrar transferencia", http.StatusInternalServerError)
		return
	}

	// Crear transacción de entrada (depósito al cliente destino)
	nombreOrigen := clienteOrigen.Persona.PrimerNombre + " " + clienteOrigen.Persona.PrimerApellido
	transaccionDestino := models.Transaccion{
		IDCliente:   uint(idDestino),
		Tipo:        "Transferencia",
		Monto:       monto,
		Fecha:       time.Now(),
		Descripcion: fmt.Sprintf("Transferencia recibida de %s", nombreOrigen),
	}

	if err := tx.Create(&transaccionDestino).Error; err != nil {
		tx.Rollback()
		log.Println("Error al crear transacción destino:", err)
		http.Error(w, "Error al registrar transferencia", http.StatusInternalServerError)
		return
	}

	// Confirmar todas las operaciones
	tx.Commit()
	log.Printf("Transferencia exitosa: $%.2f de Cliente %d a Cliente %d\n", monto, idOrigen, idDestino)
	http.Redirect(w, r, "/transacciones", http.StatusSeeOther)
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

// Mostrar formulario de transferencia
func TransferirSaldo(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	var clientes []models.Cliente
	if err := db.Preload("Persona").Find(&clientes).Error; err != nil {
		http.Error(w, "Error al cargar clientes", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Clientes": clientes,
	}

	err := tmpl.ExecuteTemplate(w, "transferir.html", data)
	if err != nil {
		log.Println("Error al renderizar template:", err)
	}
}

func GuardarTransferencia(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/transacciones/transferir", http.StatusSeeOther)
		return
	}

	db := database.GetDB()

	// Obtener valores del formulario
	idOrigen, _ := strconv.ParseUint(r.FormValue("origen"), 10, 64)
	idDestino, _ := strconv.ParseUint(r.FormValue("destino"), 10, 64)
	monto, _ := strconv.ParseFloat(r.FormValue("monto"), 64)

	// Verificar que los IDs sean diferentes
	if idOrigen == idDestino {
		http.Error(w, "No puedes transferirte a ti mismo.", http.StatusBadRequest)
		return
	}

	// Obtener los clientes
	var origen, destino models.Cliente
	if err := db.First(&origen, idOrigen).Error; err != nil {
		http.Error(w, "Cliente origen no encontrado", http.StatusNotFound)
		log.Println("Error cliente origen:", err)
		return
	}
	if err := db.First(&destino, idDestino).Error; err != nil {
		http.Error(w, "Cliente destino no encontrado", http.StatusNotFound)
		log.Println("Error cliente destino:", err)
		return
	}

	// Verificar saldo suficiente
	if origen.Saldo < monto {
		http.Error(w, "Saldo insuficiente en la cuenta origen", http.StatusBadRequest)
		return
	}

	// Mostrar info en consola
	log.Println("=== DATOS TRANSFERENCIA ===")
	log.Println("Origen:", origen.ID, "Saldo:", origen.Saldo)
	log.Println("Destino:", destino.ID, "Saldo:", destino.Saldo)
	log.Println("Monto:", monto)
	log.Println("===========================")

	// Iniciar transacción
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Actualizar saldos
	origen.Saldo -= monto
	destino.Saldo += monto

	if err := tx.Save(&origen).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Error al actualizar cuenta origen", http.StatusInternalServerError)
		log.Println("Error al guardar origen:", err)
		return
	}

	if err := tx.Save(&destino).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Error al actualizar cuenta destino", http.StatusInternalServerError)
		log.Println("Error al guardar destino:", err)
		return
	}

	// Crear registros de transacción
	transOrigen := models.Transaccion{
		IDCliente:   origen.ID,
		Tipo:        "Transferencia salida",
		Monto:       -monto,
		Fecha:       time.Now(),
		Descripcion: "Transferencia a cliente " + strconv.Itoa(int(destino.ID)),
	}

	transDestino := models.Transaccion{
		IDCliente:   destino.ID,
		Tipo:        "Transferencia entrada",
		Monto:       monto,
		Fecha:       time.Now(),
		Descripcion: "Transferencia desde cliente " + strconv.Itoa(int(origen.ID)),
	}

	// Crear transacciones
	if err := tx.Create(&transOrigen).Error; err != nil {
		tx.Rollback()
		log.Println("Error detalle transOrigen:", err)
		http.Error(w, "Error al registrar transacción origen", http.StatusInternalServerError)
		return
	}

	if err := tx.Create(&transDestino).Error; err != nil {
		tx.Rollback()
		log.Println("Error detalle transDestino:", err)
		http.Error(w, "Error al registrar transacción destino", http.StatusInternalServerError)
		return
	}

	// Confirmar transacción
	if err := tx.Commit().Error; err != nil {
		log.Println("Error al hacer commit:", err)
		http.Error(w, "Error al finalizar la transferencia", http.StatusInternalServerError)
		return
	}

	log.Println("Transferencia completada correctamente.")
	http.Redirect(w, r, "/transacciones", http.StatusSeeOther)
}
