package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Fungsi untuk inisialisasi koneksi ke database
func initDB() {
	var err error
	// Ganti dengan koneksi yang sesuai
	dsn := "root:Insider2816.@tcp(127.0.0.1:3306)/book_store?parseTime=true" //Parameter parseTime=true memastikan bahwa driver MySQL akan mengkonversi kolom DATE dan DATETIME langsung menjadi time.Time di Go
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection established")
}
