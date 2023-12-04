package core

import (
	"context"
	"log"

	"github.com/sttronn/usage-billing/db"

	_ "github.com/go-sql-driver/mysql"
)

type Interval int

const (
	Daily Interval = iota
	Monthly
	Yearly
)

type Aggregation int

const (
	SUM Aggregation = iota
	COUNT
)

// TODO: move things to types folder
type BillingMetric struct {
	Name        string
	Code        string
	Aggregation Aggregation
	FieldName   string
}

type Model int

const (
	STANDARD Model = iota
	TIERS
	PER_TRANSACTION
)

type Slab struct {
	ToValue    float64
	FromValue  float64
	FlatFee    float64
	PerUnitFee float64
}

type ChargesProperties struct {
	Slabs         []Slab
	FlatFee       float64
	PercentageFee float64
}

type Charges struct {
	Model      Model
	Properties ChargesProperties
}

type Plan struct {
	CustomerId string
	*BillingMetric
	Charges Charges
}

// type Charges {
// 	BillingMetricId string

// }

// type Plan struct {
// 	CustomerId string
// 	Active bool
// 	Interval Interval
// 	charges Charges

// }
// func (charges Charges) CalculateCharges() (float64, error) {

// }

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

type SumWithRangeAggregation struct {
}

func (aggregation SumWithRangeAggregation) AggregateStratergy(plan Plan, queries *db.Queries) (AggregatedEvent, error) {
	ctx := context.Background()
	sumByCustomerIdAndCodeNameParams := db.SumEventsByCustomerIdCodeNameStartRangeAndEndRangeParams{
		CustomerID: plan.CustomerId,
		Code:       plan.Code,
		Value:      0,
		Value_2:    200,
	}

	result, err := queries.SumEventsByCustomerIdCodeNameStartRangeAndEndRange(ctx, sumByCustomerIdAndCodeNameParams)
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

type AggregateInput struct {
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

// TOOD: instead of returning float64 return invoice object
func Billing(plan Plan, queries *db.Queries) (float64, error) {

	if plan.Charges.Model == STANDARD {
		aggregatedEvent, err := Aggregate(plan, queries)
		if err != nil {
			log.Printf("Error aggregating events: %v", err)
			return 0, err
		}
		return (aggregatedEvent.Value)*(plan.Charges.Properties.PercentageFee) + plan.Charges.Properties.FlatFee, nil
	}

	if plan.Charges.Model == TIERS {
		aggregatedEvent, err := Aggregate(plan, queries)
		if err != nil {
			log.Printf("Error aggregating events: %v", err)
			return 0, err
		}
		charges := 0.0
		//for each event that is make sure that there is combing

		for _, slab := range plan.Charges.Properties.Slabs {
			if slab.ToValue <= aggregatedEvent.Value {
				charges += (slab.FromValue-slab.ToValue)*slab.PerUnitFee + slab.FlatFee
			} else if slab.FromValue <= aggregatedEvent.Value {
				charges += (aggregatedEvent.Value - slab.FromValue) * slab.PerUnitFee * aggregatedEvent.Value
			}
		}
		return charges, nil
	}

	if plan.Charges.Model == PER_TRANSACTION {

	}

	return 0, nil

}
