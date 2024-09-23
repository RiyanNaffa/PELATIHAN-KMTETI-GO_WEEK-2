// CASE 1: Troubleshooting
// identifikasi dan koreksi kesalahan pada program di bawah

// EXPECTED OUTPUT:
// [0 1 2 4 5 6]
// 6
// Umur Imanuel tahun ini: 21
// Siswa baru, Siska berumur: 19
// [11 13 17 19 23 29]

package main

import "fmt"

func main() {
	// Silakan melakukan perubahan pada program di bawah agar output
	// ketika program dijalankan sesuai dengan EXPECETED OUTPUT

	// Part 1:	Mengapa terdapat error pada baris-baris kode berikut?
	// 			Perbaiki baris-baris berikut agar program dapat dijalankan
	// myArr := [5]uint{0, 1, 2, 4, 5, 6}
	myArr := [6]uint{0, 1, 2, 4, 5, 6}

	fmt.Println(myArr)
	fmt.Println(myArr[5])
	// -- End of Part 1 --

	// Part 2:	Output program tidak sesuai dengan EXPECETED OUTPUT
	//			Jangan lakukan perubahan langsung pada stdAge di awal.
	//        	Coba gunakan key untuk melakukan perubahan dan menambahkan
	//        	item ke map.
	stdAge := map[string]uint{
		"Agus":     19,
		"Irfan":    20,
		"Johannes": 19,
		"Imanuel":  20,
	}

	stdAge["Imanuel"] = 21
	fmt.Println("Umur Imanuel tahun ini:", stdAge["Imanuel"])
	stdAge["Siska"] = 19
	fmt.Println("Siswa baru, Siska berumur:", stdAge["Siska"])
	// -- End of Part 2 --

	// Part 3:	Baris-baris kode berikut seharusnya akan menampilkan
	//			bilangan prima antara 10 hingga 30.
	//			Identifikasi letak kesalahan logika pada program dan
	//			koreksi agar menampilkan output yang sesuai.
	var primes []int
	for i := 10; i < 30; i++ {
		if i%2 != 0 && i%3 != 0 && i%5 != 0 && i%7 != 0 {
			primes = append(primes, i)
		}
	}

	fmt.Println(primes)
}
