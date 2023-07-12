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

// NewConsumptionResponse creates a new instance of ConsumptionResponse.
func NewConsumptionResponse(period []string, dataGraph []*ConsumptionDataGraph) *ConsumptionResponse {
	return &ConsumptionResponse{
		Period:    period,
		DataGraph: dataGraph,
	}
}

// NewConsumptionDataGraph creates a new instance of ConsumptionDataGraph.
func NewConsumptionDataGraph(
	meterID int64,
	address string,
	activeEnergy,
	reactiveEnergy,
	capacitiveReactive,
	solar []float64) *ConsumptionDataGraph {
	return &ConsumptionDataGraph{
		MeterID:            meterID,
		Address:            address,
		ActiveEnergy:       activeEnergy,
		ReactiveEnergy:     reactiveEnergy,
		CapacitiveReactive: capacitiveReactive,
		Solar:              solar,
	}
}
