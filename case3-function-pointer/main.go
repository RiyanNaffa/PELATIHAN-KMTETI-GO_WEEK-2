// Case 3: Melakukan perubahan pada elemen map menggunakan function dan pointer

package main

import "fmt"

type Mahasiswa struct {
	Name    string
	IsAsdos bool
}

// Terdapat kesalahan pada fungsi updateStatus. Seharusnya fungsi tersebut
// melakukan update status asdos bagi mahasiswa bernama Caca.
// Perbaiki fungsi updateStatus, pendefinisian variabel asdos, serta
// pemanggilan fungsi updateStatus

func updateStatus(mhs []*Mahasiswa) {
	mhs[2].IsAsdos = true
}

func main() {
	m1 := Mahasiswa{Name: "Andi", IsAsdos: true}
	m2 := Mahasiswa{Name: "Budi", IsAsdos: true}
	m3 := Mahasiswa{Name: "Caca", IsAsdos: false}

	asdos := []*Mahasiswa{&m1, &m2, &m3}

	updateStatus(asdos)
	fmt.Println(m3)
}
