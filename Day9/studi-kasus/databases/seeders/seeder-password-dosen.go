package seeders

import (
	"errors"
	"fmt"
	"mahasiswa/configs"
	"mahasiswa/models"
	"mahasiswa/utils"
	"strconv"

	"gorm.io/gorm"
)

func SeederPasswordDosen() {
	err := configs.DB.Where("nik is not null AND password is not null").First(&models.Dosen{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		hashPassword, err := utils.HashPassword("admin123")
		if err != nil {
			fmt.Println("Error", err.Error())
			return
		}

		nik := "DOS1"

		var listDosen []models.Dosen

		err = configs.DB.Find(&listDosen).Error
		if err != nil {
			fmt.Println("Error", err.Error())
			return
		}

		for _, dosen := range listDosen {
			idString := strconv.FormatUint(uint64(dosen.ID), 10)

			nikDosen := nik + idString

			err = configs.DB.Model(&dosen).Updates(models.Dosen{NIK: nikDosen, Password: hashPassword}).Error
			if err != nil {
				fmt.Println("Error", err.Error())
				return
			}
		}
	}
}
