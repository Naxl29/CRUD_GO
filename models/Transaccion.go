package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaccion struct {
	gorm.Model
	IDCliente   uint      `gorm:"not null" json:"id_cliente"`
	Tipo        string    `gorm:"type:varchar(20);not null" json:"tipo"`
	Monto       float64   `gorm:"type:decimal(15,2);not null" json:"monto"`
	Fecha       time.Time `gorm:"autoCreateTime" json:"fecha"`
	Descripcion string    `gorm:"type:varchar(255)" json:"descripcion"`

	// Relación con Cliente
	Cliente Cliente `gorm:"foreignKey:IDCliente" json:"cliente,omitempty"`
}
