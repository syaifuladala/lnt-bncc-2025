package models

import "gorm.io/gorm"

type MataKuliah struct {
	gorm.Model
	Nama    string           `json:"nama"`
	DosenID uint             `json:"dosen_id"`
	Dosen   Dosen            `json:"dosen" gorm:"foreignKey:DosenID"`
	Nilai   []NilaiMahasiswa `json:"nilai" gorm:"foreignKey:MataKuliahID"`
}
