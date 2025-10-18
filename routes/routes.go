package routes

import (
	"net/http"
	"CRUD_GO/controllers"
)

func LoadRoutes() {
	// Rutas para Personas
	http.HandleFunc("/personas", controllers.VerPersonas)
	http.HandleFunc("/personas/crear", controllers.CrearPersona)
	http.HandleFunc("/personas/guardar", controllers.GuardarPersona)
	http.HandleFunc("/personas/editar", controllers.EditarPersona)
	http.HandleFunc("/personas/eliminar", controllers.EliminarPersona)

	// Rutas para Clientes
	http.HandleFunc("/clientes", controllers.VerClientes)
	http.HandleFunc("/clientes/crear", controllers.CrearCliente)
	http.HandleFunc("/clientes/guardar", controllers.GuardarCliente)
	http.HandleFunc("/clientes/editar", controllers.EditarCliente)
	http.HandleFunc("/clientes/eliminar", controllers.EliminarCliente)
}
