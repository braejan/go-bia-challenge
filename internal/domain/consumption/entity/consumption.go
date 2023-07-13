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

// NewConsumption creates a new instance of Consumption entity.
func NewConsumption(id uuid.UUID, meterID int64, activeEnergy, reactiveEnergy, capacitiveReactive, solar float64, date string) *Consumption {
	return &Consumption{
		ID:                 id,
		MeterID:            meterID,
		ActiveEnergy:       activeEnergy,
		ReactiveEnergy:     reactiveEnergy,
		CapacitiveReactive: capacitiveReactive,
		Solar:              solar,
		Date:               date,
	}
}
