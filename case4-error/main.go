// Case 4: Error handling

package main

import (
	"errors"
	"fmt"
	"math"
)

var ErrNegativeNumber = errors.New("square root of negative number will result complex number")
var ErrDevideByZero = errors.New("cannot divide by zero")
var ErrModByZero = errors.New("cannot find remainders of division by zero")

// Kita ingin mencari akar kuadrat dari sebuah bilangan. Akan tetapi, kita
// tahu bahwa akar kuadrat dari bilangan negatif akan menghasilkan
// bilangan kompleks.
// Pada main function, myNum didefinisikan bernilai -20.7, seharusnya program
// akan menampilkan error untuk kasus ini.
// Lakukan perubahan pada fungsi findRoot agar mengembalikan error yang
// sesuai.
func findRoot(num float64) (float64, error) {
	if num < 0 {
		// return 0, ErrDevideByZero
		return 0, ErrNegativeNumber
	}
	res := math.Sqrt(num)
	return res, nil
}

// End of Section -- Jangan lakukan perubahan pada main function

func main() {
	myNum := -20.7
	myNumRoot, err := findRoot(myNum)
	if err == ErrNegativeNumber {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(myNumRoot)
}
