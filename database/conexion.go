package database

import (
	"log"

	"CRUD_GO/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB instancia global de la base de datos
var DB *gorm.DB

// Connect establece la conexión con MySQL
func Connect() {
	var err error

	// Obtener DSN desde la configuración
	dsn := config.Configuracion.GetDSN()

	// Configurar GORM
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	log.Println("Conexión exitosa con la base de datos MySQL")
}

// GetDB retorna la instancia de la base de datos
func GetDB() *gorm.DB {
	return DB
}

// Close cierra la conexión con la base de datos
func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Println("Error al obtener la instancia de la base de datos:", err)
		return
	}

	err = sqlDB.Close()
	if err != nil {
		log.Println("Error al cerrar la conexión:", err)
		return
	}

	log.Println("Conexión cerrada correctamente")
}
