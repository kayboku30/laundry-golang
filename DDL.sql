-- Script DDL PostgreSQL

-- Tabel untuk pelanggan
CREATE TABLE IF NOT EXISTS customers (
    customer_id SERIAl PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    phone_number VARCHAR(15) NOT NULL UNIQUE
);

-- tabel untuk employee
CREATE TABLE IF NOT EXISTS employee (
    employee_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    phone_number VARCHAR(15) NOT NULL UNIQUE
);

-- Tabel untuk layanan laundry
CREATE TABLE IF NOT EXISTS laundry_services (
    service_id SERIAL PRIMARY KEY,
    service_name VARCHAR(100) NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);

-- Tabel transaksi laundry
CREATE TABLE IF NOT EXISTS transactions (
    transaction_id SERIAL PRIMARY KEY,
    customer_id INT REFERENCES customers(customer_id),
	employee_id INT REFERENCES employee(employee_id),
    service_id INT REFERENCES laundry_services(service_id),
    quantity INT NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    transaction_date DATE NOT NULL,
	finish_date DATE NOT NULL
);
