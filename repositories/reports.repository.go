package repository

import (
	"context"
	"database/sql"
	model "ecommerce-reporting/models"
	utils "ecommerce-reporting/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type ReportsRepository interface {
	GetSalesReport(ctx context.Context, startDate, endDate time.Time, productCategory *string, productID *int, customerLocation *string) (model.SalesReport, error)
	GetCustomerReport(ctx context.Context, signupStartDate, signupEndDate time.Time, minLifetimeValue *float64) ([]model.CustomerReport, error)
}

type reportsRepository struct {
	db *pgxpool.Pool
}

func NewReportsRepository(db *pgxpool.Pool) ReportsRepository {
	return &reportsRepository{db: db}
}

func dereferenceString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func dereferenceInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

func (r *reportsRepository) GetSalesReport(ctx context.Context, startDate, endDate time.Time, productCategory *string, productID *int, customerLocation *string) (model.SalesReport, error) {
	var report model.SalesReport

	cacheKey := fmt.Sprintf("sales_report:%v:%v:%v:%v:%v", startDate, endDate, dereferenceString(productCategory), dereferenceInt(productID), dereferenceString(customerLocation))

	err := utils.GetCache(cacheKey, &report)

	// If the cache hit, return the cached report
	if err == nil {
		log.Println("Cache hit, returning cached report")
		return report, nil
	}

	fmt.Println("Cache miss, querying database")

	// Helper function to inject filters into a base query
	injectFilters := func(baseQuery string, args []interface{}) (string, []interface{}) {
		if !startDate.IsZero() {
			baseQuery += " AND o.order_date >= $" + strconv.Itoa(len(args)+1)
			args = append(args, startDate)
		}
		if !endDate.IsZero() {
			baseQuery += " AND o.order_date <= $" + strconv.Itoa(len(args)+1)
			args = append(args, endDate)
		}
		if productCategory != nil && *productCategory != "" {
			baseQuery += " AND p.category = $" + strconv.Itoa(len(args)+1)
			args = append(args, *productCategory)
		}
		if productID != nil && *productID > 0 {
			baseQuery += " AND oi.product_id = $" + strconv.Itoa(len(args)+1)
			args = append(args, *productID)
		}
		if customerLocation != nil && *customerLocation != "" {
			baseQuery += " AND c.location = $" + strconv.Itoa(len(args)+1)
			args = append(args, *customerLocation)
		}
		return baseQuery, args
	}

	// Helper function to execute a query and scan the result into a destination pointer
	executeQuery := func(query string, args []interface{}, dest **float64) error {
		var result sql.NullFloat64
		err := r.db.QueryRow(ctx, query, args...).Scan(&result)
		if err != nil {
			return err
		}
		if result.Valid {
			*dest = new(float64)
			**dest = result.Float64
		} else {
			*dest = nil
		}
		return nil
	}

	// List of queries for total sales, average order value, and number of products
	queries := []struct {
		query string
		dest  **float64
	}{
		{
			`SELECT SUM(oi.price) AS total_sales
             FROM orders o
             INNER JOIN order_items oi ON o.id = oi.order_id
             INNER JOIN products p ON oi.product_id = p.id
             INNER JOIN customers c ON o.customer_id = c.id
             WHERE o.status = 'COMPLETED'`,
			&report.TotalSales,
		},
		{
			`SELECT AVG(oi.price) AS avg_order_value
             FROM orders o
             INNER JOIN order_items oi ON o.id = oi.order_id
             INNER JOIN products p ON oi.product_id = p.id
             INNER JOIN customers c ON o.customer_id = c.id
             WHERE o.status = 'COMPLETED'`,
			&report.AvgOrderValue,
		},
		{
			`SELECT SUM(oi.quantity) AS number_of_products
             FROM orders o
             INNER JOIN order_items oi ON o.id = oi.order_id
             INNER JOIN products p ON oi.product_id = p.id
             INNER JOIN customers c ON o.customer_id = c.id
             WHERE o.status = 'COMPLETED'`,
			&report.NumberOfProducts,
		},
	}

	// Execute each query
	for _, q := range queries {
		args := []interface{}{}
		filteredQuery, filteredArgs := injectFilters(q.query, args)
		if err := executeQuery(filteredQuery, filteredArgs, q.dest); err != nil {
			return report, err
		}
	}

	// Query for total revenue by customer
	revenueQuery := `
        SELECT c.id AS customer_id, 
               c.name AS customer_name, 
               p.name AS product_name, 
               c.location AS region, 
               COALESCE(SUM(oi.price), 0) AS total_revenue
        FROM orders o
        INNER JOIN order_items oi ON o.id = oi.order_id
        INNER JOIN customers c ON o.customer_id = c.id
        INNER JOIN products p ON oi.product_id = p.id
        WHERE o.status = 'COMPLETED'`
	revenueQuery, revenueArgs := injectFilters(revenueQuery, []interface{}{})
	revenueQuery += ` GROUP BY c.id, c.name, p.name, c.location`

	rows, err := r.db.Query(ctx, revenueQuery, revenueArgs...)
	if err != nil {
		return report, err
	}
	defer rows.Close()

	for rows.Next() {
		var detail model.CustomerRevenueDetail
		if err := rows.Scan(&detail.CustomerID, &detail.CustomerName, &detail.ProductName, &detail.Region, &detail.TotalRevenue); err != nil {
			return report, err
		}
		report.TotalRevenue = append(report.TotalRevenue, detail)
	}

	// Set the cache with a 5-minute expiration
	if err := utils.SetCache(cacheKey, report, 5*time.Minute); err != nil {
		return report, err
	}

	return report, nil
}

