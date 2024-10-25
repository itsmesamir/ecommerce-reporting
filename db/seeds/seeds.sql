INSERT INTO customers (name, email, signup_date, location, lifetime_value)
VALUES
    ('John Doe', 'john.doe@example.com', '2023-01-01', 'New York', 1500.00),
    ('Jane Smith', 'jane.smith@example.com', '2023-02-15', 'Los Angeles', 2000.00),
    ('Alice Johnson', 'alice.johnson@example.com', '2023-03-10', 'Chicago', 750.00),
    ('Bob Brown', 'bob.brown@example.com', '2023-04-25', 'Houston', 500.00),
    ('Charlie Green', 'charlie.green@example.com', '2023-03-18', 'Miami', 1200.00),
    ('David White', 'david.white@example.com', '2023-02-11', 'Denver', 1800.00),
    ('Ella Black', 'ella.black@example.com', '2023-01-25', 'Seattle', 950.00),
    ('Frank Harris', 'frank.harris@example.com', '2023-05-10', 'Phoenix', 500.00),
    ('Grace Clark', 'grace.clark@example.com', '2023-06-15', 'San Diego', 1700.00),
    ('Henry Lewis', 'henry.lewis@example.com', '2023-03-22', 'Dallas', 1400.00),
    ('Ivy Edwards', 'ivy.edwards@example.com', '2023-04-30', 'Austin', 1650.00),
    ('Jake Thompson', 'jake.thompson@example.com', '2023-05-15', 'San Francisco', 1200.00),
    ('Kate Adams', 'kate.adams@example.com', '2023-07-05', 'Las Vegas', 1300.00),
    ('Liam Lee', 'liam.lee@example.com', '2023-06-18', 'Boston', 900.00),
    ('Mia Evans', 'mia.evans@example.com', '2023-07-20', 'Orlando', 1450.00),
    ('Noah Turner', 'noah.turner@example.com', '2023-05-22', 'Philadelphia', 600.00),
    ('Olivia Hill', 'olivia.hill@example.com', '2023-08-10', 'Minneapolis', 1100.00),
    ('Paul Scott', 'paul.scott@example.com', '2023-07-25', 'Portland', 1250.00),
    ('Quinn Walker', 'quinn.walker@example.com', '2023-06-28', 'Charlotte', 1150.00),
    ('Ruby Allen', 'ruby.allen@example.com', '2023-05-30', 'Sacramento', 850.00);

INSERT INTO products (name, category, price)
VALUES
    ('Laptop', 'Electronics', 999.99),
    ('Headphones', 'Electronics', 199.99),
    ('Coffee Maker', 'Appliances', 79.99),
    ('Desk Chair', 'Furniture', 149.99),
    ('Smartphone', 'Electronics', 699.99),
    ('Blender', 'Appliances', 59.99),
    ('TV', 'Electronics', 399.99),
    ('Sofa', 'Furniture', 499.99),
    ('Gaming Console', 'Electronics', 299.99),
    ('Vacuum Cleaner', 'Appliances', 89.99),
    ('Microwave Oven', 'Appliances', 120.99),
    ('Refrigerator', 'Appliances', 1200.00),
    ('Office Desk', 'Furniture', 299.99),
    ('Smartwatch', 'Electronics', 199.99),
    ('Camera', 'Electronics', 499.99),
    ('Bookshelf', 'Furniture', 129.99),
    ('Air Purifier', 'Appliances', 229.99),
    ('Tablet', 'Electronics', 449.99),
    ('Washing Machine', 'Appliances', 599.99),
    ('Office Lamp', 'Furniture', 49.99);

INSERT INTO orders (customer_id, order_date, status)
VALUES
    (1, '2023-05-01', 'COMPLETED'),
    (2, '2023-06-01', 'PENDING'),
    (3, '2023-07-15', 'CANCELED'),
    (4, '2023-08-10', 'COMPLETED'),
    (5, '2023-08-15', 'COMPLETED'),
    (6, '2023-08-20', 'PENDING'),
    (7, '2023-09-01', 'CANCELED'),
    (8, '2023-09-10', 'PENDING'),
    (9, '2023-09-15', 'COMPLETED'),
    (10, '2023-09-20', 'COMPLETED'),
    (11, '2023-10-01', 'PENDING'),
    (12, '2023-10-05', 'COMPLETED'),
    (13, '2023-10-10', 'PENDING'),
    (14, '2023-10-15', 'CANCELED'),
    (15, '2023-10-18', 'COMPLETED'),
    (16, '2023-10-22', 'PENDING'),
    (17, '2023-10-25', 'PENDING'),
    (18, '2023-10-28', 'COMPLETED'),
    (19, '2023-11-01', 'PENDING'),
    (20, '2023-11-05', 'COMPLETED');

INSERT INTO order_items (order_id, product_id, quantity, price)
VALUES
    (1, 1, 1, 999.99),
    (1, 2, 2, 199.99),
    (2, 3, 1, 79.99),
    (3, 4, 1, 149.99),
    (4, 5, 1, 699.99),
    (5, 6, 1, 59.99),
    (6, 7, 1, 399.99),
    (7, 8, 1, 499.99),
    (8, 9, 1, 299.99),
    (9, 10, 1, 89.99),
    (10, 11, 1, 120.99),
    (11, 12, 1, 1200.00),
    (12, 13, 1, 299.99),
    (13, 14, 1, 199.99),
    (14, 15, 1, 499.99),
    (15, 16, 1, 129.99),
    (16, 17, 1, 229.99),
    (17, 18, 1, 449.99),
    (18, 19, 1, 599.99),
    (19, 20, 1, 49.99);

INSERT INTO transactions (order_id, payment_status, payment_date, total_amount)
VALUES
    (1, 'SUCCESS', '2023-05-02', 1399.97),
    (2, 'FAILED', '2023-06-02', 79.99),
    (3, 'SUCCESS', '2023-08-11', 149.99),
    (4, 'SUCCESS', '2023-08-12', 699.99),
    (5, 'SUCCESS', '2023-08-16', 59.99),
    (6, 'FAILED', '2023-08-21', 399.99),
    (7, 'FAILED', '2023-09-02', 499.99),
    (8, 'SUCCESS', '2023-09-11', 299.99),
    (9, 'SUCCESS', '2023-09-16', 89.99),
    (10, 'SUCCESS', '2023-09-21', 120.99),
    (11, 'SUCCESS', '2023-10-02', 1200.00),
    (12, 'SUCCESS', '2023-10-06', 299.99),
    (13, 'SUCCESS', '2023-10-11', 199.99),
    (14, 'FAILED', '2023-10-16', 499.99),
    (15, 'SUCCESS', '2023-10-19', 129.99),
    (16, 'SUCCESS', '2023-10-23', 229.99),
    (17, 'SUCCESS', '2023-10-26', 449.99),
    (18, 'SUCCESS', '2023-10-29', 599.99),
    (19, 'SUCCESS', '2023-11-02', 49.99),
    (20, 'SUCCESS', '2023-11-06', 49.99);
