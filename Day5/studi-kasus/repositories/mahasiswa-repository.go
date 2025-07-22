package repositories

import (
	"mahasiswa/configs"
	"mahasiswa/models"
)

func CreateMahasiswa(mahasiswa *models.Mahasiswa) error {
	err := configs.DB.Create(mahasiswa).Error
	if err != nil {
		return err
	}

	return nil
}

func ListMahasiswa() (mahasiswas []models.Mahasiswa, err error) {
	err = configs.DB.Find(&mahasiswas).Error
	if err != nil {
		return nil, err
	}

	return mahasiswas, nil
}

func UpdateMahasiswa(mahasiswa *models.Mahasiswa) error {
	err := configs.DB.Save(mahasiswa).Error
	if err != nil {
		return err
	}

	return nil
}

func GetMahasiswaByID(id uint) (mahasiswa *models.Mahasiswa, err error) {
	err = configs.DB.First(&mahasiswa, id).Error
	if err != nil {
		return nil, err
	}

	return mahasiswa, nil
}

func DeleteMahasiswa(id uint) error {
	var mahasiswa models.Mahasiswa
	err := configs.DB.Delete(&mahasiswa, id).Error
	if err != nil {
		return err
	}

	return nil
}
