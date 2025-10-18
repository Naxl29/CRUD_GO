package models

import "time"

type Transaccion struct {
	ID               int
	IDCliente        int
	Tipo             string
	Monto            float64
	Fecha            time.Time
	Descripcion      string
}