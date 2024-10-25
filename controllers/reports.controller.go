package controller

import (
	service "ecommerce-reporting/services"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type ReportsController struct {
	reportsService service.ReportsService
}

func NewReportsController(reportsService service.ReportsService) *ReportsController {
	return &ReportsController{reportsService: reportsService}
}

func (c *ReportsController) GetSalesReport(w http.ResponseWriter, r *http.Request) {
	log.Println("GetSalesReport", r.URL.Query())
	startDateStr := r.URL.Query().Get("start_date")
	endDateStr := r.URL.Query().Get("end_date")
	productCategory := r.URL.Query().Get("product_category")
	productIDStr := r.URL.Query().Get("product_id")
	customerLocation := r.URL.Query().Get("customer_location")

	var startDate, endDate time.Time
	if startDateStr != "" {
		startDate, _ = time.Parse("2006-01-02", startDateStr)
	}
	if endDateStr != "" {
		endDate, _ = time.Parse("2006-01-02", endDateStr)
	}

	productID, _ := strconv.Atoi(productIDStr)

	reports, err := c.reportsService.GenerateSalesReport(r.Context(), startDate, endDate, &productCategory, &productID, &customerLocation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reports)
}

func (c *ReportsController) GetCustomerReport(w http.ResponseWriter, r *http.Request) {
	signupStartDateStr := r.URL.Query().Get("signup_start_date")
	signupEndDateStr := r.URL.Query().Get("signup_end_date")
	lifetimeValueStr := r.URL.Query().Get("lifetime_value")

	var signupStartDate, signupEndDate time.Time
	if signupStartDateStr != "" {
		signupStartDate, _ = time.Parse("2006-01-02", signupStartDateStr)
	}
	if signupEndDateStr != "" {
		signupEndDate, _ = time.Parse("2006-01-02", signupEndDateStr)
	}

	var minLifetimeValue *float64
	if lifetimeValueStr != "" {
		val, err := strconv.ParseFloat(lifetimeValueStr, 64)
		if err == nil {
			minLifetimeValue = &val
		}
	}

	reports, err := c.reportsService.GenerateCustomerReport(r.Context(), signupStartDate, signupEndDate, minLifetimeValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reports)
}
