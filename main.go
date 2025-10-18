package main

import (
	"log"
	"net/http"

	"CRUD_GO/config"
	"CRUD_GO/database"
	"CRUD_GO/routes"
)

func main() {
	log.Println("🚀 Iniciando servidor...")

	// 1. Cargar configuración desde .env
	config.Variables_entorno()

	// 2. Conectar a la base de datos MySQL
	database.Connect()

	// 3. Ejecutar migraciones (crear tabla personas automáticamente)
	database.Migrate()

	// 4. Cargar todas las rutas definidas en routes/routes.go
	routes.LoadRoutes()

	// Iniciar el servidor, vamos a utilizar el puerto 8080 para que lo tengan en cuenta muchachos
	log.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
