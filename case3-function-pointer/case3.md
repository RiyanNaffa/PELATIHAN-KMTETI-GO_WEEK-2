# CASE 3: FUNCTION & POINTER

```go
type Mahasiswa struct {
	Name    string
	IsAsdos bool
}
```

Terdapat suatu *struct* `Mahasiswa` dengan *field* seperti di atas. Kemudian, tiga variabel `m1`, `m2`, dan `m3` dideklarasikan dan didefinisikan sebagai variabel *struct* `Mahasiswa` pada fungsi `main()` seperti berikut.

```go
m1 := Mahasiswa{Name: "Andi", IsAsdos: true}
m2 := Mahasiswa{Name: "Budi", IsAsdos: true}
m3 := Mahasiswa{Name: "Caca", IsAsdos: false}
```

Nilai dari *field* `IsAsdos` pada `m3`, yaitu Caca, ingin diubah dengan sebuah *user-defined function* lain. Sebelum itu, ketiga variabel tersebut dimasukkan ke dalam suatu *slice* bertipe `Mahasiswa`.

```go
asdos := []Mahasiswa{m1, m2, m3}
```

Suatu fungsi untuk mengubah nilai `IsAsdos` pada `m3` didefinisikan sebagai berikut. Fungsi mengambil *slice* bertipe `Mahasiswa` dengan nama parameter `mhs` dan mengubah nilai dari `IsAsdos` untuk elemen ke-3 atau elemen dengan indeks ke-2.

```go
func updateStatus(mhs []Mahasiswa) {
	mhs[2].IsAsdos = true
}
```

Fungsi tersebut kemudian dipanggil untuk melakukan modifikasi pada *slice* `asdos`.

```go
updateStatus(asdos)
fmt.Println(m3)
fmt.Println(asdos[2])
```

```
Output:
{Caca false}
{Caca true}
```

Mengapa nilai `IsAsdos` dari variabel `m3` masih bernilai *false*?

Perlu diketahui bahwa *slice* `asdos`, *slice* yang dikenai modifikasi, merupakan *slice* bertipe `Mahasiswa` yang memiliki ***value*** yang sama dengan ketiga variabel `m1`, `m2`, dan `m3`.

Akan tetapi, *slice* tersebut tidak berhubungan langsung dengan variabel `m1`, `m2`, dan `m3`. Sehingga, jika ketiga variabel tersebut dipanggil, nilainya masih sama seperti deklarasi di awal.

Untuk mengubah langsung ketiga variabel tersebut, modifikasi harus dilakukan langsung. Hal ini dapat dilakukan dengan melakukan *reference passing* pada *slice* `asdos`.

```go
asdos := []*Mahasiswa{&m1, &m2, &m3}
```

*Slice* `asdos` kini merupakan *slice* bertipe *pointer* `Mahasiswa` yang memiliki elemen yang merupakan alamat dari ketiga variabel. Sehingga, jika dilakukan modifikasi pada `asdos`, variabel `m1`, `m2`, dan `m3` juga ikut berubah.

Kemudian, fungsi `updateStatus()` juga perlu diubah supaya cocok dengan masukannya, yakni tipe *pointer* `Mahasiswa`.

```go
func updateStatus(mhs []*Mahasiswa) {
	mhs[2].IsAsdos = true
}
```

Ketika dijalankan, maka baris-baris berikut akan memberikan *output* sesuai dengan yang diharapkan.

```go
updateStatus(asdos)
fmt.Println(m3)
```

```
Output:
{Caca true}
```