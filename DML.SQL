-- Menyisipkan data pelanggan
INSERT INTO customers (name, phone_number) VALUES
    ('bam', '1234567890'),
    ('bim', '9876543210'),
    ('boom', '5555555555');

-- Menyisipkan data karyawan
INSERT INTO employee (name, phone_number) VALUES
    ('hus', '1111111111'),
    ('has', '2222222222'),
    ('hes', '3333333333');

-- Menyisipkan data layanan laundry
INSERT INTO laundry_services (service_name, price) VALUES
    ('Cuci Biasa', 10000.00),
    ('Pencucian Kering', 15000.00),
    ('Setrika', 50000.00);

-- Menyisipkan data transaksi laundry
INSERT INTO transactions (customer_id, employee_id, service_id, quantity, total_price, transaction_date, finish_date) VALUES
    (1, 1, 1, 2, 20000.00, '2024-04-01', '2024-04-02'),
    (2, 2, 2, 1, 15000.00, '2024-04-02', '2024-04-03'),
    (3, 3, 3, 3, 15000.00, '2024-04-03', '2024-04-04');
