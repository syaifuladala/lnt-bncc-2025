package models

import "gorm.io/gorm"

type Dosen struct {
	gorm.Model
	Nama       string       `json:"nama"`
	MataKuliah []MataKuliah `json:"mata_kuliah" gorm:"foreignKey:DosenID"`
}
