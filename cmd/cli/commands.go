package main

import (
	"encoding/json"

	consumptionRepo "github.com/braejan/go-bia-challenge/internal/domain/consumption/repository/postgres"
	consumptionUc "github.com/braejan/go-bia-challenge/internal/domain/consumption/usecases"
	"github.com/braejan/go-bia-challenge/internal/valueobject/postgres"
)

func main() {
	metersID := []int64{1, 2}
	beginDate := "2023-06-01"
	endDate := "2023-07-10"
	// Create a postgres configuration from environment variables
	postgresConfig := postgres.NewPostgresConfigurationFromEnv()
	// Create a db Based on the configuration
	postgresDatabase := postgres.NewBasePostgresDatabase(postgresConfig)
	// Create a consumption repository based on the db
	consumptionRepo := consumptionRepo.NewConsumptionPostgresRepository(postgresDatabase)
	// Create a consumption usecases based on the consumption repository
	consumptionUsecases, err := consumptionUc.NewConsumptionUsecases(consumptionRepo)
	if err != nil {
		panic(err)
	}
	// Get daily accumulated consumption
	consumptions, err := consumptionUsecases.GetDailyAccumulatedConsumptionByIDs(metersID, beginDate, endDate)
	if err != nil {
		panic(err)
	}
	// print consumptions JSON response
	jsonString, err := json.Marshal(consumptions)
	if err != nil {
		panic(err)
	}
	println(string(jsonString))
}
