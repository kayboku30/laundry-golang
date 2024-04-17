package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "cesena10"
	dbname   = "laundry_2"
)

func main() {
	// Buat koneksi ke database
	connectDb := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connectDb)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	// Menu aplikasi
	for {
		fmt.Println("\n=== Enigma Laundry ===")
		fmt.Println("1. Menu Customers")
		fmt.Println("2. Menu services")
		fmt.Println("3. Menu Employee")
		fmt.Println("4. Menu Transaction")
		fmt.Println("0. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			menuCustomer(db)
		case 2:
			menuService(db)
		case 3:
			menuEmployee(db)
		case 4:
			menuTransaction(db)
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func menuCustomer(db *sql.DB) {
	for {
		fmt.Println("\n=== Enigma Laundry ===")
		fmt.Println("1. Insert Customers")
		fmt.Println("2. Update Customers")
		fmt.Println("3. Delete Customers")
		fmt.Println("4. View Customers")
		fmt.Println("0. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			insertCustomer(db)
		case 2:
			updateCustomer(db)
		case 3:
			deleteCustomer(db)
		case 4:
			viewCustomer(db)
		case 0:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func menuService(db *sql.DB) {
	for {
		fmt.Println("\n=== Enigma Laundry ===")
		fmt.Println("1. Insert Services")
		fmt.Println("2. Update Services")
		fmt.Println("3. Delete Services")
		fmt.Println("4. View Services")
		fmt.Println("0. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			insertServices(db)
		case 2:
			updateServices(db)
		case 3:
			deleteServices(db)
		case 4:
			viewServices(db)
		case 0:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func menuTransaction(db *sql.DB) {
	for {
		fmt.Println("\n=== Enigma Laundry ===")
		fmt.Println("1. Insert Transaction")
		fmt.Println("2. Update Transaction")
		fmt.Println("3. Delete Transaction")
		fmt.Println("4. View Transaction")
		fmt.Println("0. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			insertTransaction(db)
		case 2:
			updateTransaction(db)
		case 3:
			deleteTransaction(db)
		case 4:
			viewTransactions(db)
		case 0:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func updateTransaction(db *sql.DB) {
	// Input ID Transaction yang akan diupdate
	var transactionID int
	var checkID int

	fmt.Print("Enter Transaction ID: ")
	fmt.Scanln(&transactionID)

	errs := db.QueryRow("SELECT COUNT(*) FROM transactions WHERE transaction_id = $1", transactionID).Scan(&checkID)
	if errs != nil {
		panic(errs)
	}
	if checkID == 0 {
		fmt.Println("\nInvalid Transaction ID.")
		fmt.Println("Press Enter to return to menu...")
		fmt.Scanln()
		return
	}

	// Input data transaction yang baru
	var customerID, employeeID, serviceID, quantity int
	var finishDate string
	fmt.Print("Enter Customer ID: ")
	fmt.Scanln(&customerID)
	fmt.Print("Enter Employee ID: ")
	fmt.Scanln(&employeeID)
	fmt.Print("Enter Service ID: ")
	fmt.Scanln(&serviceID)
	fmt.Print("Enter Quantity: ")
	fmt.Scanln(&quantity)
	fmt.Print("Enter Finish Date (YYYY-MM-DD): ")
	fmt.Scanln(&finishDate)

	// Memperoleh harga layanan dari database
	err := db.QueryRow("SELECT COUNT(*) FROM customers WHERE customer_id = $1", customerID).Scan(&checkID)
	if err != nil {
		panic(err)
	}
	if checkID == 0 {
		fmt.Println("Invalid Customer ID.")
		return
	}
	err = db.QueryRow("SELECT COUNT(*) FROM employee WHERE employee_id = $1", employeeID).Scan(&checkID)
	if err != nil {
		panic(err)
	}
	if checkID == 0 {
		fmt.Println("Invalid Employee ID.")
		return
	}

	err = db.QueryRow("SELECT COUNT(*) FROM laundry_services WHERE service_id = $1", serviceID).Scan(&checkID)
	if err != nil {
		panic(err)
	}
	if checkID == 0 {
		fmt.Println("Invalid Service ID.")
		return
	}

	// Memeriksa apakah jumlah yang dimasukkan valid
	if quantity <= 0 {
		fmt.Println("Quantity must be greater than 0.")
		return
	}
	var price float64
	errs = db.QueryRow("SELECT price FROM laundry_services WHERE service_id = $1", serviceID).Scan(&price)
	if errs != nil {
		panic(errs)
	}

	// Menghitung total harga transaksi
	totalPrice := price * float64(quantity)

	// Update data pegawai di database
	_, errf := db.Exec("UPDATE transactions SET customer_id = $1, employee_id = $2, service_id = $3, quantity = $4,total_price = $5, finish_date = $6 WHERE transaction_id = $7", customerID, employeeID, serviceID, quantity, totalPrice, finishDate, transactionID)
	if errf != nil {
		panic(errf)
	} else {
		fmt.Println("\ntransaction updated successfully.")
	}
	fmt.Println("Press Enter to return to menu...")
	fmt.Scanln()
}

func deleteTransaction(db *sql.DB) {
	// Input ID Transaksi yang akan dihapus
	var transactionid int
	var checkID int
	fmt.Print("Enter Transaction ID: ")
	fmt.Scanln(&transactionid)

	errs := db.QueryRow("SELECT COUNT(*) FROM transactions WHERE transaction_id = $1", transactionid).Scan(&checkID)
	if errs != nil {
		panic(errs)
	}
	if checkID == 0 {
		fmt.Println("\nInvalid Employee ID.")
		fmt.Println("Press Enter to return to menu...")
		fmt.Scanln()
		return
	}

	// Hapus data transaksi dari database
	_, err := db.Exec("DELETE FROM transactions WHERE transaction_id = $1", transactionid)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nTransaction deleted successfully.")
	}
	fmt.Println("Press Enter to return to menu...")
	fmt.Scanln()
}

func menuEmployee(db *sql.DB) {
	for {
		fmt.Println("\n=== Enigma Laundry ===")
		fmt.Println("1. Insert Employee")
		fmt.Println("2. Update Employee")
		fmt.Println("3. Delete Employee")
		fmt.Println("4. View Employee")
		fmt.Println("0. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			insertEmployee(db)
		case 2:
			updateEmployee(db)
		case 3:
			deleteEmployee(db)
		case 4:
			viewEmployee(db)
		case 0:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func insertEmployee(db *sql.DB) {
	// Input data Pegawai
	var name, phoneNumber string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter Name: ")
	scanner.Scan() // Baca input baris teks
	name = scanner.Text()
	fmt.Print("Enter Phone Number: ")
	fmt.Scanln(&phoneNumber)

	// Insert data pelanggan ke database
	_, err := db.Exec("INSERT INTO employee (name, phone_number) VALUES ($1, $2)", name, phoneNumber)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nEmployee inserted successfully.")
	}
	fmt.Println("Press Enter to return to menu...")
	fmt.Scanln()
}

func updateEmployee(db *sql.DB) {
	// Input ID pegawai yang akan diupdate
	var employeeID int
	var checkID int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Employee ID: ")
	fmt.Scanln(&employeeID)

	errs := db.QueryRow("SELECT COUNT(*) FROM employee WHERE employee_id = $1", employeeID).Scan(&checkID)
	if errs != nil {
		panic(errs)
	}
	if checkID == 0 {
		fmt.Println("\nInvalid Employee ID.")
		fmt.Println("Press Enter to return to menu...")
		fmt.Scanln()
		return
	}

	// Input data pegawai yang baru
	var name, phoneNumber string
	fmt.Print("Enter New Name: ")
	scanner.Scan()
	name = scanner.Text()
	fmt.Print("Enter New Phone Number: ")
	fmt.Scanln(&phoneNumber)

	// Update data pegawai di database
	_, err := db.Exec("UPDATE employee SET name = $1, phone_number = $2 WHERE employee_id = $3", name, phoneNumber, employeeID)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nEmployee updated successfully.")
	}
	fmt.Println("Press Enter to return to menu...")
	fmt.Scanln()
}

func deleteEmployee(db *sql.DB) {
	// Input ID pegawai yang akan dihapus
	var employeeID int
	var checkID int
	fmt.Print("Enter Employee ID: ")
	fmt.Scanln(&employeeID)

	errs := db.QueryRow("SELECT COUNT(*) FROM employee WHERE employee_id = $1", employeeID).Scan(&checkID)
	if errs != nil {
		panic(errs)
	}
	if checkID == 0 {
		fmt.Println("\nInvalid Employee ID.")
		fmt.Println("Press Enter to return to menu...")
		fmt.Scanln()
		return
	}

	// Hapus data pegawai dari database
	_, err := db.Exec("DELETE FROM employee WHERE employee_id = $1", employeeID)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nEmployee deleted successfully.")
	}
	fmt.Println("Press Enter to return to menu...")
	fmt.Scanln()
}

func viewEmployee(db *sql.DB) {
	// Query database untuk menampilkan data employee
	rows, err := db.Query("SELECT * FROM employee")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	fmt.Println("=== Employee ===")
	fmt.Println("ID | Name | Phone Number")
	for rows.Next() {
		var (
			employeeID  int
			name        string
			phoneNumber string
		)
		if err := rows.Scan(&employeeID, &name, &phoneNumber); err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("%d | %s | %s\n", employeeID, name, phoneNumber)
		}

	}
	fmt.Println("\nPress Enter to return to menu...")
	fmt.Scanln()
}

func updateServices(db *sql.DB) {
	var serviceID int
	var checkID int
	fmt.Print("Enter Services ID: ")
	fmt.Scanln(&serviceID)

	errs := db.QueryRow("SELECT COUNT(*) FROM public.laundry_services WHERE service_id = $1", serviceID).Scan(&checkID)
	if errs != nil {
		panic(errs)
	}
	if checkID == 0 {
		fmt.Println("\nInvalid Service ID.")
		fmt.Println("Press Enter to return to menu...")
		fmt.Scanln()
		return
	}

	// Input data Service yang baru
	var name string
	var price float64
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter New Service Name: ")
	scanner.Scan()
	name = scanner.Text()
	fmt.Print("Enter New Price: ")
	fmt.Scanln(&price)

	// Update data service di database
	_, err := db.Exec("UPDATE public.laundry_services SET service_name = $1, price = $2 WHERE Service_id = $3", name, price, serviceID)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nService updated successfully.")
	}
	fmt.Println("Press Enter to return to menu...")
	fmt.Scanln()
}

func deleteServices(db *sql.DB) {
	var serviceID int
	var checkID int
	fmt.Print("Enter Service ID: ")
	fmt.Scanln(&serviceID)

	errs := db.QueryRow("SELECT COUNT(*) FROM laundry_services WHERE service_id = $1", serviceID).Scan(&checkID)
	if errs != nil {
		panic(errs)
	}
	if checkID == 0 {
		fmt.Println("\nInvalid Service ID.")
		fmt.Println("Press Enter to return to menu...")
		fmt.Scanln()
		return
	}

	// Hapus data Service dari database
	_, err := db.Exec("DELETE FROM laundry_services WHERE service_id = $1", serviceID)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nService deleted successfully.")
	}
	fmt.Println("Press Enter to return to menu...")
	fmt.Scanln()
}

func viewServices(db *sql.DB) {
	// Query database untuk menampilkan jenis service
	rows, err := db.Query("SELECT * FROM public.laundry_services")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	fmt.Println("=== Customers ===")
	fmt.Println("ID | Service Name | Price")
	for rows.Next() {
		var (
			serviceID   int
			serviceName string
			price       string
		)
		if err := rows.Scan(&serviceID, &serviceName, &price); err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("%d | %s | %s\n", serviceID, serviceName, price)
		}

	}
	fmt.Println("\nPress Enter to return to menu...")
	fmt.Scanln()
}

func viewCustomer(db *sql.DB) {
	// Query database untuk menampilkan pelanggan
	rows, err := db.Query("SELECT * FROM customers")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	fmt.Println("=== Customers ===")
	fmt.Println("ID | Name | Phone Number")
	for rows.Next() {
		var (
			customerID  int
			name        string
			phoneNumber string
		)
		if err := rows.Scan(&customerID, &name, &phoneNumber); err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("%d | %s | %s\n", customerID, name, phoneNumber)
		}

	}
	fmt.Println("\nPress Enter to return to menu...")
	fmt.Scanln()
}

func insertCustomer(db *sql.DB) {
	// Input data pelanggan
	var name, phoneNumber string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter Name: ")
	scanner.Scan() // Baca input baris teks
	name = scanner.Text()
	fmt.Print("Enter Phone Number: ")
	fmt.Scanln(&phoneNumber)

	// Insert data pelanggan ke database
	_, err := db.Exec("INSERT INTO customers (name, phone_number) VALUES ($1, $2)", name, phoneNumber)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nCustomer inserted successfully.")
	}
	fmt.Println("Press Enter to return to menu...")
	fmt.Scanln()
}

func updateCustomer(db *sql.DB) {
	// Input ID pelanggan yang akan diupdate
	scanner := bufio.NewScanner(os.Stdin)
	var customerID int
	var checkID int
	fmt.Print("Enter Customer ID: ")
	fmt.Scanln(&customerID)

	errs := db.QueryRow("SELECT COUNT(*) FROM customers WHERE customer_id = $1", customerID).Scan(&checkID)
	if errs != nil {
		panic(errs)
	}
	if checkID == 0 {
		fmt.Println("\nInvalid Customer ID.")
		fmt.Println("Press Enter to return to menu...")
		fmt.Scanln()
		return
	}

	// Input data pelanggan yang baru
	var name, phoneNumber string
	fmt.Print("Enter New Name: ")
	scanner.Scan()
	name = scanner.Text()
	fmt.Print("Enter New Phone Number: ")
	fmt.Scanln(&phoneNumber)

	// Update data pelanggan di database
	_, err := db.Exec("UPDATE customers SET name = $1, phone_number = $2 WHERE customer_id = $3", name, phoneNumber, customerID)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Customer updated successfully.")
	}
	fmt.Println("\nPress Enter to return to menu...")
	fmt.Scanln()
}

func deleteCustomer(db *sql.DB) {
	// Input ID pelanggan yang akan dihapus
	var customerID int
	var checkID int
	fmt.Print("Enter Customer ID: ")
	fmt.Scanln(&customerID)

	errs := db.QueryRow("SELECT COUNT(*) FROM customers WHERE customer_id = $1", customerID).Scan(&checkID)
	if errs != nil {
		panic(errs)
	}
	if checkID == 0 {
		fmt.Println("\nInvalid Customer ID.")
		fmt.Println("Press Enter to return to menu...")
		fmt.Scanln()
		return
	}

	// Hapus data pelanggan dari database
	_, err := db.Exec("DELETE FROM customers WHERE customer_id = $1", customerID)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Customer deleted successfully.")
	}
	fmt.Println("\nPress Enter to return to menu...")
	fmt.Scanln()
}

func viewTransactions(db *sql.DB) {
	// Query database untuk menampilkan transaksi

	rows, err := db.Query("SELECT t.transaction_id, c.name as customer_name, s.service_name, e.name as employee_name,t.quantity, t.total_price, t.transaction_date,t.finish_date FROM transactions t JOIN customers c ON t.customer_id = c.customer_id JOIN laundry_services s ON t.service_id = s.service_id JOIN employee e ON t.employee_id = e.employee_id")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	fmt.Println("=== Transactions ===")
	fmt.Println("ID | Customer name | Employee Name | Service Name | Quantity | Total Price | Transaction date | Finish Date ")
	for rows.Next() {
		var (
			transactionID   int
			customerName    string
			employeeNaame   string
			serviceName     string
			quantity        int
			total_price     float64
			transactionDate string
			finishDate      string
		)
		if err := rows.Scan(&transactionID, &customerName, &employeeNaame, &serviceName, &quantity, &total_price, &transactionDate, &finishDate); err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("%d | %s | %s | %s | %d | %.2f | %s | %s\n", transactionID, customerName, employeeNaame, serviceName, quantity, total_price, transactionDate, finishDate)
		}

	}
	fmt.Println("\nPress Enter to return to menu...")
	fmt.Scanln()
}

func insertTransaction(db *sql.DB) {

	fmt.Println("\n=== Insert Transaction ===")
	// Menampilkan daftar pelanggan untuk referensi
	var customerID, employeeID, serviceID, quantity int
	var finishDate string
	fmt.Print("Enter Customer ID: ")
	fmt.Scanln(&customerID)
	fmt.Print("Enter Employee ID: ")
	fmt.Scanln(&employeeID)
	fmt.Print("Enter Service ID: ")
	fmt.Scanln(&serviceID)
	fmt.Print("Enter Quantity: ")
	fmt.Scanln(&quantity)
	fmt.Print("Enter Finish Date (YYYY-MM-DD): ")
	fmt.Scanln(&finishDate)

	// Memeriksa apakah customerID dan serviceID yang dimasukkan valid
	var checkID int
	err := db.QueryRow("SELECT COUNT(*) FROM customers WHERE customer_id = $1", customerID).Scan(&checkID)
	if err != nil {
		panic(err)
	}
	if checkID == 0 {
		fmt.Println("Invalid Customer ID.")
		return
	}

	err = db.QueryRow("SELECT COUNT(*) FROM laundry_services WHERE service_id = $1", serviceID).Scan(&checkID)
	if err != nil {
		panic(err)
	}
	if checkID == 0 {
		fmt.Println("Invalid Service ID.")
		return
	}

	// Memeriksa apakah jumlah yang dimasukkan valid
	if quantity <= 0 {
		fmt.Println("Quantity must be greater than 0.")
		return
	}

	// Memeriksa apakah format tanggal selesai valid
	_, err = time.Parse("2006-01-02", finishDate)
	if err != nil {
		fmt.Println("Invalid End Date format. Please use YYYY-MM-DD format.")
		return
	}

	// Memperoleh harga layanan dari database
	var price float64
	err = db.QueryRow("SELECT price FROM laundry_services WHERE service_id = $1", serviceID).Scan(&price)
	if err != nil {
		panic(err)
	}

	// Menghitung total harga transaksi
	totalPrice := price * float64(quantity)

	// Memasukkan transaksi ke dalam database
	_, err = db.Exec("INSERT INTO transactions (customer_id,employee_id, service_id, quantity, total_price, transaction_date, finish_date) VALUES ($1, $2, $3, $4,$5, CURRENT_DATE, $6)",
		customerID, employeeID, serviceID, quantity, totalPrice, finishDate)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nTransaction inserted successfully.")
	}
	fmt.Println("Press Enter to return to menu...")
	fmt.Scanln()
}

func insertServices(db *sql.DB) {
	// Input data pelanggan
	var name, price string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter Service Name: ")
	scanner.Scan()
	name = scanner.Text()
	fmt.Print("Enter Price: ")
	fmt.Scanln(&price)

	// Insert data pelanggan ke database
	_, err := db.Exec("INSERT INTO laundry_services (service_name, price) VALUES ($1, $2)", name, price)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nService inserted successfully.")
	}
	fmt.Println("Press Enter to return to menu...")
	fmt.Scanln()
}
