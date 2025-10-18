package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config estructura para almacenar la configuración
type Config struct {
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	ServerPort  string
	Environment string
}

// Configuracion variable global para acceder a la configuración
var Configuracion *Config

// Variables_entorno carga las variables de entorno
func Variables_entorno() {
	// Cargar archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Advertencia: No se pudo cargar el archivo .env, usando variables del sistema")
	}

	// Crear instancia de configuración
	Configuracion = &Config{
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBPort:      getEnv("DB_PORT", "3306"),
		DBUser:      getEnv("DB_USER", "root"),
		DBPassword:  getEnv("DB_PASSWORD", ""),
		DBName:      getEnv("DB_NAME", "CRUD_GO"),
		ServerPort:  getEnv("SERVER_PORT", "8080"),
		Environment: getEnv("ENVIRONMENT", "development"),
	}

	log.Println("✅ Configuración cargada exitosamente")
}

// getEnv obtiene una variable de entorno o retorna un valor por defecto
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// GetDSN retorna el Data Source Name para MySQL
func (c *Config) GetDSN() string {
	// Formato DSN para MySQL: user:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)
}

// GetServerAddress retorna la dirección del servidor
func (c *Config) GetServerAddress() string {
	return ":" + c.ServerPort
}
