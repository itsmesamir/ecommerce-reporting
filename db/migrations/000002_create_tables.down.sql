-- Drop indexes first
DROP INDEX IF EXISTS idx_payment_date;
DROP INDEX IF EXISTS idx_order_date;

-- Drop tables in reverse order to avoid foreign key constraint issues
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS customers;

-- Drop enums last
DROP TYPE IF EXISTS payment_status;
DROP TYPE IF EXISTS order_status;
