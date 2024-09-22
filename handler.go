package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Fungsi untuk mendapatkan semua buku
func getBooks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, author, published_date, price FROM books")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		// Scan langsung ke time.Time
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublishedDate.Time, &book.Price); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Fungsi untuk mendapatkan buku berdasarkan ID
func getBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var book Book
	err = db.QueryRow("SELECT id, title, author, published_date, price FROM books WHERE id = ?", id).
		Scan(&book.ID, &book.Title, &book.Author, &book.PublishedDate, &book.Price)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// Fungsi untuk menambahkan buku baru
func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book

	// Decode JSON dari body request
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Menyimpan data ke database menggunakan book.PublishedDate.Time
	result, err := db.Exec("INSERT INTO books (title, author, published_date, price) VALUES (?, ?, ?, ?)",
		book.Title, book.Author, book.PublishedDate.Time, book.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mendapatkan ID dari buku yang baru saja dibuat
	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	book.ID = int(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// Fungsi untuk memperbarui buku berdasarkan ID
func updateBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE books SET title = ?, author = ?, published_date = ?, price = ? WHERE id = ?",
		book.Title, book.Author, book.PublishedDate.Time, book.Price, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	book.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// Fungsi untuk menghapus buku berdasarkan ID
func deleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Book deleted"})
}
