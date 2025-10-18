package models

import "gorm.io/gorm"

type Persona struct {
	gorm.Model
	PrimerNombre    string `gorm:"type:varchar(50);not null" json:"primer_nombre"`
	SegundoNombre   string `gorm:"type:varchar(50)" json:"segundo_nombre"`
	PrimerApellido  string `gorm:"type:varchar(50);not null" json:"primer_apellido"`
	SegundoApellido string `gorm:"type:varchar(50)" json:"segundo_apellido"`
	Documento       string `gorm:"type:varchar(20);unique;not null" json:"documento"`
	Correo          string `gorm:"type:varchar(100);unique;not null" json:"correo"`
	Direccion       string `gorm:"type:varchar(200);not null" json:"direccion"`
	Telefono        string `gorm:"type:varchar(15);not null" json:"telefono"`
}
