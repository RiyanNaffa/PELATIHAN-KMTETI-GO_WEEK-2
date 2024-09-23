# CASE 2: BOOK STORE

## Deskripsi Kasus

Transaksi pada suatu toko buku dapat dimodelkan dengan sebuah *struct* `Transaction` yang memiliki *field*:
1. `Total`: total harga (*unsigned integer*)
2. `Date`: tanggal transaksi (*string*)
3. `Books`: buku-buku yang dibeli (*slice of* `Book`)

Setiap buku pada toko tersebut juga dimodelkan dalam bentuk *struct* `Book` dengan *field*:
1. `Title`: judul buku (*string*)
2. `Writer`: penulis buku (*string*)
3. `Price`: harga buku (*unsigned integer*)
4. `Count`: banyak buku yang dibeli (*unsigned integer*)

Terdapat sebuah transaksi yang dideskripsikan sebagai berikut:

Pada tanggal **23-01-2021**, terjadi sebuah transaksi. Pembeli membeli 2 buah buku. Buku pertama berjudul "**Laut Bercerita**" karangan **Leila S. Chudori** seharga **108000** sebanyak **1** buah. Buku berikutnya "**Bumi**" karangan **Tere Liye** seharga **97000** sebanyak **1** buah.

## Deklarasi *Struct*

*Struct* `Transaction` dapat ditulis sebagai berikut.

```go
type Transaction struct {
	Total uint
	Date  string
	Books []Book
}
```

Sedangkan *struct* `Book` sendiri dapat ditulis sebagai berikut.

```go
type Book struct {
	Title  string
	Writer string
	Price  uint
	Count  uint
}
```

## Definisi Buku-Buku

Buku-buku pada kasus tersebut, yakni "Hujan" dan "Laut Bercerita" dibuat menjadi dua buah variabel bertipe *struct* `Book` dengan nama `hujan` dan `laut` berturut-turut.

```go
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
```

## Definisi Transaksi

Transaksi dibuat menjadi sebuah variabel *struct* `Transaction` dengan nama `tx1`.

```go
tx1 := Transaction{}
```

Pendefinisian `tx1` tidak dilakukan secara *literal* seperti kedua buku sebelumnya, tetapi dilakukan secara sekuensial untuk memudahkan pemrograman.

Pertama-tama, buku-buku yang ingin dibeli didefinisikan dahulu pada *field* `Books` dengan melakukan *append*.

```go
tx1.Books = append(tx1.Books, hujan, laut)
```

*Slice* `txt1.Books` akan bernilai seperti di bawah ini.

```
[{Hujan Tere Liye 97000 1} {Laut Bercerita Leila S. Chudori 108000 1}]
```

Kemudian, penghitungan total harga transaksi dilakukan. Awalnya, total harga didefinisikan dengan nilai 0 sebelum dilakukan iterasi penjumlahan harga buku. 

Panjang *slice* `txt1.Books` dihitung terlebih dahulu untuk menentukan banyak iterasi yang perlu dilakukan untuk tiap bukunya. Lalu, untuk setiap buku pada *slice*, total harga (`.Total`) ditambah dengan banyak buku (`.Books[i].Count`) dikali dengan harga buku (`.Books[i].Price`).

```go
tx1.Total = 0

n := len(tx1.Books)
for i := 0; i < n; i++ {
    tx1.Total += tx1.Books[i].Count * tx1.Books[i].Price
}
```

Nilai dari `tx1.Total` akan menjadi seperti berikut.

```
205000
```

Terakhir, tanggal transaksi didefinisikan pada tanggal 23-01-2021.

```go
tx1.Date = "23-01-2021"
```

Variabel `tx1` telah selesai didefinisikan.

```go
fmt.Println(tx1)
```

```
Output:
{205000 23-01-2021 [{Hujan Tere Liye 97000 1} {Laut Bercerita Leila S. Chudori 108000 1}]}
```