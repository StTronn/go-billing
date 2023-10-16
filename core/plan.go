package core

type Interval int

const (
	Daily Interval = iota
	Monthly
	Yearly
)

type Plan struct {
	CustomerId string
	*BillingMetric
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
