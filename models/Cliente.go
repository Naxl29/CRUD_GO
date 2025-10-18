package models

import "time"

type Cliente struct {
	ID            int
	IDPersona     int
	TipoCuenta    string
	Saldo         float64
	FechaRegistro time.Time
}
