package utils

import (
	"errors"
	"unicode"
)

func ValidasiNilai(nilai []int) error {
	for _, n := range nilai {
		if n < 0 || n > 100 {
			return errors.New("nilai harus dalam rentang 0 - 100")
		}
	}

	return nil
}

func CekNoHP(noHP *string) error {
	if noHP == nil {
		return nil
	}

	if len(*noHP) < 10 || len(*noHP) > 13 {
		return errors.New("nomor HP harus 10-13 digit")
	}

	for _, r := range *noHP {
		if !unicode.IsDigit(r) {
			return errors.New("nomor HP hanya boleh berisi angka")
		}
	}

	return nil
}
