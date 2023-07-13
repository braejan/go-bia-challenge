package main

import (
	"net/http"

	consumptionRepo "github.com/braejan/go-bia-challenge/internal/domain/consumption/repository/postgres"
	"github.com/braejan/go-bia-challenge/internal/domain/consumption/usecases"
	"github.com/braejan/go-bia-challenge/internal/valueobject/postgres"
	"github.com/gin-gonic/gin"
)

var consumptionUC usecases.ConsumptionUsecases

// init function called to initialize the usecases
func init() {
	// Create a postgres configuration from environment variables
	postgresConfig := postgres.NewPostgresConfigurationFromEnv()
	// Create a db Based on the configuration
	postgresDatabase := postgres.NewBasePostgresDatabase(postgresConfig)
	// Create a consumption repository based on the db
	consumptionRepo := consumptionRepo.NewConsumptionPostgresRepository(postgresDatabase)
	// Create a consumption usecases based on the consumption repository
	uc, err := usecases.NewConsumptionUsecases(consumptionRepo)
	if err != nil {
		panic(err)
	}
	consumptionUC = uc
}

func main() {
	router := gin.Default()

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to an url matching:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/consumption", func(c *gin.Context) {
		meterIDs := c.Query("meter_ids")
		beginDate := "2023-06-01"
		endDate := "2023-07-10"
		// Get daily accumulated consumption
		consumptions, err := consumptionUC.GetDailyAccumulatedConsumptionByIDs(meterIDs, beginDate, endDate)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, consumptions)
	})
	router.Run(":8080")

}
