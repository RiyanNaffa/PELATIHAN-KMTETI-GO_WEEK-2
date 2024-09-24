# CASE 4: ERROR HANDLING

Untuk menjalankan program ini, dua buah *library* lainnya selain `fmt` perlu di-*import*.

```go
import (
	"errors"
	"fmt"
	"math"
)
```

Didefinisikan tiga buah variabel bertipe *error* yang bekerja sesuai dengan apa yang tertera sebagai berikut.

```go
var ErrNegativeNumber = errors.New("square root of negative number will result complex number")
var ErrDevideByZero = errors.New("cannot divide by zero")
var ErrModByZero = errors.New("cannot find remainders of division by zero")
```

Permasalahan pada program adalah ketika fungsi `findRoot()` diberi masukan berupa bilangan negatif. Karena fungsi ini hanya memproses bilangan riil saja, masukan bilangan negatif harus memberikan nilai *error* yang tepat, yakni `ErrNegativeNumber`.

Akan tetapi, kode yang telah disediakan memiliki kembalian nilai *error* yang tidak sesuai, yakni `ErrDevideByZero`.

```go
func findRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, ErrDevideByZero
	}
	res := math.Sqrt(num)
	return res, nil
}
```

Jika kode berikut dijalankan, hasilnya tidak sesuai dengan apa yang diharapkan.

```go
myNum := -20.7
myNumRoot, err := findRoot(myNum)
if err == ErrNegativeNumber {
    fmt.Println(err.Error())
    return
}
fmt.Println(myNumRoot)
```

```
Output:
0
```

`Output` adalah 0. Hal ini disebabkan karena fungsi `findRoot()` mengembalikan nilai 0 dan `ErrDevideByZero` dengan masukan angka negatif. Jika kode di atas dijalankan, pada blok seleksi, *error* hasil operasi tidak cocok dengan operasi logika, sehingga blok seleksi di bawahnya akan dilewati begitu saja dan langsung mencetak angkanya.

Koreksi kode di atas adalah dengan mengoreksi nilai kembalian pada fungsi `findRoot()`.

```go
func findRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, ErrNegativeNumber
	}
	// ...
}
```

Jika dijalankan, kode tersebut akan menghasilkan *output* yang sesuai.

```go
myNum := -20.7
myNumRoot, err := findRoot(myNum)
if err == ErrNegativeNumber {
    fmt.Println(err.Error())
    return
}
fmt.Println(myNumRoot)
```

```
Output:
square root of negative number will result complex number
```