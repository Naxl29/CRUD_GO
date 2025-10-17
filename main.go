package main

import (
    "fmt"
    "log"
    "net/http"

    "CRUD_GO/config"
    "CRUD_GO/routes"
)

func main() {
	fmt.Println("🚀 Iniciando servidor...")

	// Conectar base de datos
	config.ConnectDB()

	// Configurar rutas
	r := routes.SetupRoutes()

	// Iniciar servidor
	fmt.Println("✅ Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
