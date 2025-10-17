package database

import (
	"log"

	"github.com/Naxl29/CRUD_GO/models"
)

// Migrate ejecuta las migraciones de la base de datos
func Migrate() {
	log.Println("Ejecutando migraciones...")

	// AutoMigrate crea las tablas en orden
	err := DB.AutoMigrate(
		&models.Cliente{},
		&models.Cuenta{},
		&models.Transaccion{},
	)

	if err != nil {
		log.Fatal("Error al ejecutar migraciones:", err)
	}

	log.Println("Migraciones ejecutadas exitosamente")
}
