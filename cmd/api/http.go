package main

import (
	"log"
	"net/http"

	"github.com/braejan/go-bia-challenge/internal/domain/consumption/entity"
	consumptionRepo "github.com/braejan/go-bia-challenge/internal/domain/consumption/repository/postgres"
	"github.com/braejan/go-bia-challenge/internal/domain/consumption/usecases"
	consumptionVO "github.com/braejan/go-bia-challenge/internal/valueobject/consumption"
	"github.com/braejan/go-bia-challenge/internal/valueobject/postgres"
	"github.com/gin-gonic/gin"
)

var consumptionUC usecases.ConsumptionUsecases

type ucFunction func(ids string, beginDate, endDate string) (consumptions entity.ConsumptionResponse, err error)

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
	router.GET("/consumption", func(c *gin.Context) {
		meterIDs := c.Query("meter_ids")
		beginDate := c.Query("start_date")
		endDate := c.Query("end_date")
		kindPeriod := c.Query("kind_period")
		ucFinalFunction, err := getUCFunction(kindPeriod)
		if err != nil {
			log.Println("error: ", err)
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		consumptions, err := ucFinalFunction(meterIDs, beginDate, endDate)
		if err != nil {
			log.Println("error: ", err)
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, consumptions)
	})
	router.Run(":8080")

}

func getUCFunction(kindPeriod string) (function ucFunction, err error) {
	switch kindPeriod {
	case "daily":
		function = consumptionUC.GetDailyAccumulatedConsumptionByIDs
	case "weekly":
		function = consumptionUC.GetWeeklyAccumulatedConsumptionByIDs
	case "monthly":
		function = consumptionUC.GetMonthlyAccumulatedConsumptionByIDs
	default:
		err = consumptionVO.ErrInvalidKindPeriod
	}
	return
}
