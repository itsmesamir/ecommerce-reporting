# Ecommerce Reporting Application

## Overview

This application is designed to generate various reports for an e-commerce platform.

## Setup Instructions

### 1. Clone the Repository

First, clone the repository to your local machine:

```bash
git clone https://github.com/itsmesamir/ecommerce-reporting.git
cd ecommerce-reporting
```

### 2. Configure Database Connection

Create a .env file and update the credentials accordingly. This usually involves setting your database username, password, host, and database name.

### 3. Run Database Migrations

```bash
# Navigate to the scripts directory
cd scripts

# Run the migrations (example command)
go run migrate.go up
```

### 4. Build the Application

```bash
# Navigate to the cmd directory
cd cmd

# Run the application

go run main.go
```

### 6. Execute Reports

To generate reports, you can use the following endpoints:

Sales Report:

```bash
GET /api/sales-report
```

This endpoint retrieves the sales report.

Customer Report:

```bash
GET /api/customer-report
```

This endpoint retrieves the customer report.

You can use a tool like Postman or curl to make requests to these endpoints. For example:

```bash
curl http://localhost:8080/api/sales-report
curl http://localhost:8080/api/customer-report
```

### 7. Rate Limiting

The application includes a Token Bucket rate limiter that allows 10 requests per minute to each endpoint. If the limit is exceeded, the application will return a 429 Too Many Requests error.

### 8. Caching with Redis

To improve performance, the application uses Redis for caching report data. When a report is generated, the result is cached in Redis. Subsequent requests for the same report will retrieve the data from the cache instead of querying the database, significantly reducing response time.

#### How to Use Redis Caching:

Ensure Redis is running on your machine or server.
Modify the report generation logic to store the results in Redis after fetching from the database.
Before fetching data from the database, check if the report is available in Redis.
