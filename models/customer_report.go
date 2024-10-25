package model

type CustomerReport struct {
	TotalCustomers        int     `json:"total_customers"`
	CustomerSegment       string  `json:"customer_segment"`
	AverageOrderFrequency float64 `json:"average_order_frequency"`
}
