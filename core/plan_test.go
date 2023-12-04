package core_test

import (
	"testing"

	"database/sql"

	"github.com/stretchr/testify/assert"
	"github.com/sttronn/usage-billing/core"
	"github.com/sttronn/usage-billing/db"
)

func TestBilling(t *testing.T) {

	db_driver, err := sql.Open("mysql", "root:password@/go_billing_test?parseTime=true")
	if err != nil {
		t.Fatal(err)
	}
	defer db_driver.Close()
	// Create a mock database connection

	queries := db.New(db_driver)

	// Create a test plan
	plan := core.Plan{
		CustomerId: "customer1",
		BillingMetric: &core.BillingMetric{
			Name:        "Test Metric",
			Code:        "txn_event",
			Aggregation: core.SUM,
			FieldName:   "value",
		},
		Charges: core.Charges{
			Model: core.TIERS,
			Properties: core.ChargesProperties{
				Slabs: []core.Slab{
					{FromValue: 0, ToValue: 100, FlatFee: 10, PerUnitFee: 0.1},
					{FromValue: 100, ToValue: 1000, FlatFee: 50, PerUnitFee: 0.05},
					{FromValue: 1000, ToValue: 10000, FlatFee: 200, PerUnitFee: 0.01},
				},
				FlatFee:       0,
				PercentageFee: 0,
			},
		},
	}

	// Create a mock aggregated event
	//event := core.AggregatedEvent{Value: 500}

	// Call the Billing function
	result, err := core.Billing(plan, queries)

	// Check the result
	// expected := 60.0
	// if result != expected {
	// 	t.Errorf("Billing(%v, %v) = %v, expected %v", plan, db, result, expected)
	// }

	// // Check the error
	// if err != nil {
	// 	t.Errorf("Billing(%v, %v) returned an unexpected error: %v", plan, db, err)
	// }

	assert.NotNil(t, queries)
	assert.Equal(t, plan.BillingMetric.Name, "Test Metric")
	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Equal(t, result, 60.0)
}
