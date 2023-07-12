package consumption

import "errors"

var (
	// ErrNilConsumptionRepository is returned when the consumption repository is nil.
	ErrNilConsumptionRepository = errors.New("consumption repository is nil")
)
