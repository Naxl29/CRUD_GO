package models

import (
	"time"

	"gorm.io/gorm"
)

type Cliente struct {
	gorm.Model
	IDPersona     uint      `gorm:"not null" json:"id_persona"`
	TipoCuenta    string    `gorm:"type:varchar(20);not null" json:"tipo_cuenta"`
	Saldo         float64   `gorm:"type:decimal(15,2);default:0" json:"saldo"`
	FechaRegistro time.Time `gorm:"autoCreateTime" json:"fecha_registro"`

	// Relación con Persona
	Persona Persona `gorm:"foreignKey:IDPersona" json:"persona,omitempty"`
}
