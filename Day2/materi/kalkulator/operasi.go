package kalkulator

import "errors"

func Tambah(angka1 int, angka2 int) int {
	return angka1 + angka2
}

func Kurang(angka1 int, angka2 int) int {
	return angka1 - angka2
}

func Kali(angka1 int, angka2 int) int {
	return angka1 * angka2
}

func Bagi(angka1 int, angka2 int) (int, error) {
	if angka2 == 0 {
		return 0, errors.New("tidak bisa dibagi 0")
	}
	return angka1 / angka2, nil
}
