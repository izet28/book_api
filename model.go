package main

import (
	"fmt"
	"time"
)

// Custom type untuk mempermudah serialisasi/deserialisasi time.Time
type Date struct {
	time.Time
}

// Format untuk serialisasi/deserialisasi
const dateFormat = "2006-01-02"

// Custom UnmarshalJSON untuk format "YYYY-MM-DD"
func (d *Date) UnmarshalJSON(b []byte) error {
	str := string(b)
	str = str[1 : len(str)-1] // Menghapus tanda kutip
	parsedTime, err := time.Parse(dateFormat, str)
	if err != nil {
		return err
	}
	d.Time = parsedTime
	return nil
}

// Custom MarshalJSON untuk format "YYYY-MM-DD"
func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, d.Format(dateFormat))), nil
}

// Struct Book merepresentasikan data buku
type Book struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Author        string  `json:"author"`
	PublishedDate Date    `json:"published_date" sql:"type:date"` // Pastikan untuk mendefinisikan tipe date
	Price         float64 `json:"price"`
}
