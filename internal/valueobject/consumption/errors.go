package consumption

import "errors"

var (
	// ErrNilConsumptionRepository is returned when the consumption repository is nil.
	ErrNilConsumptionRepository = errors.New("consumption repository is nil")
	// ErrInvalidMeterID is returned when the meter id list is invalid.
	ErrInvalidMeterID = errors.New("invalid meter ids")
	// ErrEmptyDate is returned when the date is empty.
	ErrEmptyDate = errors.New("empty date")
	// ErrInvalidDate is returned when the date is invalid.
	ErrInvalidDate = errors.New("invalid date")
	// ErrInvalidKindPeriod is returned when the kind period is invalid.
	ErrInvalidKindPeriod = errors.New("invalid kind period")
)
