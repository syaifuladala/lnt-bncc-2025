package main

import (
	"fmt"
	"mahasiswa/models"
	"mahasiswa/utils"
	"strings"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Terjadi Kesalahan:", r)
		}
	}()

	// Contoh data mahasiswa
	noHP1 := "08123456789"
	noHP3 := "08987654321"

	mahasiswas := []models.Mahasiswa{
		{
			Nama:  "Rina",
			Umur:  25,
			Hobi:  []string{"Membaca", "Menulis", "Coding"},
			NoHP:  &noHP1,
			Nilai: []int{90, 85, 88},
			Alamat: models.Alamat{
				Jalan:   "Jl. Mawar No. 10",
				Kota:    "Bandung",
				KodePos: "40123",
			},
		},
		{
			Nama:  "Budi",
			Umur:  22,
			Hobi:  []string{"Bermain Musik", "Olahraga"},
			NoHP:  nil, // belum diisi
			Nilai: []int{50, 45, 55},
			Alamat: models.Alamat{
				Jalan:   "Jl. Melati No. 5",
				Kota:    "Jakarta",
				KodePos: "10230",
			},
		},
		{
			Nama:  "Sari",
			Umur:  24,
			Hobi:  []string{"Fotografi", "Traveling", "Membaca"},
			NoHP:  &noHP3,
			Nilai: []int{75, 80, 70},
			Alamat: models.Alamat{
				Jalan:   "Jl. Anggrek No. 7",
				Kota:    "Surabaya",
				KodePos: "60241",
			},
		},
	}

	// Tampilkan data mahasiswa
	for i, mhs := range mahasiswas {
		// validasi nomor HP dan nilai
		err := utils.CekNoHP(mhs.NoHP)
		if err != nil {
			panic(fmt.Sprintf("Data mahasiswa %s: %v", mhs.Nama, err))
		}

		err = utils.ValidasiNilai(mhs.Nilai)
		if err != nil {
			panic(fmt.Sprintf("Data mahasiswa %s: %v", mhs.Nama, err))
		}

		// kalkulasi rata-rata nilai dan status kelulusan
		rata2 := utils.RataRataNilai(mhs.Nilai)
		lulus, keteranganLulus := utils.CekStatusKelulusan(rata2)
		mhs.Lulus = lulus
		mhs.KeteranganLulus = keteranganLulus

		fmt.Println("===============================")
		fmt.Printf("Mahasiswa ke-%d\n", i+1)
		fmt.Println("Nama   :", mhs.Nama)
		fmt.Println("Umur   :", mhs.Umur)
		fmt.Println("Hobi   :", strings.Join(mhs.Hobi, ", "))
		fmt.Printf("Alamat : %s, %s - %s\n", mhs.Alamat.Jalan, mhs.Alamat.Kota, mhs.Alamat.KodePos)

		if mhs.NoHP != nil {
			fmt.Println("No HP  :", *mhs.NoHP)
		} else {
			fmt.Println("No HP  : (belum diisi)")
		}

		fmt.Println("Nilai  :", mhs.Nilai)
		fmt.Printf("Rata-rata Nilai : %.2f\n", rata2)
		fmt.Println("Lulus  :", mhs.Lulus)
		fmt.Println("Keterangan :", mhs.KeteranganLulus)
	}
	fmt.Println("=================================")
}
