package repositories

import (
	"mahasiswa/configs"
	"mahasiswa/models"
)

func ListMataKuliah() (mataKuliah []models.MataKuliah, err error) {
	err = configs.DB.Find(&mataKuliah).Error
	if err != nil {
		return nil, err
	}

	return mataKuliah, nil
}

func GetMataKuliahByID(id uint) (mataKuliah *models.MataKuliah, err error) {
	err = configs.DB.First(&mataKuliah, id).Error
	if err != nil {
		return nil, err
	}

	return mataKuliah, nil
}

func CreateMataKuliah(mataKuliah models.MataKuliah) error {
	err := configs.DB.Create(&mataKuliah).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateMataKuliah(id uint, mataKuliah *models.MataKuliah) error {
	err := configs.DB.Model(&models.MataKuliah{}).Where("id = ?", id).Updates(mataKuliah).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteMataKuliah(id uint) error {
	err := configs.DB.Delete(&models.MataKuliah{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
