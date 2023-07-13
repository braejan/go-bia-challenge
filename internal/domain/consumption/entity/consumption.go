package entity

import (
	"github.com/google/uuid"
)

// Consumption is the consumption entity struct.
type Consumption struct {
	ID                 uuid.UUID
	MeterID            int64
	ActiveEnergy       float64
	ReactiveEnergy     float64
	CapacitiveReactive float64
	Solar              float64
	Date               string
}
