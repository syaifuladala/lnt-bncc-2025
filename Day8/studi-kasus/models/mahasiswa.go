package models

import "gorm.io/gorm"

type Mahasiswa struct {
	gorm.Model
	Nama     string           `json:"nama"`
	NIM      string           `json:"nim" gorm:"unique"`
	Umur     int              `json:"umur"`
	Hobi     string           `json:"hobi"`
	Alamat   string           `json:"alamat"`
	Kota     string           `json:"kota"`
	KodePos  string           `json:"kode_pos"`
	NoHP     *string          `json:"no_hp"`
	RataRata float64          `json:"rata_rata"`
	Lulus    bool             `json:"lulus" gorm:"default:0"`
	Nilai    []NilaiMahasiswa `json:"nilai_mahasiswa" gorm:"foreignKey:MahasiswaID"`
}
