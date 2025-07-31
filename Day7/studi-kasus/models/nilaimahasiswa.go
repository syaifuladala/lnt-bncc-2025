package models

import "gorm.io/gorm"

type NilaiMahasiswa struct {
	gorm.Model
	MahasiswaID  uint       `json:"mahasiswa_id"`
	Mahasiswa    Mahasiswa  `json:"mahasiswa" gorm:"foreignKey:MahasiswaID"`
	MataKuliahID uint       `json:"mata_kuliah_id"`
	MataKuliah   MataKuliah `json:"mata_kuliah" gorm:"foreignKey:MataKuliahID"`
	Nilai        int        `json:"nilai"`
}
