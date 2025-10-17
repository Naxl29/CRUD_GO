package routes

import (
	"net/http"
	"CRUD_GO/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/personas", controllers.VerPersonas)
	http.HandleFunc("/personas/crear", controllers.CrearPersona)
	http.HandleFunc("/personas/guardar", controllers.GuardarPersona)
	http.HandleFunc("/personas/editar", controllers.EditarPersona)
	http.HandleFunc("/personas/eliminar", controllers.EliminarPersona)
}
