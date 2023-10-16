// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type BillingMetricAggregation string

const (
	BillingMetricAggregationSUM   BillingMetricAggregation = "SUM"
	BillingMetricAggregationCOUNT BillingMetricAggregation = "COUNT"
)

func (e *BillingMetricAggregation) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = BillingMetricAggregation(s)
	case string:
		*e = BillingMetricAggregation(s)
	default:
		return fmt.Errorf("unsupported scan type for BillingMetricAggregation: %T", src)
	}
	return nil
}

type NullBillingMetricAggregation struct {
	BillingMetricAggregation BillingMetricAggregation
	Valid                    bool // Valid is true if BillingMetricAggregation is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullBillingMetricAggregation) Scan(value interface{}) error {
	if value == nil {
		ns.BillingMetricAggregation, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.BillingMetricAggregation.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullBillingMetricAggregation) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.BillingMetricAggregation), nil
}

type Author struct {
	ID   int64
	Name string
	Bio  sql.NullString
}

type BillingMetric struct {
	ID          int32
	Name        string
	Code        string
	Aggregation BillingMetricAggregation
	FieldName   string
}

type Event struct {
	ID         int32
	CustomerID string
	Code       string
	Timestamp  time.Time
	Value      float64
}