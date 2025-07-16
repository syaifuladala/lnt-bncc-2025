package main

import (
	"fmt"
	"kalkulator/kalkulator"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Pulih dari panic: ", r)
		}
	}()
	fmt.Println("Before panic")
	panic("Terjadi kesalahan")
	fmt.Println("Harusnya tidak dieksekusi")

	x := 10
	y := 5

	halo()

	hasilPenjumlahan := tambah(x, y)
	fmt.Println("hasil tambah: ", hasilPenjumlahan)

	hasilPembagian, statusBagi := bagi(x, y)
	fmt.Printf("hasil bagi: %d, status: %s \n", hasilPembagian, statusBagi)

	variadicJumlah := jumlahkan(x, y, 1)
	fmt.Println("variadic jumlah: ", variadicJumlah)

	// modular
	fmt.Println(kalkulator.Tambah(10, 5))

	kurang := kalkulator.Kurang(x, y)
	fmt.Println(kurang)

	hasilBagi, err := kalkulator.Bagi(10, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(hasilBagi)
	}

}

func halo() {
	fmt.Println("Halo!")
}

func tambah(angka1 int, angka2 int) (hasil int) {
	hasil = angka1 + angka2
	return hasil
}

func bagi(angka1 int, angka2 int) (hasil int, status string) {
	hasil = angka1 / angka2
	status = "sukses"
	return hasil, status
}

func jumlahkan(angka ...int) int {
	total := 0

	for _, v := range angka {
		total += v
	}

	return total
}
