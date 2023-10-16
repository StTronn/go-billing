package core_test

// import (
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/sttronn/usage-billing/core"
// )

// func TestBillingMetricAggregateSum(t *testing.T) {
// 	events := []core.Event{
// 		{
// 			CustomerId: "customer1",
// 			Code:       "txn_event",
// 			TimeStamp:  time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
// 			Properties: map[string]interface{}{
// 				"txn": 1.0,
// 			},
// 		},
// 		{
// 			CustomerId: "customer1",
// 			Code:       "txn_event",
// 			TimeStamp:  time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
// 			Properties: map[string]interface{}{
// 				"txn": 20.0,
// 			},
// 		},
// 	}

// 	billingMetric := core.BillingMetric{
// 		Code:        "txn_event",
// 		FieldName:   "txn",
// 		Aggregation: core.SUM,
// 	}

// 	result, err := billingMetric.Aggregate(events)
// 	assert.NoError(t, err)
// 	assert.Equal(t, 21.0, result)

// 	billingMetric.Code = "txn_event"

// }

// func TestBillingMetricAggregateCount(t *testing.T) {
// 	events := []core.Event{
// 		{},
// 		{},
// 		{},
// 	}

// 	billingMetric := core.BillingMetric{
// 		Aggregation: core.COUNT,
// 	}

// 	result, err := billingMetric.Aggregate(events)
// 	assert.NoError(t, err)
// 	assert.Equal(t, 3.0, result)
// }

// func TestBillingMetricAggregateInvalidAggregation(t *testing.T) {
// 	events := []core.Event{}

// 	billingMetric := core.BillingMetric{
// 		Aggregation: 2,
// 	}

// 	_, err := billingMetric.Aggregate(events)
// 	assert.Error(t, err)
// 	assert.EqualError(t, err, "invalid aggregation type: 2")
// }

// func TestBillingMetricAggregateInvalidValueType(t *testing.T) {
// 	events := []core.Event{
// 		{
// 			Properties: map[string]interface{}{
// 				"foo": "not a float",
// 			},
// 		},
// 	}

// 	billingMetric := core.BillingMetric{
// 		Code:        "foo",
// 		Aggregation: core.SUM,
// 	}

// 	_, err := billingMetric.Aggregate(events)
// 	assert.Error(t, err)
// 	assert.EqualError(t, err, "invalid value type for foo")
// }
