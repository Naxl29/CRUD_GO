package routes

import (
	"CRUD_GO/controllers"
	"net/http"
)

func LoadRoutes() {
	// Ruta principal - Página de inicio con listado de cuentas
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			controllers.VerIndex(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

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

	// Rutas para Transacciones
	http.HandleFunc("/transacciones", controllers.VerTransacciones)
	http.HandleFunc("/transacciones/crear", controllers.CrearTransaccion)
	http.HandleFunc("/transacciones/guardar", controllers.GuardarTransaccion)
	http.HandleFunc("/transacciones/editar", controllers.EditarTransaccion)
	http.HandleFunc("/transacciones/eliminar", controllers.EliminarTransaccion)
<<<<<<< HEAD

	// API para obtener datos para formularios
	http.HandleFunc("/api/clientes", controllers.ObtenerClientes)
	http.HandleFunc("/api/personas", controllers.ObtenerPersonas)
=======
	// Transferencia entre clientes
	http.HandleFunc("/transacciones/transferir", controllers.TransferirSaldo)
	http.HandleFunc("/transacciones/guardarTransferencia", controllers.GuardarTransferencia)


>>>>>>> feature/transaccion
}
