package core

type Charge interface {
	Apply()
	NewCharge()
}

type BaseCharge struct {
	BillingMetricId string
}
