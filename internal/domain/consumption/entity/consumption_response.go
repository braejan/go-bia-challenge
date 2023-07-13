package entity

type ConsumptionResponse struct {
	Period    []string                `json:"period"`
	DataGraph []*ConsumptionDataGraph `json:"data_graph"`
}

type ConsumptionDataGraph struct {
	MeterID            int64     `json:"meter_id"`
	Address            string    `json:"address"`
	ActiveEnergy       []float64 `json:"active_energy"`
	ReactiveEnergy     []float64 `json:"reactive_energy"`
	CapacitiveReactive []float64 `json:"capacitive_reactive"`
	Solar              []float64 `json:"solar"`
}
