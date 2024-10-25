package service

import (
	"context"
	model "ecommerce-reporting/models"
	repository "ecommerce-reporting/repositories"
	"time"
)

type ReportsService interface {
	GenerateSalesReport(ctx context.Context, startDate, endDate time.Time, productCategory *string, productID *int, customerLocation *string) (model.SalesReport, error)
	GenerateCustomerReport(ctx context.Context, signupStartDate, signupEndDate time.Time, minLifetimeValue *float64) ([]model.CustomerReport, error)
}

type reportsService struct {
	repo repository.ReportsRepository
}

func NewReportsService(repo repository.ReportsRepository) ReportsService {
	return &reportsService{repo: repo}
}

func (s *reportsService) GenerateSalesReport(ctx context.Context, startDate, endDate time.Time, productCategory *string, productID *int, customerLocation *string) (model.SalesReport, error) {
	return s.repo.GetSalesReport(ctx, startDate, endDate, productCategory, productID, customerLocation)
}

func (s *reportsService) GenerateCustomerReport(ctx context.Context, signupStartDate, signupEndDate time.Time, minLifetimeValue *float64) ([]model.CustomerReport, error) {
	return s.repo.GetCustomerReport(ctx, signupStartDate, signupEndDate, minLifetimeValue)
}
