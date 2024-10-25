package main

import (
	"log"
	"net/http"

	controller "ecommerce-reporting/controllers"
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
	api.HandleFunc("/sales-report", ctrl.GetSalesReport).Methods("GET")
	api.HandleFunc("/customer-report", ctrl.GetCustomerReport).Methods("GET")

	// Start the server
	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
