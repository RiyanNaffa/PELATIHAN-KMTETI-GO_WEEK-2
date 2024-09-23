# CASE 1: TROUBLESHOOTING

## Part 1

```go
myArr := [5]uint{0, 1, 2, 4, 5, 6}

fmt.Println(myArr)
fmt.Println(myArr[5])
```

*Array* `myArr`, yang merupakan tipe data statis, dideklarasikan dengan ukuran **5 elemen `uint`**. Pada inisialisasi nilai awalnya, diketahui bahwa *array* diberi **6 buah** elemen `uint`. Ketidakcocokan elemen inilah yang memunculkan *error*.

```
Output:
index 5 is out of bounds (>= 5)
invalid argument: index 5 out of bounds [0:5]
```

Perbaikan yang tepat adalah untuk menambahkan ukuran *array*.

```go
myArr := [6]uint{0, 1, 2, 4, 5, 6}
```

Kini, perintah berikut menghasilkan *output* yang diinginkan.

```go
fmt.Println(myArr)
fmt.Println(myArr[5])
```

```
Output:
[0 1 2 4 5 6]
6
```

## Part 2

```go
stdAge := map[string]uint{
		"Agus":     19,
		"Irfan":    20,
		"Johannes": 19,
		"Imanuel":  20,
}
```

Subkasus pertama adalah mengubah nilai dari suatu *key* tanpa harus melakukan perubahan pada inisialisasi variabel.
```go
fmt.Println("Umur Imanuel tahun ini:", stdAge["Imanuel"])
```

```
Output:
Umur Imanuel tahun ini: 20
```

Supaya umur Imanuel bernilai 21, nilai dari *key* `Imanuel` harus diubah lewat modifikasi.
```go
stdAge["Imanuel"] = 21
fmt.Println("Umur Imanuel tahun ini:", stdAge["Imanuel"])
```

```
Output: 
Umur Imanuel tahun ini: 21
```

Subkasus selanjutnya adalah nilai dari *key* yang tidak ada. Perintah di bawah akan mencoba mencari *key* `Siska`. Akan tetapi, `Siska` tidak dapat ditemukan sehingga mengembalikan nilai nol dari tipe data *map* (`uint`).

```go
fmt.Println("Siswa baru, Siska berumur:", stdAge["Siska"])
```

```
Output:
Siswa baru, Siska berumur: 0
```

Dengan mendeklarasikan dan menginisialisasikan pasangan *key* dan *value* baru dari *map* `stdAge`, `Siska` dapat ditemukan.

```go
stdAge["Siska"] = 19
fmt.Println("Siswa baru, Siska berumur:", stdAge["Siska"])
```

```
Output:
Siswa baru, Siska berumur: 19
```

## Part 3

```go
var primes []int
for i := 10; i < 30; i++ {
    if i%2 != 0 && i%3 != 0 && i%5 != 0 && i%7 != 0 {
    }
    primes = append(primes, i)
}

fmt.Println(primes)
```

```
Output:
[10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29]
```

Baris untuk memasukkan angka ke dalam *array* bilangan prima `primes` terletak di luar blok seleksi `if`, sehingga baris tersebut akan dijalankan apa pun nilai `i` pada *for-loop*.

Supaya menghasilkan *output* yang diinginkan, baris tersebut dapat dipindah ke dalam blok `if`.

```go
var primes []int
for i := 10; i < 30; i++ {
    if i%2 != 0 && i%3 != 0 && i%5 != 0 && i%7 != 0 {
        primes = append(primes, i)
    }
}

fmt.Println(primes)
```

```
Output:
[11 13 17 19 23 29]
```