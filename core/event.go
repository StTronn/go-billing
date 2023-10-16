package core

import (
	"context"
	"database/sql"

	"github.com/sttronn/usage-billing/db"

	_ "github.com/go-sql-driver/mysql"
)

type Event db.Event

type AggregatedEvent struct {
	Value float64
}

func createQueries() (*db.Queries, error) {
	db_driver, err := sql.Open("mysql", "root:password@/go_billing?parseTime=true")
	if err != nil {
		return nil, err
	}
	return db.New(db_driver), nil
}

type AggregationStratergy interface {
	AggregateStratergy(plan Plan, queries *db.Queries) (AggregatedEvent, error)
}

type SumAggregation struct {
}

func (aggregation SumAggregation) AggregateStratergy(plan Plan, queries *db.Queries) (AggregatedEvent, error) {
	ctx := context.Background()
	sumByCustomerIdAndCodeNameParams := db.SumEventsByCustomerIdAndCodeNameParams{
		CustomerID: plan.CustomerId,
		Code:       plan.Code,
	}

	result, err := queries.SumEventsByCustomerIdAndCodeName(ctx, sumByCustomerIdAndCodeNameParams)
	if err != nil {
		return AggregatedEvent{}, err
	}

	return AggregatedEvent{Value: result.Sum}, err
}

type CountAggregation struct {
}

func (aggregation CountAggregation) AggregateStratergy(plan Plan, queries *db.Queries) (AggregatedEvent, error) {
	ctx := context.Background()

	countByCustomerIdAndCodeParam := db.CountEventsByCustomerIdAndCodeParams{
		CustomerID: plan.CustomerId,
		Code:       plan.Code,
	}
	result, err := queries.CountEventsByCustomerIdAndCode(ctx, countByCustomerIdAndCodeParam)

	if err != nil {
		return AggregatedEvent{}, err
	}
	return AggregatedEvent{Value: float64(result.Count)}, err

}

// TODO: this changes when charges comes into play
func Aggregate(plan Plan, queries *db.Queries) (AggregatedEvent, error) {
	cases := map[Aggregation]AggregationStratergy{
		SUM:   SumAggregation{},
		COUNT: CountAggregation{},
	}

	aggregator, ok := cases[plan.Aggregation]
	if !ok {
		return AggregatedEvent{}, nil
	}

	// Call the aggregator's aggregate method to get the result
	result, err := aggregator.AggregateStratergy(plan, queries)

	if err != nil {
		return AggregatedEvent{}, err
	}

	return result, nil

}

func (event Event) Create(queries *db.Queries) (int64, error) {
	ctx := context.Background()

	result, err := queries.CreateEvent(ctx, db.CreateEventParams{
		CustomerID: event.CustomerID,
		Code:       event.Code,
		Timestamp:  event.Timestamp,
		Value:      event.Value,
	})
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()

}
