package core

type Aggregation int

const (
	SUM Aggregation = iota
	COUNT
)

type BillingMetric struct {
	Name        string
	Code        string
	Aggregation Aggregation
	FieldName   string
}

type Sum struct {
}

// func (s Sum) aggregate(events []Event, billingMetric BillingMetric) (float64, error) {

// 	filteredEvents := []Event{}
// 	for _, event := range events {
// 		if code, ok := event.Properties["Code"].(string); ok && code == billingMetric.Code {
// 			filteredEvents = append(filteredEvents, event)
// 		}
// 	}

// 	aggr := 0.0
// 	//this will be sql query when using db
// 	for _, event := range events {
// 		//convert to float
// 		value, ok := event.Properties[billingMetric.FieldName].(float64)
// 		if !ok {
// 			return 0, fmt.Errorf("invalid value type for %s", billingMetric.Code)
// 		}

// 		//aggregate based on sum
// 		aggr += value
// 	}
// 	return aggr, nil
// }

// type Count struct {
// }

// func (c Count) aggregate(events []Event, billingMetric BillingMetric) (float64, error) {
// 	return float64(len(events)), nil
// }

// func (billMetric BillingMetric) Aggregate(events []Event) (float64, error) {
// 	cases := map[Aggregation]Aggregate{
// 		SUM:   Sum{},
// 		COUNT: Count{},
// 	}

// 	aggregator, ok := cases[billMetric.Aggregation]
// 	if !ok {
// 		return 0, fmt.Errorf("invalid aggregation type: %v", billMetric.Aggregation)
// 	}

// 	// Call the aggregator's aggregate method to get the result
// 	result, err := aggregator.aggregate(events, billMetric)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return result, nil

// }
