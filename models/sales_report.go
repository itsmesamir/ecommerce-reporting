package model

type SalesReport struct {
	TotalSales       *float64                `json:"total_sales"`
	AvgOrderValue    *float64                `json:"avg_order_value"`
	NumberOfProducts *float64                `json:"number_of_products"`
	TotalRevenue     []CustomerRevenueDetail `json:"total_revenue"`
}

type CustomerRevenueDetail struct {
	CustomerID   int     `json:"customer_id"`
	CustomerName string  `json:"customer_name"`
	ProductName  string  `json:"product_name"`
	Region       string  `json:"region"`
	TotalRevenue float64 `json:"total_revenue"`
}
