// Case 2: Book Store

package main

import "fmt"

// Sebuah transaksi di toko buku akan disimpan ke dalam database.
// Untuk itu, perlu dibuat model dari transaksi menggunakan struct.
// Setiap transaksi memiliki beberapa field, yaitu:
// - Total: total harga yang perlu dibayar
// - Date: tanggal terjadinya transaksi berupa string
// - Books: buku-buku yang dibeli oleh pembeli

// Setiap buku juga memiliki beberapa field, yaitu:
// - Title: judul buku
// - Writer: nama penulis buku
// - Price: harga buku
// - Count: banyak buku yang dibeli oleh pembeli

// Pada tanggal 23-01-2021, terjadi transaksi dengan total harga sebesar
// 32000. Pembeli membeli 2 buah buku. Buku pertama berjudul "Laut Bercerita"
// karangan Leila S. Chudori seharga 108000 sebanyak 1 buah. Buku berikutnya
// berjudul "Bumi" karangan Tere Liye seharga 97000 sebanyak 1 buah.

type Book struct {
	Title  string
	Writer string
	Price  uint
	Count  uint
}
type Transaction struct {
	Total uint
	Date  string
	Books []Book
}

func main() {
	hujan := Book{
		Title:  "Hujan",
		Writer: "Tere Liye",
		Price:  97000,
		Count:  1,
	}
	laut := Book{
		Title:  "Laut Bercerita",
		Writer: "Leila S. Chudori",
		Price:  108000,
		Count:  1,
	}

	// Modifikasi pendefinisian variable tx1 sehingga sesuai dengan kasus
	// yang diceritakan.

	tx1 := Transaction{}
	tx1.Books = append(tx1.Books, hujan, laut)

	tx1.Total = 0
	// Iterasi sepanjang slice Transaction.Books untuk menghitung harga
	n := len(tx1.Books)
	for i := 0; i < n; i++ {
		tx1.Total += tx1.Books[i].Count * tx1.Books[i].Price
	}

	tx1.Date = "23-01-2021"

	fmt.Println(tx1)
}
