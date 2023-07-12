package postgres

import (
	"context"

	"github.com/braejan/go-bia-challenge/internal/domain/consumption/entity"
	"github.com/braejan/go-bia-challenge/internal/valueobject/postgres"
	"github.com/lib/pq"
)

type ConsumptionPostgresRepository struct {
	baseDB postgres.PostgresDatabase
}

// NewConsumptionPostgresRepository creates a new instance of ConsumptionRepository interface implementation.
func NewConsumptionPostgresRepository(baseDB postgres.PostgresDatabase) *ConsumptionPostgresRepository {
	return &ConsumptionPostgresRepository{
		baseDB: baseDB,
	}
}

// ConsumptionRepository interface implementation.

// GetConsumptionsByIDSGroupedDaily returns a list of consumptions grouped by day.
const (
	getConsumptionsByIDSGroupedDailyQuery = `
SELECT
    DATE_TRUNC('week', date) AS week_start_date,
    meter_id,
    SUM(active_energy) AS total_active_energy,
    SUM(reactive_energy) AS total_reactive_energy,
    SUM(capacitive_reactive) AS total_capacitive_reactive,
    SUM(solar) AS total_solar
FROM
    consumptions
WHERE
    meter_id = ANY ($1)
    AND date >= $2
    AND date <= $3
GROUP BY
    week_start_date,
    meter_id
ORDER BY
    meter_id asc,
    week_start_date desc
`
)

func (postgresRepo *ConsumptionPostgresRepository) GetConsumptionsByIDSGroupedDaily(ids []int64, beginDate, endDate string) (consumptions []*entity.Consumption, err error) {
	db, err := postgresRepo.baseDB.Open()
	if err != nil {
		return nil, err
	}
	defer postgresRepo.baseDB.Close(db)
	idArray := pq.Array(ids)
	rows, err := db.QueryContext(context.Background(), getConsumptionsByIDSGroupedDailyQuery, idArray, beginDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var consumption entity.Consumption
		err = rows.Scan(
			&consumption.Date,
			&consumption.MeterID,
			&consumption.ActiveEnergy,
			&consumption.ReactiveEnergy,
			&consumption.CapacitiveReactive,
			&consumption.Solar,
		)
		if err != nil {
			return nil, err
		}
		consumptions = append(consumptions, &consumption)
	}
	return
}

// GetConsumptionsByIDSGroupedWeekly returns a list of consumptions grouped by week.
func (postgresRepo *ConsumptionPostgresRepository) GetConsumptionsByIDSGroupedWeekly(ids []int64, beginDate, endDate string) (consumptions []*entity.Consumption, err error) {
	return
}

// GetConsumptionsByIDSGroupedMonthly returns a list of consumptions grouped by month.
func (postgresRepo *ConsumptionPostgresRepository) GetConsumptionsByIDSGroupedMonthly(ids []int64, beginDate, endDate string) (consumptions []*entity.Consumption, err error) {
	return
}
