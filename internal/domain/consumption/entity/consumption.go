package entity

import (
	"github.com/google/uuid"
)

/*
CREATE TABLE consumptions (
    id       UUID PRIMARY KEY,
    meter_id BIGINT,
    active_energy FLOAT,
    reactive_energy FLOAT,
    capacitive_reactive FLOAT,
    solar FLOAT,
    date TIMESTAMP NOT NULL DEFAULT NOW()
);
*/

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
