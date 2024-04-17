# Enigma Laundry Console App

Aplikasi konsol sederhana untuk mencatat transaksi di toko laundry Enigma.

## Struktur Database

Struktur database menggunakan PostgreSQL dan terdiri dari dua tabel master dan satu tabel transaksi:

1. Tabel `customers`:
    - customer_id (int, primary key)
    - name (text)
    - phone_number (text)

2. Tabel `laundry_services`:
    - service_id (int, primary key)
    - service_name (text)
    - price (numeric)

3. Tabel `transactions`:
    - transaction_id (int, primary key)
    - customer_id (int, foreign key ke customers)
    - employee_id (int, foreign key ke employee)
    - service_id (int, foreign key ke laundry_services)
    - quantity (int)
    - total_price (numeric)
    - transaction_date (date)
    - finish_date (date)

4. Tabel `employee`:
    - employee_id (int, primary key)
    - name (text)
    - phone_number (text)

## Cara Menggunakan Aplikasi

1. Pastikan Anda memiliki database PostgreSQL yang berjalan dan telah membuat struktur database sesuai dengan skema yang dijelaskan di atas.

2. Pastikan sudah menyesuaikan info database dengan local database yang dimiliki (seperti password, nama database, user)

3. Jalankan aplikasi dengan menjalankan file `main.go`.

4. Anda akan disambut dengan menu utama. Pilih salah satu opsi yang diinginkan.

5. Untuk setiap submenu (Customer, Transaction, Employee, Services), Anda dapat memilih opsi untuk menambah, memperbarui, menghapus, atau melihat data.

6. Pada menu Insert Customer atau Insert Transaction, pastikan untuk memasukkan input dengan benar, termasuk memperhatikan penggunaan spasi pada nama pelanggan.

7. Jika ingin keluar dari aplikasi, pilih opsi "0" di menu utama.

## Teknologi yang Digunakan

- Golang
- PostgreSQL