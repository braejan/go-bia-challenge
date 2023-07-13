package consumption

import "errors"

var (
	// ErrNilConsumptionRepository is returned when the consumption repository is nil.
	ErrNilConsumptionRepository = errors.New("consumption repository is nil")
	// ErrInvalidMeterID is returned when the meter id list is invalid.
	ErrInvalidMeterID = errors.New("invalid meter ids")
)
