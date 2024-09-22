package main

import (
	"log"
	"net/http"
)

func main() {
	// Inisialisasi koneksi ke database
	initDB()
	defer db.Close()

	// Route endpoints
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Jika ada query id, tampilkan buku berdasarkan ID, jika tidak tampilkan semua buku
			if r.URL.Query().Get("id") != "" {
				getBook(w, r)
			} else {
				getBooks(w, r)
			}
		case http.MethodPost:
			createBook(w, r)
		case http.MethodPut:
			updateBook(w, r)
		case http.MethodDelete:
			deleteBook(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Jalankan server
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
