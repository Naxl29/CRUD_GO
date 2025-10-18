package main

import (
	"log"
	"net/http"

	"CRUD_GO/config"
	"CRUD_GO/database"
	"CRUD_GO/routes"
)

func main() {
	
	// 1. Cargar configuración desde .env
	config.Variables_entorno()

	// 2. Conectar a la base de datos MySQL
	database.Connect()

	// 3. Ejecutar migraciones
	database.Migrate()

	// 4. Servir archivos estáticos
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 5. Cargar todas las rutas
	routes.LoadRoutes()

	// Iniciar el servidor en el puerto 8080
	log.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
