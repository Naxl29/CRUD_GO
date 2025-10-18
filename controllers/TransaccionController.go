package controllers

import (
	"CRUD_GO/models"
	"log"
	"net/http"
	"time"
	"strconv"
	"CRUD_GO/database"
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
