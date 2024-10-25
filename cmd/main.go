package main

import (
	"log"
	"net/http"
	"time"

	"ecommerce-reporting/constants"
	controller "ecommerce-reporting/controllers"
	"ecommerce-reporting/middleware" // Import your middleware package
	repository "ecommerce-reporting/repositories"
	service "ecommerce-reporting/services"
	utils "ecommerce-reporting/utils"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize database connection
	utils.ConnectDB()
	defer utils.CloseDB()

	// Initialize dependencies
	repo := repository.NewReportsRepository(utils.DB)
	svc := service.NewReportsService(repo)
	ctrl := controller.NewReportsController(svc)

	// Set up router
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	// Create a rate limiter with a limit given by MaxRequests and a time window of 1 minute
	rl := middleware.NewRateLimiter(constants.MaxRequests,
		time.Minute)

	api.HandleFunc("/sales-report", func(w http.ResponseWriter, r *http.Request) {
		rl.ServeHTTP(w, r, ctrl.GetSalesReport)
	}).Methods("GET")

	api.HandleFunc("/customer-report", func(w http.ResponseWriter, r *http.Request) {
		rl.ServeHTTP(w, r, ctrl.GetCustomerReport)
	}).Methods("GET")

	// Start the server
	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