func (r *reportsRepository) GetCustomerReport(ctx context.Context, signupStartDate, signupEndDate time.Time, minLifetimeValue *float64) ([]model.CustomerReport, error) {
	var whereClauses []string
	var args []interface{}
	argIndex := 1 // Start index for query arguments

	if !signupStartDate.IsZero() {
		whereClauses = append(whereClauses, fmt.Sprintf("c.signup_date >= $%d", argIndex))
		args = append(args, signupStartDate)
		argIndex++
	}

	if !signupEndDate.IsZero() {
		whereClauses = append(whereClauses, fmt.Sprintf("c.signup_date <= $%d", argIndex))
		args = append(args, signupEndDate)
		argIndex++
	}

	if minLifetimeValue != nil && *minLifetimeValue > 0 {
		whereClauses = append(whereClauses, fmt.Sprintf("c.lifetime_value = $%d", argIndex))
		args = append(args, *minLifetimeValue)
		argIndex++
	}

	query := `
        WITH customerFrequency AS (
            SELECT c.id, COUNT(o.id) AS order_frequency
            FROM customers c
            LEFT JOIN orders o ON o.customer_id = c.id
            %s
            GROUP BY c.id
        )

        SELECT
            COUNT(c.id) AS total_customers,
            CASE
                WHEN c.lifetime_value < 500 THEN 'Low valued customer'
                WHEN c.lifetime_value BETWEEN 500 AND 1000 THEN 'Medium valued customer'
                WHEN c.lifetime_value > 1000 THEN 'High valued customer'
            END AS customer_segment,
            AVG(cf.order_frequency) AS average_order_frequency
        FROM customers c
        INNER JOIN customerFrequency cf ON c.id = cf.id
        %s
        GROUP BY customer_segment;
    `

	var whereClause string
	if len(whereClauses) > 0 {
		whereClause = "WHERE " + strings.Join(whereClauses, " AND ")
	}

	query = fmt.Sprintf(query, whereClause, whereClause)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []model.CustomerReport
	for rows.Next() {
		var report model.CustomerReport
		if err := rows.Scan(&report.TotalCustomers, &report.CustomerSegment, &report.AverageOrderFrequency); err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}
	return reports, nil
}
