package repositories

import (
	"mahasiswa/configs"
	"mahasiswa/models"
)

func ListDosen() (dosen []models.Dosen, err error) {
	err = configs.DB.Find(&dosen).Error
	if err != nil {
		return nil, err
	}

	return dosen, nil
}

func GetDosenByID(id uint) (dosen *models.Dosen, err error) {
	err = configs.DB.First(&dosen, id).Error
	if err != nil {
		return nil, err
	}

	return dosen, nil
}

func CreateDosen(dosen models.Dosen) error {
	err := configs.DB.Create(&dosen).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateDosen(id uint, dosen *models.Dosen) error {
	err := configs.DB.Model(&models.Dosen{}).Where("id = ?", id).Updates(dosen).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteDosen(id uint) error {
	err := configs.DB.Delete(&models.Dosen{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func GetDosenByNIK(nik string) (dosen *models.Dosen, err error) {
	err = configs.DB.Where("nik", nik).First(&dosen).Error
	if err != nil {
		return nil, err
	}

	return dosen, nil
}
