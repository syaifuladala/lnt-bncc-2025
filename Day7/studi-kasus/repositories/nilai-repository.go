package repositories

import (
	"mahasiswa/configs"
	"mahasiswa/models"
)

func InputNilaiMahasiswa(nilai models.NilaiMahasiswa) error {
	err := configs.DB.Create(&nilai).Error
	if err != nil {
		return err
	}

	return nil
}

func GetMataKuliahMahasiswaByID(mahasiswaID uint) (nilaiMahasiswa []models.NilaiMahasiswa, err error) {
	err = configs.DB.Where("mahasiswa_id = ?", mahasiswaID).Order("mata_kuliah_id").Preload("MataKuliah.Dosen").Find(&nilaiMahasiswa).Error
	if err != nil {
		return nil, err
	}

	return nilaiMahasiswa, nil
}
