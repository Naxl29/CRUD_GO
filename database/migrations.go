package database

import (
	"log"

	"CRUD_GO/models"
)

// Migrate ejecuta las migraciones de la base de datos
func Migrate() {
	log.Println(" Ejecutando migraciones...")

	err := DB.AutoMigrate(
		&models.Persona{},
		&models.Cliente{},
		&models.Transaccion{},
	)

	if err != nil {
		log.Fatal("Error al ejecutar migraciones:", err)
	}

	log.Println("Migraciones ejecutadas exitosamente")
}
