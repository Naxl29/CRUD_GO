package main

import (
    "net/http"
    "log"
    //"fmt"

    //"CRUD_GO/config"
    "CRUD_GO/routes"
)

func main() {
	// Acá se cargan todas las rutas definidas en routes/routes.go
	routes.LoadRoutes()

	// Iniciar el servidor, vamos a utilizar el puerto 8080 para que lo tengan en cuenta muchachos
	log.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}