package main

import (
	"fmt"
	"strings"
)

func main() {
	// Struct Alamat
	type Alamat struct {
		Jalan   string
		Kota    string
		KodePos string
	}

	// Struct Mahasiswa
	type Mahasiswa struct {
		Nama   string
		Umur   int
		Hobi   []string
		Alamat Alamat
		NoHP   *string
	}

	// Contoh data mahasiswa
	noHP1 := "08123456789"
	noHP3 := "08987654321"

	mahasiswas := []Mahasiswa{
		{
			Nama: "Rina",
			Umur: 25,
			Hobi: []string{"Membaca", "Menulis", "Coding"},
			Alamat: Alamat{
				Jalan:   "Jl. Mawar No. 10",
				Kota:    "Bandung",
				KodePos: "40123",
			},
			NoHP: &noHP1,
		},
		{
			Nama: "Budi",
			Umur: 22,
			Hobi: []string{"Bermain Musik", "Olahraga"},
			Alamat: Alamat{
				Jalan:   "Jl. Melati No. 5",
				Kota:    "Jakarta",
				KodePos: "10230",
			},
			NoHP: nil, // belum diisi
		},
		{
			Nama: "Sari",
			Umur: 24,
			Hobi: []string{"Fotografi", "Traveling", "Membaca"},
			Alamat: Alamat{
				Jalan:   "Jl. Anggrek No. 7",
				Kota:    "Surabaya",
				KodePos: "60241",
			},
			NoHP: &noHP3,
		},
	}

	// Tampilkan data mahasiswa
	for i, mhs := range mahasiswas {
		fmt.Println("=================================")
		fmt.Println("Mahasiswa ke-", i+1)
		fmt.Println("Nama   : ", mhs.Nama)
		fmt.Println("Umur   : ", mhs.Umur)
		fmt.Println("Hobi   : ", strings.Join(mhs.Hobi, ", "))
		// fmt.Printf("Alamat : %s , %s - %s\n", mhs.Alamat.Jalan, mhs.Alamat.Kota, mhs.Alamat.KodePos)
		alamatLengkap := fmt.Sprintf("Alamat : %s , %s - %s", mhs.Alamat.Jalan, mhs.Alamat.Kota, mhs.Alamat.KodePos)
		fmt.Println(alamatLengkap)

		if mhs.NoHP != nil {
			fmt.Println("No HP  : ", *mhs.NoHP)
		} else {
			fmt.Println("No HP  : (belum diisi)")
		}
	}
	fmt.Println("=================================")
}
