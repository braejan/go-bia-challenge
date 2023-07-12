package entity

import (
	"github.com/google/uuid"
)

// Consumption is the consumption entity struct.
type Consumption struct {
	ID           uuid.UUID
	MeterID      int64
	Consumption  float64
	CreationDate string
}

// NewConsumption creates a new instance of Consumption entity.
func NewConsumption(id uuid.UUID, meterID int64, consumption float64, creationDate string) *Consumption {
	return &Consumption{
		ID:           id,
		MeterID:      meterID,
		Consumption:  consumption,
		CreationDate: creationDate,
	}
}
